package adapter

import (
	"github.com/labstack/echo"
	"go-ca-webapi/02_cleanarchitecture/application"
)

func NewWebServer() application.WebServer {
	return &webServer{}
}

type webServer struct {
}

func (w *webServer) Start() error {
	e := echo.New()

	e.Logger.Fatal(e.Start(":8080"))
	return nil
}