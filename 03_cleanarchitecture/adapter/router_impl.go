package adapter

import (
	"github.com/labstack/echo"
	"go-ca-webapi/02_cleanarchitecture/adapter/controller"
	"go-ca-webapi/02_cleanarchitecture/application"
	"net/http"
)

func NewRouter() application.Router {
	// FIXME:
	return &router{}
}

type router struct {
	item controller.Item
}

func (r *router) SetUpRouter() {
	// FIXME:

	http.Handle("/", e)

	// 個人の認証を要するWebAPI用のルート
	authGroup := e.Group("/api/v1")

	item := controller.NewItem(authGroup)

}
