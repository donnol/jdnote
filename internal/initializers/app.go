package initializers

import (
	stdctx "context"
	"fmt"
	stdlog "log"
	"net/http"
	"os"
	"time"

	"github.com/donnol/jdnote/utils/cache"
	"github.com/donnol/jdnote/utils/config"
	"github.com/donnol/jdnote/utils/context"
	"github.com/donnol/jdnote/utils/jwt"
	"github.com/donnol/jdnote/utils/queue"
	"github.com/donnol/jdnote/utils/route"
	"github.com/donnol/jdnote/utils/store/db"
	"github.com/donnol/jdnote/utils/store/influx"
	"github.com/donnol/jdnote/utils/store/redis"
	"github.com/donnol/jdnote/utils/timer"

	"github.com/donnol/tools/inject"
	"github.com/donnol/tools/log"

	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/jmoiron/sqlx"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/robfig/cron/v3"

	_ "net/http/pprof"

	_ "github.com/lib/pq" // github.com/lib/pq postgresql驱动
)

type App struct {
	opt *Option

	config config.Config

	db              db.DB
	redisClient     *redis.Client
	influxdb        *influx.Client
	influxAPIWriter api.WriteAPI

	cache   cache.Cache
	logger  log.Logger
	trigger queue.Trigger

	ioc         *inject.Ioc
	arounderMap inject.ArounderMap
	proxy       inject.Proxy

	jwtToken *jwt.Token
	router   *route.Router
	server   *http.Server
	cron     *cron.Cron
}

const (
	ProjectEnv = "PROJECT_ENV"

	ProjectEnvDev = "PROJECT_ENV_DEV"

	ProjectEnvProd = "PROJECT_ENV_PROD"
)

func GetProjectEnv() string {
	if v, ok := os.LookupEnv(ProjectEnv); ok && v == ProjectEnvProd {
		return ProjectEnvProd
	}
	return ProjectEnvDev
}

func New(ctx stdctx.Context, setters ...OptionSetter) (*App, context.Context) {
	// 选项
	opt := &Option{}
	for _, setter := range setters {
		setter(opt)
	}

	// 检查必填项，如果没设置，或报错，或使用默认值
	if err := opt.checkRequire(); err != nil {
		panic(err)
	}

	// 新建app
	app := &App{}

	// 选项
	app.opt = opt

	// 配置,来自环境变量，如docker run时用-e指定
	// defaultConfig 默认配置
	switch GetProjectEnv() {
	case ProjectEnvProd:
		app.config = normal
	default:
		app.config = dev
	}
	fmt.Printf("config: %+v\n", app.config)

	// 数据库: mysql或pg
	var err error
	app.db, err = db.Open(db.Option{
		DriverName:     app.config.DB.Scheme,
		DataSourceName: app.config.DB.String(),
	})
	if err != nil {
		panic(err)
	}

	// redis
	app.redisClient = redis.NewClient(&redis.Options{
		Addr:     app.config.Redis.Addr,
		Password: app.config.Redis.Password,
	})

	// cache
	app.cache = cache.New(cache.Option{
		RedisClient: app.redisClient,
	})

	// logger
	app.logger = log.New(os.Stdout, "", stdlog.LstdFlags|stdlog.Lshortfile)

	// trigger
	app.trigger = queue.NewTrigger(queue.Option{
		RedisClient: app.redisClient,
	})

	// influxdb
	app.influxdb = influx.Open(influx.Option{
		Host:  app.config.InfluxDB.Host,
		Token: app.config.InfluxDB.Token,
	}, nil)
	app.influxAPIWriter = app.influxdb.WriteAPI(app.config.InfluxAPIWriter.OrgName, app.config.InfluxAPIWriter.BucketName)

	// ctx
	cusCtx := context.New(ctx, app.db)

	// 第三方sdk: oss等

	//  其它
	// jwt
	app.jwtToken = jwt.New([]byte(app.config.JWT.Secret))

	// ioc
	app.ioc = inject.NewIoc(true)

	// proxy
	app.proxy = inject.NewProxy()

	// defaultRouter 默认路由
	app.router = route.NewRouter(route.Option{
		SessionKey: sessionKey,
		JwtToken:   app.jwtToken,
	})

	// timer
	opts := []cron.Option{
		cron.WithLocation(time.Local),
	}
	app.cron = cron.New(opts...)

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

type ProviderOption struct {
	Provider interface{}
	Mock     interface{}
}

func (app *App) RegisterArounderMap(arounderMap map[inject.ProxyContext]inject.AroundFunc) {
	app.arounderMap = arounderMap
}

func (app *App) MustRegisterProvider(opts ...ProviderOption) {
	for _, opt := range opts {
		v := app.proxy.Around(opt.Provider, opt.Mock, GetArounder(app.arounderMap))
		if err := app.ioc.RegisterProvider(v); err != nil {
			panic(err)
		}
	}
}

func (app *App) MustInject(v interface{}) {
	if err := app.ioc.Inject(v); err != nil {
		panic(err)
	}
}

func (app *App) RegisterRouterWithInject(ctx context.Context, v interface{}) {
	app.MustInject(v)

	app.router.Register(ctx, v, route.RegisterOption{
		InfluxAPIWriter: app.influxAPIWriter,
		ReqTimeout:      app.opt.timeout,
	})
}

func (app *App) StaticFS(relativePath string, fs http.FileSystem) {
	app.router.StaticFS(relativePath, fs)
}

func (app *App) RegisterTimerHandler(spec string, f timer.FuncJob) {
	jobWrapper := cron.Recover(cron.DefaultLogger)
	job := jobWrapper(f)
	entryID, err := app.cron.AddJob(spec, job)
	if err != nil {
		panic(err)
	}
	app.logger.Infof("Cron AddJob: %v\n", entryID)
}

func (app *App) RunTimer() error {
	app.cron.Start()

	return nil
}

func (app *App) StopTimer() stdctx.Context {
	return app.cron.Stop()
}

func (app *App) RunPprof() error {
	addr := app.config.Pprof.Port.ToAddr()

	// 启动pprof
	go func() {
		app.logger.Debugf("Pprof server start: %s\n", addr)

		app.logger.Errorf("pprof ListenAndServe err: %+v\n", http.ListenAndServe(addr, nil))
	}()

	return nil
}

func (app *App) RunPrometheus() error {
	addr := app.config.Prometheus.Port.ToAddr()

	// 启动prometheus
	go func() {
		app.logger.Debugf("Prometheus server start: %s\n", addr)

		http.Handle("/metrics", promhttp.Handler())
		app.logger.Errorf("prometheus ListenAndServe err: %+v\n", http.ListenAndServe(addr, nil))
	}()

	return nil
}

func (app *App) Run() error {
	port := app.config.Server.Port

	if err := app.StartServer(port.Raw()); err != nil {
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
