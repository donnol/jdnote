package route

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"runtime"
	"strings"
	"unicode"

	"github.com/donnol/jdnote/config"
	pg "github.com/donnol/jdnote/store/db/postgresql"
	"github.com/donnol/jdnote/utils/jwt"
	utillog "github.com/donnol/jdnote/utils/log"
	"github.com/gin-gonic/gin"
)

// 路径分割符
const (
	pathSep = "/"
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

	body []byte // 参数
}

// Parse 解析
func (p *Param) Parse(v interface{}) error {
	if err := json.Unmarshal(p.body, v); err != nil {
		return err
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

// Register 注册
func (r *Router) Register(param interface{}, f HandlerFunc) {
	// 通过f的名字获取method，path
	method, path := getMethodPathFromFunc(f)

	switch method {
	case http.MethodPost:
		r.Engine.POST(path, defaultHandlerFunc(http.MethodPost, param, f))
	case http.MethodPut:
		r.Engine.PUT(path, defaultHandlerFunc(http.MethodPut, param, f))
	case http.MethodGet:
		r.Engine.GET(path, defaultHandlerFunc(http.MethodGet, param, f))
	case http.MethodDelete:
		r.Engine.DELETE(path, defaultHandlerFunc(http.MethodDelete, param, f))
	default:
		panic("Not support method now.")
	}
}

// RegisterStruct 注册结构体
// 结构体名字作为路径的第一部分，路径后面部分由可导出方法名映射来
func (r *Router) RegisterStruct(v interface{}) {
	// 初始化
	v = initParamWithDB(v, pg.New())

	// 反射获取Type
	var structName string
	refType := reflect.TypeOf(v)
	refValue := reflect.ValueOf(v)
	if refType.Kind() == reflect.Ptr {
		structName = refType.Elem().Name()
	} else {
		structName = refType.Name()
	}

	// 找出method field
	for i := 0; i < refType.NumMethod(); i++ {
		field := refType.Method(i)
		value := refValue.Method(i)

		//  路径
		method, path := getMethodPath(field.Name)
		path = addPathPrefix(path, structName)

		// 方法
		valueFunc := value.Interface().(func(Param) (Result, error))

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
}

func addPathPrefix(path, prefix string) string {
	return pathSep + strings.ToLower(prefix) + path
}

// getMethodPathFromFunc 通过f的名字获取method，path
func getMethodPathFromFunc(f HandlerFunc) (method, path string) {
	// 利用反射和运行时获取函数名
	refValue := reflect.ValueOf(f)
	fn := runtime.FuncForPC(refValue.Pointer())
	fullFuncName := fn.Name()

	return getMethodPath(fullFuncName)
}

// getMethodPath 获取methid, path
func getMethodPath(fullFuncName string) (method, path string) {
	const sep = "."

	upperFunc := func(r rune) bool {
		return unicode.IsUpper(r)
	}

	// 过滤函数名的包名部分
	lastDotIndex := strings.LastIndex(fullFuncName, sep)
	funcName := fullFuncName[lastDotIndex+1:]

	// 找到函数名里的首个大写字母，并以此作为依据将字符串分割
	firstUpperIndex := strings.IndexFunc(funcName, upperFunc)
	if firstUpperIndex == 0 {
		// 如果方法是可导出的，首字母就是大写，需要过滤掉
		firstUpperIndex = strings.IndexFunc(funcName[1:], upperFunc) + 1
	}
	method = funcName[:firstUpperIndex]
	method = methodMap(method)

	// 如果剩下的路径部分还有大写字母，需要分为多段路径
	tmpPath := funcName[firstUpperIndex:]
	for {
		tmpPath = strings.ToLower(tmpPath[:1]) + tmpPath[1:]
		firstUpperIndex = strings.IndexFunc(tmpPath, upperFunc)
		if firstUpperIndex == -1 {
			path += pathSep + strings.ToLower(tmpPath)
			return
		}
		path += pathSep + strings.ToLower(tmpPath[:firstUpperIndex])

		tmpPath = tmpPath[firstUpperIndex:]
	}

	return
}

func methodMap(m string) (r string) {
	m = strings.ToLower(m)
	switch m {
	case "get":
		r = http.MethodGet
	case "add":
		r = http.MethodPost
	case "mod":
		r = http.MethodPut
	case "del":
		r = http.MethodDelete
	default:
		r = http.MethodPost
	}
	return
}

// initParamWithDB 初始化-使用反射初始化param里的DB
func initParamWithDB(param interface{}, db pg.DB) interface{} {
	// 校验类型
	refType := reflect.TypeOf(param)
	refValue := reflect.ValueOf(param)
	if refType.Kind() == reflect.Ptr {
		refType = refType.Elem()
		refValue = refValue.Elem()
	}
	if refType.Kind() != reflect.Struct {
		panic("Please input struct param!")
	}

	// db类型
	dbType := reflect.TypeOf((*pg.DB)(nil)).Elem()
	dbValue := reflect.ValueOf(db)

	// 注入DB
	setValue(refType, dbType, refValue, dbValue)

	return param
}

func setValue(refType, dbType reflect.Type, refValue, dbValue reflect.Value) {
	for i := 0; i < refType.NumField(); i++ {
		field := refType.Field(i)
		if field.Type == dbType { // 类型相同，直接赋值
			v := refValue.Field(i)
			v.Set(dbValue)
		} else if field.Type.Implements(dbType) { // 内嵌类型，递归遍历
			setValue(field.Type, dbType, refValue.Field(i), dbValue)
		}
	}
}

var structHandlerFunc = func(method string, f HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		// 获取参数
		var body []byte
		switch method {
		case http.MethodPost:
			fallthrough
		case http.MethodPut:
			body, err = c.GetRawData()
		case http.MethodGet:
			fallthrough
		case http.MethodDelete:
			var queryMap = make(map[string]interface{})
			values := c.Request.URL.Query()
			for k, v := range values {
				if len(v) == 1 {
					queryMap[k] = v[0]
				} else {
					queryMap[k] = v
				}
			}
			body, err = json.Marshal(queryMap)
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		log.Printf("body: %s\n", body)

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

		// 注入用户信息，并执行业务方法
		p := Param{UserID: userID, body: body}
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

		c.JSON(http.StatusOK, r)
	}
}

var defaultHandlerFunc = func(method string, param interface{}, f HandlerFunc) gin.HandlerFunc {
	// 如果有实现New方法，则调用
	if v, ok := param.(Newer); ok {
		param = v.New()
	} else {
		// 注入DB
		param = initParamWithDB(param, pg.New())
	}

	return func(c *gin.Context) {
		var err error

		switch method {
		case http.MethodPost:
			fallthrough
		case http.MethodPut:
			err = c.ShouldBindJSON(param)
		case http.MethodGet:
			fallthrough
		case http.MethodDelete:
			err = c.ShouldBindQuery(param)
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 检查参数
		if v, ok := param.(Checker); ok {
			if err := v.Check(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
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

		p := Param{UserID: userID, RequestParam: param}
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

		c.JSON(http.StatusOK, r)
	}
}
