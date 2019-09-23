package adapter

import (
	"github.com/labstack/echo"
	"go-ca-webapi/02_cleanarchitecture/application"
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
}
