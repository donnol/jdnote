package route

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/donnol/jdnote/utils/context"
	"github.com/donnol/jdnote/utils/errors"
	"github.com/donnol/jdnote/utils/jwt"

	"github.com/donnol/tools/log"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/rcrowley/go-metrics"
	"golang.org/x/time/rate"
)

// header相关
const (
	ContentDispositionHeaderKey         = "Content-Disposition"
	ContentDispositionHeaderValueFormat = `attachment; filename="%s"`

	setCookieHeaderKey     = "Set-Cookie"
	contentTypeHeaderKey   = "Content-Type"
	contentTypeHeaderValue = "application/json; charset=utf-8"

	// 跨域
	accessOriginHeaderKey         = "Access-Control-Allow-Origin"
	accessOriginHeaderValue       = "*"
	accessCreadentialsHeaderKey   = "Access-Control-Allow-Credentials"
	accessCreadentialsHeaderValue = "true"
)

// Router 路由
type Router struct {
	*gin.Engine

	sessionKey string
	jwtToken   *jwt.Token
}

type Option struct {
	SessionKey string
	JwtToken   *jwt.Token
}

// NewRouter 新建路由
func NewRouter(opt Option) *Router {
	router := gin.Default()

	router.Use(gzip.Gzip(gzip.DefaultCompression))

	gin.DefaultWriter = io.MultiWriter(os.Stdout)

	return &Router{
		Engine: router,

		sessionKey: opt.SessionKey,
		jwtToken:   opt.JwtToken,
	}
}

// HandlerFunc 处理函数
// 使用别名，可以互相替换，但是不能添加方法
// 使用类型，不可以互相替换，需要转型，但是可以添加方法
type HandlerFunc = func(context.Context, Param) (Result, error)

type limiterOption struct {
	rate float64 // 代表每秒可以向Token桶中产生多少token
	b    int     // 代表Token桶的容量大小
}

// Register 注册结构体
// 结构体名字作为路径的第一部分，路径后面部分由可导出方法名映射来
func (r *Router) Register(ctx context.Context, v interface{}) {
	// 计时开始
	start := time.Now()

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

	// 找出路由属性
	routeAtrr := getRouteAttr(refTypeRaw)

	// 找出method field
	var routeNum int
	for i := 0; i < refType.NumMethod(); i++ {
		field := refType.Method(i)
		value := refValue.Method(i)

		// 方法
		valueFunc, ok := value.Interface().(HandlerFunc)
		if !ok {
			continue
		}
		routeNum++

		// 路径
		method, path := getMethodPath(field.Name)
		path = addPathPrefix(path, structName)
		path = addPathPrefix(path, routeAtrr.groupName)

		// 处理器配置
		var ho = handlerOption{
			sessionKey: r.sessionKey,
			jwtToken:   r.jwtToken,
		}
		if routeAtrr.isFile {
			ho.isFile = true
		} else {
			if _, ok := routeAtrr.fileMap[strings.ToLower(field.Name)]; ok {
				ho.isFile = true
			}
		}
		if routeAtrr.isTx {
			ho.useTx = true
		} else {
			if _, ok := routeAtrr.methodTxMap[field.Name]; ok {
				ho.useTx = true
			}
		}

		handler := structHandlerFunc(ctx, method, valueFunc, ho)

		wo := wrapOption{
			fieldName: field.Name,
			method:    method,
			path:      path,
		}

		// 添加中间件：我要知道我要不要用，用什么，用的参数
		// 限流: 每个路径对应一个限流器
		handler = wrapLimiter(handler, routeAtrr, wo)

		// 指标
		handler = wrapMetrics(handler, wo)

		// 注册路由
		switch method {
		case http.MethodPost,
			http.MethodPut,
			http.MethodGet,
			http.MethodDelete:
			r.Engine.Handle(method, path, handler)
		default:
			panic("Not support method now.")
		}
	}

	// 计时结束
	end := time.Now()
	log.Debugf("Register %s struct %d routers use time: %v\n\n", structName, routeNum, end.Sub(start))
}

type wrapOption struct {
	fieldName string
	method    string
	path      string
}

func wrapLimiter(handler gin.HandlerFunc, routeAtrr routeAttr, wo wrapOption) gin.HandlerFunc {
	var limiter *rate.Limiter
	if lo, ok := routeAtrr.limiterMap[limiterTagRateName]; ok {
		limiter = rate.NewLimiter(rate.Limit(lo.rate), lo.b)
	} else if mlo, ok := routeAtrr.methodLimiterMap[wo.fieldName]; ok {
		limiter = rate.NewLimiter(rate.Limit(mlo.rate), mlo.b)
	}
	if limiter == nil {
		return handler
	}

	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, Result{Error: errors.Error{
				Code: errors.ErrorCodeRouter,
				Msg:  "Too Many Requests",
			}})
			return
		}
		handler(c)
	}
}

