package auth

import (
	"crypto/subtle"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	// JWTConfig defines the config for JWT middleware.
	JWTConfig struct {
		// Skipper defines a function to skip middleware.
		Skipper middleware.Skipper

		// BeforeFunc defines a function which is executed just before the middleware.
		BeforeFunc middleware.BeforeFunc

		// SuccessHandler defines a function which is executed for a valid token.
		SuccessHandler JWTSuccessHandler

		// ErrorHandler defines a function which is executed for an invalid token.
		// It may be used to define a custom JWT error.
		ErrorHandler JWTErrorHandler

		// map[SigningMethod]SigningKey
		Signing map[string]interface{}

		// Context key to store user information from the token into context.
		// Optional. Default value "user".
		ContextKey string

		// Claims are extendable claims data defining token content.
		// Optional. Default value jwt.MapClaims
		Claims jwt.Claims

		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup string

		// AuthScheme to be used in the Authorization header.
		// Optional. Default value "Bearer".
		AuthScheme string

		keyFunc jwt.Keyfunc

		// Dual_Cip to enable anyone cip
		DualCip bool
	}

	// JWTSuccessHandler defines a function which is executed for a valid token.
	JWTSuccessHandler func(echo.Context)

	// JWTErrorHandler defines a function which is executed for an invalid token.
	JWTErrorHandler func(error) error

	jwtExtractor func(echo.Context) (string, error)

	// used to generate SigningKey
	SigningKeyFunc func() (interface{}, error)
)

// Algorithms
const (
	AlgorithmHS256 = "HS256"
	AlgorithmRS256 = "RS256"
)

// Errors
var (
	ErrJWTMissing = echo.NewHTTPError(http.StatusBadRequest, "missing or malformed jwt")
	ErrJWTInvalid = echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired jwt")
	secretKey     interface{}
	Valid         = ValidAuth
)

var (
	// DefaultJWTConfig is the default JWT auth middleware config.
	// claims:
	//     iat: token创建时间
	//     exp: token到期时间
	//     nbf: token有效开始时间
	//
	//     iss: token创建者(名称)
	//     sip: token创建者(ip)
	//     sub: token使用者(用户名)
	//     cip: token使用者(ip)
	//     role: 角色权限管理(TODO)
	DefaultJWTConfig = JWTConfig{
		Skipper:     middleware.DefaultSkipper,
		Signing:     map[string]interface{}{AlgorithmHS256: nil},
		ContextKey:  "user",
		TokenLookup: "header:" + echo.HeaderAuthorization,
		AuthScheme:  "Bearer",
		Claims:      jwt.MapClaims{},
	}
)

// JWT returns a JSON Web Token (JWT) auth middleware.
//
// For valid token, it sets the user in context and calls next handler.
// For invalid token, it returns "401 - Unauthorized" error.
// For missing token, it returns "400 - Bad Request" error.
//
// See: https://jwt.io/introduction
// See `JWTConfig.TokenLookup`
func JWT(key interface{}) echo.MiddlewareFunc {
	c := DefaultJWTConfig
	c.Signing[AlgorithmHS256] = key
	secretKey = key
	return JWTWithConfig(c)
}

func GetSecretKey() interface{} {
	return secretKey
}

// JWTWithConfig returns a JWT auth middleware with config.
// See: `JWT()`.
func JWTWithConfig(config JWTConfig) echo.MiddlewareFunc {
	// Defaults
	if config.Skipper == nil {
		config.Skipper = DefaultJWTConfig.Skipper
	}
	if config.Signing == nil {
		panic("echo: jwt middleware requires signing info")
	}
	for k, v := range config.Signing {
		if k == "" {
			panic("echo: jwt middleware get invalid signing method")
		}
		if v == nil {
			panic("echo: jwt middleware requires signing key")
		}
		if k == AlgorithmHS256 {
			secretKey = v
		}
	}
	if config.ContextKey == "" {
		config.ContextKey = DefaultJWTConfig.ContextKey
	}
	if config.Claims == nil {
		config.Claims = DefaultJWTConfig.Claims
	}
	if config.TokenLookup == "" {
		config.TokenLookup = DefaultJWTConfig.TokenLookup
	}
	if config.AuthScheme == "" {
		config.AuthScheme = DefaultJWTConfig.AuthScheme
	}
	config.keyFunc = func(t *jwt.Token) (interface{}, error) {
		// Check the signing method
		if key, ok := config.Signing[t.Method.Alg()]; ok {
			switch key.(type) {
			case SigningKeyFunc:
				r, err := key.(SigningKeyFunc)()
				if err != nil {
					return nil, err
				}
				return r, nil
			default:
				return key, nil
			}
		}
		return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
	}

	// Initialize
	parts := strings.Split(config.TokenLookup, ":")
	extractor := jwtFromHeader(parts[1], config.AuthScheme)
	switch parts[0] {
	case "query":
		extractor = jwtFromQuery(parts[1])
	case "cookie":
		extractor = jwtFromCookie(parts[1])
	case "header|session":
		extractor = jwtFromHeaderAndCookie(parts[1], config.AuthScheme)

	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			if config.BeforeFunc != nil {
				config.BeforeFunc(c)
			}

			auth, err := extractor(c)
			if err != nil {
				if config.ErrorHandler != nil {
					return config.ErrorHandler(err)
				}
				return err
			}
			token := new(jwt.Token)
			isMapClaims := false
			// Issue #647, #656
			if _, ok := config.Claims.(jwt.MapClaims); ok {
				isMapClaims = true
				token, err = jwt.Parse(auth, config.keyFunc)
			} else {
				t := reflect.ValueOf(config.Claims).Type().Elem()
				claims := reflect.New(t).Interface().(jwt.Claims)
				token, err = jwt.ParseWithClaims(auth, claims, config.keyFunc)
			}
			if err == nil && token.Valid {
				// Store user information from token into context.
				if isMapClaims {
					mClaims := token.Claims.(jwt.MapClaims)
					if sub, ok := mClaims["sub"]; ok {
						c.Set("username", sub.(string))
					}
					if tid, ok := mClaims["tenancy_key"]; ok {
						c.Set("tenancy_id", tid.(string))
					}
					if err = Valid(token.Raw, mClaims, config.DualCip, c); err == nil {
						c.Set(config.ContextKey, token)
						if config.SuccessHandler != nil {
							config.SuccessHandler(c)
						}
						return next(c)
					}
				} else {
					c.Set(config.ContextKey, token)
					if config.SuccessHandler != nil {
						config.SuccessHandler(c)
					}
					return next(c)
				}
			}
			if config.ErrorHandler != nil {
				return config.ErrorHandler(err)
			}
			return &echo.HTTPError{
				Code:     http.StatusUnauthorized,
				Message:  "invalid or expired jwt",
				Internal: err,
			}
		}
	}
}

