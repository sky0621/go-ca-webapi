package application

func NewApp(env Env, webServer WebServer) *App {
	return &App{
		env:       env,
		webServer: webServer,
	}
}

type App struct {
	env       Env
	webServer WebServer
}

func (a *App) Start() error {

	return a.webServer.Start()
}