func wrapMetrics(handler gin.HandlerFunc, wo wrapOption) gin.HandlerFunc {
	m := metrics.NewMeter()
	if err := metrics.Register(wo.method+" "+wo.path, m); err != nil {
		log.Warnf("Register metrics failed: %+v\n", err)
		return handler
	}
	m.Mark(0)

	// TODO:改为存储到时序数据库
	// go metrics.Log(
	// 	metrics.DefaultRegistry,
	// 	5*time.Second,
	// 	log.New(os.Stdout, "metrics: ", log.Lmicroseconds),
	// )

	return func(c *gin.Context) {
		handler(c)
		m.Mark(1)
	}
}

type routeAttr struct {
	groupName        string
	isFile           bool
	fileMap          map[string]struct{}
	isTx             bool
	methodTxMap      map[string]struct{}
	limiterMap       map[string]limiterOption
	methodLimiterMap map[string]limiterOption
}

const (
	fileTagSep           = ","
	fileTagName          = "file"
	methodTxTagName      = "tx"
	limiterTagMethodName = "method"
	limiterTagMethodSep  = ";"
	limiterTagRateName   = "rate"
)

var (
	groupType   = reflect.TypeOf(Group{})
	fileType    = reflect.TypeOf(File{})
	methodType  = reflect.TypeOf(Method{})
	limiterType = reflect.TypeOf(Limiter{})
)

func getRouteAttr(refTypeRaw reflect.Type) (ra routeAttr) {
	var groupName string
	var fileMap = make(map[string]struct{})
	var isFile bool
	var methodTxMap = make(map[string]struct{})
	var isTx bool
	var limiterMap = make(map[string]limiterOption)
	var methodLimiterMap = make(map[string]limiterOption)
	for i := 0; i < refTypeRaw.NumField(); i++ {
		field := refTypeRaw.Field(i)

		switch field.Type {
		// Group属性
		case groupType:
			groupName = strings.ToLower(field.Name)

		// File属性
		case fileType:
			fileTag, ok := field.Tag.Lookup(fileTagName)
			// 没有使用tag指定方法，则全部方法都是
			if !ok {
				isFile = true
			} else {
				fileTagList := strings.Split(fileTag, fileTagSep)
				for _, single := range fileTagList {
					singleLower := strings.ToLower(single)
					fileMap[singleLower] = struct{}{}
				}
			}

		// Method属性
		case methodType:
			// 事务
			methodTxTag, ok := field.Tag.Lookup(methodTxTagName)
			if !ok {
				isTx = true
			} else {
				methodTxTags := strings.Split(methodTxTag, fileTagSep)
				for _, single := range methodTxTags {
					methodTxMap[single] = struct{}{}
				}
			}

		// Limiter属性
		case limiterType:
			if methodTag, ok := field.Tag.Lookup(limiterTagMethodName); ok { // 有指定方法
				limiterTags := strings.Split(methodTag, limiterTagMethodSep)
				for _, single := range limiterTags {
					name, values, _, err := resolveCallExpr(single)
					if err != nil {
						panic(err)
					}
					rate := values[0].(float64)
					b := values[1].(int)
					methodLimiterMap[name] = limiterOption{
						rate: rate,
						b:    b,
					}
				}
			}
			if rateTag, ok := field.Tag.Lookup(limiterTagRateName); ok { // 全部指定
				_, values, _, err := resolveCallExpr(rateTag)
				if err != nil {
					panic(err)
				}
				rate := values[0].(float64)
				b := values[1].(int)
				limiterMap[limiterTagRateName] = limiterOption{
					rate: rate,
					b:    b,
				}
			}
		}
	}

	ra.groupName = groupName
	ra.isFile = isFile
	ra.fileMap = fileMap
	ra.isTx = isTx
	ra.methodTxMap = methodTxMap
	ra.limiterMap = limiterMap
	ra.methodLimiterMap = methodLimiterMap

	return
}

type handlerOption struct {
	isFile     bool // 是否文件上传/下载接口
	useTx      bool // 是否使用事务
	sessionKey string
	jwtToken   *jwt.Token
}

