package route

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/donnol/jdnote/config"
	pg "github.com/donnol/jdnote/store/db/postgresql"
	"github.com/donnol/jdnote/utils/jwt"
	utillog "github.com/donnol/jdnote/utils/log"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/schema"
)

// 参数相关
var (
	decoder = schema.NewDecoder()
)

// cookie相关
var (
	sessionKey = "jd_session"

	jwtToken = jwt.New([]byte(config.DefaultConfig.JWT.Secret))
)

// DefaultRouter 默认路由
var DefaultRouter = NewRouter()

// Router 路由
type Router struct {
	*gin.Engine
}

// NewRouter 新建路由
func NewRouter() *Router {
	router := gin.Default()
	gin.DefaultWriter = io.MultiWriter(os.Stdout)
	return &Router{
		Engine: router,
	}
}

// Param 通用参数
type Param struct {
	UserID       int         `json:"userID"`       // 用户ID
	RequestParam interface{} `json:"requestParam"` // 请求参数

	// 方法
	method string

	// 参数
	body   []byte
	values url.Values
}

// Parse 解析
func (p *Param) Parse(v interface{}) error {
	var err error

	// 解析
	switch p.method {
	case http.MethodPost:
		fallthrough
	case http.MethodPut:
		err = json.Unmarshal(p.body, v)
	case http.MethodGet:
		fallthrough
	case http.MethodDelete:
		err = decoder.Decode(v, p.values)
	}
	if err != nil {
		return err
	}

	// 检查参数
	if vv, ok := v.(Checker); ok {
		if err := vv.Check(); err != nil {
			return err
		}
	}

	return nil
}

// Error 错误
type Error struct {
	Code int    `json:"code"` // 请求返回码，一般0表示正常，非0表示异常
	Msg  string `json:"msg"`  // 信息，一般是出错时的描述信息
}

// Error 实现error接口
func (e Error) Error() string {
	return fmt.Sprintf("Code: %d, Msg: %s", e.Code, e.Msg)
}

// 确保Error实现了error接口
var _ error = Error{}

// Result 通用结果
type Result struct {
	Error

	Data interface{} `json:"data"` // 正常返回时的数据

	// 给登陆接口使用
	CookieAfterLogin int `json:"-"` // 登陆时需要设置登陆态的用户信息
}

// PresentData 用具体结构体展现数据
func (r *Result) PresentData(v interface{}) error {
	b, err := json.Marshal(r.Data)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, v); err != nil {
		return err
	}

	return nil
}

// HandlerFunc 处理函数
type HandlerFunc func(Param) (Result, error)

// Register 注册结构体
// 结构体名字作为路径的第一部分，路径后面部分由可导出方法名映射来
func (r *Router) Register(v interface{}) {
	// 计时开始
	start := time.Now()

	// 初始化
	// 如果有实现New方法，则调用
	if vv, ok := v.(Newer); ok {
		v = vv.New()
	} else {
		var err error

		// 注入DB
		v, err = pg.InitParamWithDB(v)
		if err != nil {
			panic(err)
		}

		// 注入Logger
		v, err = utillog.InitParamWithLogger(v)
		if err != nil {
			panic(err)
		}

		utillog.Debugf("After init: %+v\n", v)
	}

	// 反射获取Type
	var structName string
	refType := reflect.TypeOf(v)
	refTypeRaw := refType
	refValue := reflect.ValueOf(v)
	if refType.Kind() == reflect.Ptr {
		structName = refType.Elem().Name()
		refTypeRaw = refType.Elem()
	} else {
		structName = refType.Name()
	}

	// 找出attr field
	var groupName string
	groupType := reflect.TypeOf(Group{})
	for i := 0; i < refTypeRaw.NumField(); i++ {
		field := refTypeRaw.Field(i)
		if field.Type == groupType {
			groupName = strings.ToLower(field.Name)
		}
	}

	// 找出method field
	for i := 0; i < refType.NumMethod(); i++ {
		field := refType.Method(i)
		value := refValue.Method(i)

		// 路径
		method, path := getMethodPath(field.Name)
		path = addPathPrefix(path, structName)
		path = addPathPrefix(path, groupName)

		// 方法
		valueFunc, ok := value.Interface().(func(Param) (Result, error))
		if !ok {
			continue
		}

		// 注册路由
		switch method {
		case http.MethodPost:
			r.Engine.POST(path, structHandlerFunc(http.MethodPost, valueFunc))
		case http.MethodPut:
			r.Engine.PUT(path, structHandlerFunc(http.MethodPut, valueFunc))
		case http.MethodGet:
			r.Engine.GET(path, structHandlerFunc(http.MethodGet, valueFunc))
		case http.MethodDelete:
			r.Engine.DELETE(path, structHandlerFunc(http.MethodDelete, valueFunc))
		default:
			panic("Not support method now.")
		}
	}

	// 计时结束
	end := time.Now()
	utillog.Debugf("Register %s router use time: %v\n", structName, end.Sub(start))
}

// structHandlerFunc 结构体处理函数
func structHandlerFunc(method string, f HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		// 获取参数
		var body []byte
		var values url.Values
		switch method {
		case http.MethodPost:
			fallthrough
		case http.MethodPut:
			body, err = c.GetRawData()
		case http.MethodGet:
			fallthrough
		case http.MethodDelete:
			values = c.Request.URL.Query()
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 获取用户信息
		var userID int
		cookie, err := c.Cookie(sessionKey)
		if err == nil {
			userID, err = jwtToken.Verify(cookie)
			if err != nil {
				utillog.Warnf("token verify failed, err: %+v\n", err)
				userID = 0
			}
		}

		// 注入用户和参数信息，并执行业务方法
		p := Param{UserID: userID, method: method, body: body, values: values}
		r, err := f(p)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 设置header
		// 格式
		c.Header("Content-Type", "application/json; charset=utf-8")
		// 跨域
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		// cookie
		if r.CookieAfterLogin != 0 {
			var session string
			session, err = jwtToken.Sign(r.CookieAfterLogin)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			var maxAge = 3600 * 24 * 7
			cookie := fmt.Sprintf("%s=%s; HttpOnly; max-age=%d", sessionKey, session, maxAge)
			c.Header("Set-Cookie", cookie)
		}

		// 调用过滤器，过滤返回内容
		if v, ok := r.Data.(Filter); ok {
			r.Data = v.Filter()
		}

		// 返回
		c.JSON(http.StatusOK, r)
	}
}
