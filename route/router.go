package route

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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

// Register 注册
func (r *Router) Register(method, path string, param interface{}, f func(interface{}) (interface{}, error)) {
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

var defaultHandlerFunc = func(method string, param interface{}, f func(interface{}) (interface{}, error)) gin.HandlerFunc {
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

		r, err := f(param)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, r)
	}
}
