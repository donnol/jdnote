package route

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/donnol/jdnote/config"
	"github.com/donnol/jdnote/models"
	"github.com/donnol/jdnote/utils/context"
	utilerrors "github.com/donnol/jdnote/utils/errors"
	"github.com/donnol/jdnote/utils/jwt"
	utillog "github.com/donnol/jdnote/utils/log"
	"github.com/donnol/jdnote/utils/store/db"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/schema"
	"github.com/pkg/errors"
)

// 参数相关
var (
	decoder = schema.NewDecoder()
)

// cookie相关
var (
	sessionKey = "jd_session"

	jwtToken = jwt.New([]byte(config.Default().JWT.Secret))
)

// header相关
var (
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

// defaultRouter 默认路由
var defaultRouter = NewRouter()

// DefaultRouter 获取默认路由
func DefaultRouter() *Router {
	return defaultRouter
}

// Register 注册
func Register(v interface{}) {
	defaultRouter.Register(v)
}

// Router 路由
type Router struct {
	*gin.Engine
}

// NewRouter 新建路由
func NewRouter() *Router {
	router := gin.Default()

	router.Use(gzip.Gzip(gzip.DefaultCompression))

	gin.DefaultWriter = io.MultiWriter(os.Stdout)

	return &Router{
		Engine: router,
	}
}

// Param 通用参数
type Param struct {
	// 方法
	method string

	// 参数
	body   []byte
	values url.Values

	// 文件
	multipartReader *multipart.Reader
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
		return errors.WithStack(err)
	}

	// 检查参数
	if vv, ok := v.(Checker); ok {
		if err := vv.Check(); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

// ParseMultipartForm 解析内容
func (p *Param) ParseMultipartForm(maxFileSize int64, v interface{}) ([]byte, error) {
	var body []byte

	if p.multipartReader == nil {
		return body, fmt.Errorf("Bad multipart reader")
	}

	// 使用ReadForm
	form, err := p.multipartReader.ReadForm(maxFileSize)
	if err != nil {
		return body, err
	}

	// 获取参数
	if err := decoder.Decode(v, form.Value); err != nil {
		return body, err
	}

	// 获取内容
	buf := new(bytes.Buffer)
	for _, single := range form.File {
		for _, one := range single {
			file, err := one.Open()
			if err != nil {
				return body, err
			}
			defer file.Close()

			_, err = buf.ReadFrom(file)
			if err != nil {
				return body, err
			}
		}
	}
	body = buf.Bytes()

	return body, nil
}

// Result 通用结果
type Result struct {
	utilerrors.Error

	Data interface{} `json:"data"` // 正常返回时的数据

	// 给登陆接口使用
	CookieAfterLogin int `json:"-"` // 登陆时需要设置登陆态的用户信息

	// 下载内容时使用
	Content

	err error // 错误
}

// Content 内容
type Content struct {
	ContentLength int64             `json:"-"`
	ContentType   string            `json:"-"`
	ContentReader io.Reader         `json:"-"`
	ExtraHeaders  map[string]string `json:"-"`
}

// MakeContentFromBuffer 新建内容
func MakeContentFromBuffer(filename string, buf *bytes.Buffer) Content {
	var r Content

	writer := multipart.NewWriter(buf)
	r.ContentLength = int64(buf.Len())
	r.ContentType = writer.FormDataContentType()
	r.ContentReader = buf
	r.ExtraHeaders = map[string]string{
		ContentDispositionHeaderKey: fmt.Sprintf(
			ContentDispositionHeaderValueFormat,
			filename,
		),
	}

	return r
}

// SetErr 设置错误
func (r *Result) SetErr(err error) *Result {
	if r.err == nil && err != nil {
		r.err = err
	}
	return r
}

// Err 获取错误
func (r *Result) Err() error {
	return r.err
}

// ErrIsNil 错误是否存在
func (r *Result) ErrIsNil() bool {
	return r.err == nil
}

// Unwrap 如果错误不为nil，则panic
func (r *Result) Unwrap() interface{} {
	if r.err != nil {
		panic(r.err)
	}
	return r.Data
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
type HandlerFunc func(context.Context, Param) Result

// Register 注册结构体
// 结构体名字作为路径的第一部分，路径后面部分由可导出方法名映射来
func (r *Router) Register(v interface{}) {
	// 计时开始
	start := time.Now()

	// 初始化
	// 如果有实现New方法，则调用
	if vv, ok := v.(Newer); ok {
		v = vv.New()
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
	const (
		fileTagLeft     = "("
		fileTagRight    = ")"
		fileTagSep      = ","
		fileTagName     = "file"
		methodTxTagName = "tx"
	)
	var groupName string
	var fileMap = make(map[string]struct{})
	var isFile bool
	var methodTxMap = make(map[string]struct{})
	var isTx bool
	groupType := reflect.TypeOf(Group{})
	fileType := reflect.TypeOf(File{})
	methodType := reflect.TypeOf(Method{})
	for i := 0; i < refTypeRaw.NumField(); i++ {
		field := refTypeRaw.Field(i)

		// Group属性
		if field.Type == groupType {
			groupName = strings.ToLower(field.Name)
		}

		// File属性
		if field.Type == fileType {
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
		}

		// Method属性
		if field.Type == methodType {
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
		}
	}

	// 找出method field
	for i := 0; i < refType.NumMethod(); i++ {
		field := refType.Method(i)
		value := refValue.Method(i)

		// 方法
		valueFunc, ok := value.Interface().(func(context.Context, Param) Result)
		if !ok {
			continue
		}

		// 路径
		method, path := getMethodPath(field.Name)
		path = addPathPrefix(path, structName)
		path = addPathPrefix(path, groupName)

		// 处理器配置
		var ho = handlerOption{}
		if isFile {
			ho.isFile = true
		} else {
			if _, ok := fileMap[strings.ToLower(field.Name)]; ok {
				ho.isFile = true
			}
		}
		if isTx {
			ho.useTx = true
		} else {
			if _, ok := methodTxMap[field.Name]; ok {
				ho.useTx = true
			}
		}

		// 注册路由
		switch method {
		case http.MethodPost:
			r.Engine.POST(path, structHandlerFunc(http.MethodPost, valueFunc, ho))
		case http.MethodPut:
			r.Engine.PUT(path, structHandlerFunc(http.MethodPut, valueFunc, ho))
		case http.MethodGet:
			r.Engine.GET(path, structHandlerFunc(http.MethodGet, valueFunc, ho))
		case http.MethodDelete:
			r.Engine.DELETE(path, structHandlerFunc(http.MethodDelete, valueFunc, ho))
		default:
			panic("Not support method now.")
		}
	}

	// 计时结束
	end := time.Now()
	utillog.Debugf("Register %s router use time: %v\n", structName, end.Sub(start))
}

type handlerOption struct {
	isFile bool // 是否文件上传/下载接口
	useTx  bool // 是否使用事务
}

// structHandlerFunc 结构体处理函数
func structHandlerFunc(method string, f HandlerFunc, ho handlerOption) gin.HandlerFunc {
	utillog.Debugf("handler option: %+v\n", ho)

	return func(c *gin.Context) {
		var err error

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
			c.JSON(http.StatusNotAcceptable, Result{Error: utilerrors.Error{
				Code: utilerrors.ErrorCodeRouter,
				Msg:  fmt.Sprintf("%+v", err),
			}})
			return
		}

		// 获取用户信息
		var userID int
		cookie, err := c.Cookie(sessionKey)
		if err == nil {
			verifyUserID, err := jwtToken.Verify(cookie)
			if err != nil {
				utillog.Warnf("Verify cookie failed: %+v\n", err)
			} else {
				userID = verifyUserID
			}
		} else {
			utillog.Warnf("Get cookie failed: %+v\n", err)
		}

		// 这里要知道路由是不是文件上传/下载接口，然后将内容传递/返回给f
		var multipartReader *multipart.Reader
		if ho.isFile && method == http.MethodPost {
			multipartReader, err = c.Request.MultipartReader()
			if err != nil {
				c.JSON(http.StatusMethodNotAllowed, Result{Error: utilerrors.Error{
					Code: utilerrors.ErrorCodeRouter,
					Msg:  fmt.Sprintf("%+v", err),
				}})
				return
			}
		}

		// 注入上下文、用户和参数信息，并执行业务方法
		var r Result
		var statusCode = http.StatusOK
		p := Param{method: method, body: body, values: values, multipartReader: multipartReader}
		dbBase := models.NewBase()
		logger := utillog.Default()
		if ho.useTx {
			// 事务-统一从这里开启。ao和db不需要理会事务，只需要使用ctx.DB()返回的实例去操作即可
			// 即使是相同的请求，每次进来都会是一个新事务，所以基本上是没有事务嵌套的问题的
			err = dbBase.WithTx(func(tx db.DB) error {
				ctx := context.New(tx, logger, userID)

				r = f(ctx, p)

				return nil
			})
		} else {
			db := dbBase.DB()
			ctx := context.New(db, logger, userID)
			r = f(ctx, p)
			err = r.Err()
		}
		// 处理错误
		if e, ok := err.(utilerrors.Error); ok {
			if e.IsNormal() {
				statusCode = http.StatusBadRequest
			} else if e.IsFatal() {
				statusCode = http.StatusInternalServerError
			}
			r.Error = e
		} else {
			if err != nil {
				c.JSON(http.StatusForbidden, Result{Error: utilerrors.Error{
					Code: utilerrors.ErrorCodeRouter,
					Msg:  fmt.Sprintf("%+v", err),
				}})
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
			cookie, err := MakeCookie(r.CookieAfterLogin)
			if err != nil {
				c.JSON(http.StatusInternalServerError, Result{Error: utilerrors.Error{
					Code: utilerrors.ErrorCodeRouter,
					Msg:  fmt.Sprintf("%+v", err),
				}})
				return
			}
			c.Header(setCookieHeaderKey, cookie.String())
		}

		// 调用过滤器，过滤返回内容
		if v, ok := r.Data.(Filter); ok {
			r.Data = v.Filter()
		}

		// 返回文件内容
		if ho.isFile && method == http.MethodGet {
			c.DataFromReader(statusCode, r.ContentLength, r.ContentType, r.ContentReader, r.ExtraHeaders)
			return
		}

		// 返回
		c.JSON(statusCode, r)
	}
}

// MakeCookie 新建令牌
func MakeCookie(userID int) (cookie http.Cookie, err error) {
	session, err := jwtToken.Sign(userID)
	if err != nil {
		return
	}

	days := 7
	var maxAge = 3600 * 24 * days

	cookie.Name = sessionKey
	cookie.Value = session
	cookie.MaxAge = maxAge
	cookie.Expires = time.Now().AddDate(0, 0, days)
	cookie.Path = "/"
	cookie.HttpOnly = true

	return
}
