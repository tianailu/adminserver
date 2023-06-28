package i18n

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestCN(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	req.Header.Set(HeaderAcceptLanguage, "zh-cn")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := func(c echo.Context) error {
		return c.String(http.StatusOK, T("hello"))
	}

	conf := Config{Filepath: "../../testdata/i18n"}

	mw := MiddlewareWithConfig(conf)
	r := mw(handler)(c)
	assert.NoError(t, r)
	assert.Equal(t, "吃了吗", rec.Body.String())

}

func TestEN(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := func(c echo.Context) error {
		return c.String(http.StatusOK, T("hello"))
	}

	conf := Config{Filepath: "../../testdata/i18n"}

	mw := MiddlewareWithConfig(conf)
	r := mw(handler)(c)
	assert.NoError(t, r)
	assert.Equal(t, "hi!", rec.Body.String())
}
