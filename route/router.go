package route

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/donnol/jdnote/config"
	"github.com/donnol/jdnote/utils/jwt"
	utillog "github.com/donnol/jdnote/utils/log"
	"github.com/gin-gonic/gin"
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
}

// Result 通用结果
type Result struct {
	Code int         `json:"code"` // 请求返回码，一般0表示正常，非0表示异常
	Msg  string      `json:"msg"`  // 信息，一般是出错时的描述信息
	Data interface{} `json:"data"` // 正常返回时的数据

	// 给登陆接口使用
	CookieAfterLogin int `json:"-"` // 登陆时需要设置登陆态的用户信息
}

// Register 注册
func (r *Router) Register(method, path string, param interface{}, f func(Param) (Result, error)) {
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

var defaultHandlerFunc = func(method string, param interface{}, f func(Param) (Result, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		// 如果有实现New方法，则调用
		if v, ok := param.(Newer); ok {
			param = v.New()
		}

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
