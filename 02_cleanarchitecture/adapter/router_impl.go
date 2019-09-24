package adapter

import (
	"github.com/labstack/echo"
	"go-ca-webapi/02_cleanarchitecture/application"
	"net/http"
)

func NewRouter(e *echo.Echo) application.Router {
	// FIXME:

	return &router{e: e}
}

type router struct {
	e *echo.Echo
}

func (r *router) SetUp() {
	// FIXME:

	http.Handle("/", r.e)

	// 個人の認証を要するWebAPI用のルート
	authGroup := r.e.Group("/api/v1")

}
