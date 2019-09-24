package adapter

import (
	"fmt"
	"github.com/labstack/echo"
	"go-ca-webapi/02_cleanarchitecture/application"
)

func NewWebServer(env application.Env) application.WebServer {
	return &webServer{env: env}
}

type webServer struct {
	env application.Env
}

func (w *webServer) Start() error {
	e := echo.New()

	return e.Start(fmt.Sprintf(":%s", w.env.WebServerPort()))
}
