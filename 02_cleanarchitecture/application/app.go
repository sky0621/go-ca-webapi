package application

import "go-ca-webapi/02_cleanarchitecture/adapter"

func NewApp() *App {
	return &App{}
}

type App struct {
}

func (a *App) Start() {

	adapter.NewWebServer()
}
