package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"os"
	"strings"
)

const (
	defaultRsaPubKeyFile = "/tmp/rest-rsa.pub"
)

// 用于生成单点登录token验签公钥
var ReadPubKey SigningKeyFunc = func() (interface{}, error) {
	f, err := os.Open(defaultRsaPubKeyFile)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return jwt.ParseRSAPublicKeyFromPEM(data)
}

// 跳过的路由
func Skipper(c echo.Context) bool {

	if c.Request().Method == echo.OPTIONS {
		return true
	}

	if strings.HasPrefix(c.Path(), "verify") || strings.HasPrefix(c.Path(), "/login") {
		return true
	}

	return false
}
