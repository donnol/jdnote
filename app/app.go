package app

import (
	stdctx "context"
	"fmt"
	stdlog "log"
	"net/http"
	"os"
	"time"

	"github.com/donnol/jdnote/utils/config"
	"github.com/donnol/jdnote/utils/context"
	"github.com/donnol/jdnote/utils/jwt"
	"github.com/donnol/jdnote/utils/queue"
	"github.com/donnol/jdnote/utils/route"
	"github.com/donnol/jdnote/utils/store/db"
	"github.com/donnol/tools/inject"
	"github.com/donnol/tools/log"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq" // github.com/lib/pq postgresql驱动
)

type App struct {
	*Base

	config   config.Config
	db       db.DB
	logger   log.Logger
	trigger  queue.Trigger
	jwtToken *jwt.Token
	ioc      *inject.Ioc
	router   *route.Router
	server   *http.Server
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

	// logger
	app.logger = log.New(os.Stdout, "", stdlog.LstdFlags|stdlog.Lshortfile)

	// trigger
	app.trigger = queue.NewTrigger(queue.Option{})

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

func (app *App) Logger() log.Logger {
	return app.logger
}

func (app *App) Trigger() queue.Trigger {
	return app.trigger
}

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

func (app *App) StaticFS(relativePath string, fs http.FileSystem) {
	app.router.StaticFS(relativePath, fs)
}

func (app *App) Run() error {
	port := app.config.Server.Port

	if err := app.StartServer(port); err != nil {
		return err
	}

	return nil
}

// StartServer 开启服务
func (app *App) StartServer(port int) error {
	app.logger.Debugf("Server start at %v. Listening '%v'", time.Now().Format("2006-01-02 15:04:05"), port)

	addr := fmt.Sprintf(":%d", port)
	app.server = &http.Server{
		Addr:    addr,
		Handler: app.router,
	}
	if err := app.server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func (app *App) ShutdownServer(ctx stdctx.Context) error {
	return app.server.Shutdown(ctx)
}

func (app *App) Cancel() {
	if idb, ok := app.db.(*sqlx.DB); ok {
		idb.Close()
	}
}
