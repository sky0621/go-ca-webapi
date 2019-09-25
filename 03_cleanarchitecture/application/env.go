package application

import "os"

func NewEnv() Env {
	return &env{
		webServerPort: os.Getenv("GO_CA_WEBAPI_SERVER_PORT"),
	}
}

type Env interface {
	WebServerPort() string
}

type env struct {
	webServerPort string
}

func (e *env) WebServerPort() string {
	return e.webServerPort
}
