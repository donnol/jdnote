package app

import (
	stdctx "context"
	"fmt"
	"net/http"
	"time"

	"github.com/donnol/jdnote/utils/config"
	"github.com/donnol/jdnote/utils/context"
	"github.com/donnol/jdnote/utils/jwt"
	"github.com/donnol/jdnote/utils/route"
	"github.com/donnol/jdnote/utils/store/db"
	"github.com/donnol/tools/inject"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq" // github.com/lib/pq postgresql驱动
)

// DB DB接口
type DB = db.DB

type App struct {
	*Base

	config   config.Config
	db       DB
	jwtToken *jwt.Token
	ioc      *inject.Ioc
	router   *route.Router
}

func New(ctx stdctx.Context) (*App, context.Context) {
	app := &App{
		Base: NewBase(),
	}

	// 配置,来自文件或acm
	// defaultConfig 默认配置
	app.config = normal

	// 数据库: mysql或pg,redis
	// defaultDB 默认db
	app.db = func() *sqlx.DB {
		db, err := sqlx.Open(app.config.DB.Scheme, app.config.DB.String())
		if err != nil {
			panic(err)
		}
		if err := db.Ping(); err != nil {
			panic(err)
		}

		return db
	}()

	// ctx
	cusCtx := context.New(ctx, app.db, 0)

	// 第三方sdk: oss等

	//  其它
	// jwt
	app.jwtToken = jwt.New([]byte(app.config.JWT.Secret))

	// ioc
	app.ioc = inject.NewIoc(true)

	// defaultRouter 默认路由
	app.router = route.NewRouter(route.Option{
		SessionKey: sessionKey,
		JwtToken:   app.jwtToken,
	})

	return app, cusCtx
}

// session
const (
	sessionKey = "jd_session"
)

func (app *App) GetConfig() config.Config {
	return app.config
}

// MakeCookie 新建令牌
func (app *App) MakeCookie(userID int) (cookie http.Cookie, err error) {
	session, err := app.jwtToken.Sign(userID)
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

// Router 获取默认路由
func (app *App) Router() *route.Router {
	return app.router
}

func (app *App) MustRegisterProvider(vs ...interface{}) {
	for _, v := range vs {
		if err := app.ioc.RegisterProvider(v); err != nil {
			panic(err)
		}
	}
}

// Register 注册
func (app *App) Register(ctx context.Context, v interface{}) {
	// 初始化
	if err := app.ioc.Inject(v); err != nil {
		panic(err)
	}

	app.router.Register(ctx, v)
}

func (app *App) Run() {

}

// StartServer 开启服务
func (app *App) StartServer(port int) error {
	addr := fmt.Sprintf(":%d", port)
	if err := app.router.Run(addr); err != nil {
		return err
	}

	return nil
}