// structHandlerFunc 结构体处理函数
func structHandlerFunc(ctx context.Context, method string, f HandlerFunc, ho handlerOption) gin.HandlerFunc {
	log.Debugf("handler option: %+v\n", ho)

	// 处理请求的函数
	return func(c *gin.Context) {
		var r Result
		var err error

		// 先确定时间和地点，然后是用户和请求
		now := time.Now()
		nowTimestamp := now.Unix()
		r.Timestamp = nowTimestamp

		remoteAddr := c.Request.RemoteAddr

		var userID int
		cookie, err := c.Cookie(ho.sessionKey)
		if err == nil {
			verifyUserID, err := ho.jwtToken.Verify(cookie)
			if err != nil {
				log.Warnf("Verify cookie failed: %+v\n", err)
			} else {
				userID = verifyUserID
			}
		} else {
			log.Warnf("Get cookie failed: %+v\n", err)
		}

		reqUUID, err := uuid.NewV4()
		if err != nil {
			log.Default().Errorf("New request id failed: %+v\n", err)
		}
		reqID := reqUUID.String()
		r.RequestID = reqID

		// 获取参数
		var body []byte
		var values url.Values
		switch method {
		case http.MethodPost:
			fallthrough
		case http.MethodPut:
			// GetRawData = ioutil.ReadAll(c.Request.Body)
			// 所以下面的multipartReader会一直报‘multipart: NextPart: EOF’错误
			if !ho.isFile {
				body, err = c.GetRawData()
			}
		case http.MethodGet:
			fallthrough
		case http.MethodDelete:
			values = c.Request.URL.Query()
		}
		if err != nil {
			r.Error.Code = errors.ErrorCodeRouter
			r.Error.Msg = fmt.Sprintf("%+v", err)
			c.JSON(http.StatusNotAcceptable, r)
			return
		}

		// 这里要知道路由是不是文件上传/下载接口，然后将内容传递/返回给f
		var multipartReader *multipart.Reader
		if ho.isFile && method == http.MethodPost {
			multipartReader, err = c.Request.MultipartReader()
			if err != nil {
				r.Error.Code = errors.ErrorCodeRouter
				r.Error.Msg = fmt.Sprintf("%+v", err)
				c.JSON(http.StatusMethodNotAllowed, r)
				return
			}
		}

		// 注入上下文、用户和参数信息，并执行业务方法
		var statusCode = http.StatusOK
		var param = Param{method: method, body: body, values: values, multipartReader: multipartReader}
		ctx = context.WithValue(ctx, context.TimestampKey, nowTimestamp)
		ctx = context.WithValue(ctx, context.RemoteAddrKey, remoteAddr)
		ctx = context.WithValue(ctx, context.UserKey, userID)
		ctx = context.WithValue(ctx, context.RequestKey, reqID)
		if ho.useTx {
			// 事务-统一从这里开启。srv和db不需要理会事务，只需要使用ctx.DB()返回的实例去操作即可
			// 即使是相同的请求，每次进来都会是一个新事务，所以基本上是没有事务嵌套的问题的
			err = context.WithTx(ctx, func(ctx context.Context) error {
				var err error

				r, err = f(ctx, param)
				if err != nil {
					return err
				}

				return nil
			})
		} else {
			r, err = f(ctx, param)
		}
		r.Timestamp = nowTimestamp
		r.RequestID = reqID
		// 处理错误
		if e, ok := err.(errors.Error); ok {
			if e.IsNormal() {
				statusCode = http.StatusBadRequest
			} else if e.IsFatal() {
				statusCode = http.StatusInternalServerError
			}
			r.Error = e
		} else {
			if err != nil {
				r.Error.Code = errors.ErrorCodeRouter
				r.Error.Msg = fmt.Sprintf("%+v", err)
				c.JSON(http.StatusForbidden, r)
				return
			}
		}

		// 设置header
		// 格式
		c.Header(contentTypeHeaderKey, contentTypeHeaderValue)
		// 跨域
		c.Header(accessOriginHeaderKey, accessOriginHeaderValue)
		c.Header(accessCreadentialsHeaderKey, accessCreadentialsHeaderValue)
		// cookie
		if r.CookieAfterLogin != 0 {
			cookie, err := MakeCookie(r.CookieAfterLogin, CookieOption{
				SessionKey: ho.sessionKey,
				JwtToken:   ho.jwtToken,
			})
			if err != nil {
				r.Error.Code = errors.ErrorCodeRouter
				r.Error.Msg = fmt.Sprintf("%+v", err)
				c.JSON(http.StatusInternalServerError, r)
				return
			}
			c.Header(setCookieHeaderKey, cookie.String())
		}

		// 调用过滤器，过滤返回内容
		// Filter的存在，使用同一结构，能在不同请求里返回不一样的字段
		if v, ok := r.Data.(Filter); ok {
			r.Data = v.Filter()
		}

		// 返回文件内容
		if ho.isFile && method == http.MethodGet {
			// 重新设置文件的Content-Type
			c.Header(contentTypeHeaderKey, r.ContentType)
			c.DataFromReader(statusCode, r.ContentLength, r.ContentType, r.ContentReader, r.ExtraHeaders)
			return
		}

		// 返回
		c.JSON(statusCode, r)
	}
}

type CookieOption struct {
	SessionKey string
	JwtToken   *jwt.Token
}

// MakeCookie 新建令牌
func MakeCookie(userID int, co CookieOption) (cookie http.Cookie, err error) {
	session, err := co.JwtToken.Sign(userID)
	if err != nil {
		return
	}

	days := 7
	var maxAge = 3600 * 24 * days

	cookie.Name = co.SessionKey
	cookie.Value = session
	cookie.MaxAge = maxAge
	cookie.Expires = time.Now().AddDate(0, 0, days)
	cookie.Path = "/"
	cookie.HttpOnly = true

	return
}