// jwtFromHeader returns a `jwtExtractor` that extracts token from the request header.
func jwtFromHeader(header string, authScheme string) jwtExtractor {
	return func(c echo.Context) (string, error) {
		auth := c.Request().Header.Get(header)
		if strings.HasPrefix(auth, authScheme) {
			l := len(authScheme)
			if len(auth) > l+1 && auth[:l] == authScheme {
				return auth[l+1:], nil
			}
		} else if len(auth) > 0 {
			return auth, nil
		}
		return "", ErrJWTMissing
	}
}

// jwtFromQuery returns a `jwtExtractor` that extracts token from the query string.
func jwtFromQuery(param string) jwtExtractor {
	return func(c echo.Context) (string, error) {
		token := c.QueryParam(param)
		if token == "" {
			return "", ErrJWTMissing
		}
		return token, nil
	}
}

// jwtFromCookie returns a `jwtExtractor` that extracts token from the named cookie.
func jwtFromCookie(name string) jwtExtractor {
	return func(c echo.Context) (string, error) {
		cookie, err := c.Cookie(name)
		if err != nil {
			return "", ErrJWTMissing
		}
		return cookie.Value, nil
	}
}

// jwtFromHeader returns a `jwtExtractor` that extracts token from the request header or session cookie.
func jwtFromHeaderAndCookie(header string, authScheme string) jwtExtractor {
	return func(c echo.Context) (string, error) {
		auth := c.Request().Header.Get(header)
		if strings.HasPrefix(auth, authScheme) {
			l := len(authScheme)
			if len(auth) > l+1 && auth[:l] == authScheme {
				return auth[l+1:], nil
			}
		} else if len(auth) > 0 {
			return auth, nil
		}

		cookie, err := c.Cookie("session")
		if err != nil {
			return "", ErrJWTMissing
		}

		return cookie.Value, nil
	}
}

// 弃用: 请使用自定义设置.
// Deprecated: do not use default valid func.
func valid(raw string, m jwt.MapClaims, dc bool, c echo.Context) error {
	// 验证sip
	vErr := new(jwt.ValidationError)
	sip, _ := m["sip"].(string)
	if VerifyString(sip, c.Request().Host, false) == false {
		vErr.Inner = fmt.Errorf("token: invalid server ip")
		return vErr
	}
	// 验证cip
	cip, _ := m["cip"].(string)
	if VerifyString(cip, c.RealIP(), false) == false {
		vErr.Inner = fmt.Errorf("token: invalid client ip")
		return vErr
	}
	return nil
}

func ValidAuth(raw string, m jwt.MapClaims, dc bool, c echo.Context) error {
	// 验证sip
	vErr := new(jwt.ValidationError)
	sip, _ := m["sip"].(string)
	if VerifyString(sip, c.Request().Host, false) == false {
		vErr.Inner = fmt.Errorf("token: invalid server ip")
		return vErr
	}

	// 验证cip [dual_cip = false]
	if !dc {
		cip, _ := m["cip"].(string)
		if VerifyString(cip, c.Request().Header.Get(echo.HeaderXRealIP), false) == false {
			vErr.Inner = fmt.Errorf("token: invalid client ip")
			return vErr
		}
	}

	//审计认证只验证token有效性,不验证是否在TokenList中
	if strings.HasPrefix(c.Path(), "/admin/v1/audit") {
		return nil
	}

	return nil
}

// 比较两个字符串，相同或字符串为空返回true，否则返回false
func VerifyString(s string, cmp string, required bool) bool {
	if s == "" {
		return !required
	}
	if subtle.ConstantTimeCompare([]byte(s), []byte(cmp)) != 0 {
		return true
	} else {
		return false
	}
}
