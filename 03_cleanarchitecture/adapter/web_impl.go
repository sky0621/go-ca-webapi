package adapter

import (
	"fmt"
	"github.com/labstack/echo"
	"go-ca-webapi/02_cleanarchitecture/application"
)

func NewWebServer(env application.Env, router application.Router) application.WebServer {
	return &webServer{
		env:    env,
		router: router,
	}
}

type webServer struct {
	env    application.Env
	router application.Router
}

func (w *webServer) Start() error {
	e := echo.New()

	// FIXME: middlewareのセッティング実装！

	w.router.SetUpRouter()

	return e.Start(fmt.Sprintf(":%s", w.env.WebServerPort()))
}
