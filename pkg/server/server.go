package server

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tianailu/adminserver/api"
	"github.com/tianailu/adminserver/api/server"
	"github.com/tianailu/adminserver/config"
	"github.com/tianailu/adminserver/pkg/cors"
	"github.com/tianailu/adminserver/pkg/db/mysql"
	"github.com/tianailu/adminserver/pkg/i18n"
	"log"
	"net"
	"net/http"
	"os"
)

var (
	App *echo.Echo
)

type AdminServer struct {
	Host           string
	Port           string
	Scheme         string
	Mode           string
	PrintRoutes    string
	CertFile       string
	KeyFile        string
	AdminSecretKey string
}

func NewAdminServer() server.Server {
	as := &AdminServer{}
	return as
}

func (ad *AdminServer) Initialize() {
	var (
		args config.Args
	)
	args.Parse()
	settings, err := config.Parse(config.DefaultConfigType, args.ConfigPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// 日志
	logPath := settings.GetConfig("stdout_logger")["filepath"]
	logFile, logErr := os.OpenFile(logPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if logErr != nil {
		log.Printf("日志文件打开失败：%v \n", logErr)
		os.Exit(1)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	App = echo.New()
	App.DisableHTTP2 = true
	App.HTTPErrorHandler = ad.HTTPErrorHandler

	// Middleware
	App.Use(middleware.Logger())
	App.Use(middleware.Recover())
	// 跨域访问
	corsConfig := cors.CORSConfig{
		AllowCredentials: true,
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
	}

	App.Use(cors.CORSWithConfig(corsConfig))

	// 初始化mysql
	mysqlConf := settings.GetConfig("mysql")
	fmt.Printf("1:%s-2:%s-3:%s-4:%s-5:%s-6:%s-7:%s \n", mysqlConf["username"], mysqlConf["password"], mysqlConf["ip"], mysqlConf["port"], mysqlConf["db_name"],
		mysqlConf["conn_max_lifetime"], mysqlConf["conn_max_idle"])
	mysql.InitMysqlDB(mysqlConf["username"], mysqlConf["password"], mysqlConf["ip"], mysqlConf["port"], mysqlConf["db_name"],
		mysqlConf["conn_max_lifetime"], mysqlConf["conn_max_idle"])

	ad.registerRouter()
	ad.Mode = settings.ConfigEr.String("mode")
	ad.Scheme = settings.ConfigEr.String("scheme")
	ad.Host = settings.ConfigEr.String("host")
	ad.Port = settings.ConfigEr.String("port")
	ad.PrintRoutes = settings.ConfigEr.String("print_routes")
	ad.CertFile = settings.ConfigEr.String("cert_file")
	ad.KeyFile = settings.ConfigEr.String("cert_file")
	ad.AdminSecretKey = settings.GetConfig("auth")["admin_secret_key"]
}

func (ad *AdminServer) registerRouter() {
	admin := App.Group("/adm")
	admin1 := admin.Group("/v1")
	admin1.Use(middleware.JWT([]byte(ad.AdminSecretKey)))
	api.InitGroupAdminRouter(admin1)
	// Login route
	api.InitAdminRouter(App)
}

func (ad *AdminServer) Start() {
	//App.HideBanner = true
	// 打印routes
	if ad.PrintRoutes != "" && ad.Mode == "debug" {
		d, _ := json.MarshalIndent(App.Routes(), "", "  ")
		fmt.Println(string(d))
	}

	// 启动服务
	addr := func(port string) string {
		return net.JoinHostPort(ad.Host, ad.Port)
	}
	if ad.Scheme == "http" {
		App.Logger.Fatal(App.Start(addr("80")))
	} else {
		App.Logger.Fatal(App.StartTLS(addr("443"), ad.CertFile, ad.KeyFile))
	}
}

func (ad *AdminServer) HTTPErrorHandler(err error, c echo.Context) {
	// http错误处理函数
	var (
		code = http.StatusInternalServerError
		msg  interface{}
	)

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
		if he.Internal != nil {
			err = fmt.Errorf("%v, %v", err, he.Internal)
		}
	} else if c.Echo().Debug {
		msg = err.Error()
	} else {
		msg = http.StatusText(code)
	}
	if v, ok := msg.(string); ok {
		msg = echo.Map{"message": i18n.T(v)}
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead { // Issue #608
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, msg)
		}
		if err != nil {
			c.Logger().Error(err)
		}
	}
}
