package i18n

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	goi18n "github.com/nicksnyder/go-i18n/i18n"
	"github.com/tianailu/adminserver/pkg/utility/path"
)

type (
	Config struct {
		Skipper  middleware.Skipper
		Filepath string
		Langs    []string // TODO暂时没用
	}
)

const (
	DefaultFilepath      string = "/var/i18n"
	DefaultLanguage      string = "en"
	HeaderAcceptLanguage string = "Accept-Language"
)

var (
	DefaultConfig = Config{
		Filepath: DefaultFilepath,
		Skipper:  middleware.DefaultSkipper, // TODO需要吗
	}
	tfunc goi18n.TranslateFunc
)

func T(tid string, args ...interface{}) string {
	if tfunc == nil {
		return tid
	}
	return tfunc(tid, args...)
}

func SetT(lan string, lans ...string) {
	t, err := goi18n.Tfunc(lan, lans...)
	if err == nil {
		tfunc = t
	}
}

func Middleware() echo.MiddlewareFunc {
	return MiddlewareWithConfig(DefaultConfig)
}

func MiddlewareWithConfig(config Config) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = DefaultConfig.Skipper
	}
	if config.Filepath == "" {
		config.Filepath = DefaultFilepath
	}
	// 判断路径是否存在
	// 遍历目录获取下面的所有文件
	// 初始化i18n
	// 翻译文件名zh-CN 或 zh_cn不区分大小写
	ok, _ := path.IsDir(config.Filepath)
	if ok {
		files, _ := path.ListFiles(config.Filepath, false)
		for _, f := range files {
			goi18n.MustLoadTranslationFile(f)
		}
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}
			if lan := c.Request().Header.Get(HeaderAcceptLanguage); lan != "" {
				// TODO 增加根据cookie加载语言信息
				// header zh-CN,zh;q=0.9, fr-CH, fr;q=0.9, en;q=0.8, de;
				SetT(lan, DefaultLanguage)
			} else {
				// 如果不存在该header,按英文处理
				SetT(DefaultLanguage)
			}
			return next(c)
		}
	}
}
