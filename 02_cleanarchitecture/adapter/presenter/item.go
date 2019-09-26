package presenter

import (
	"github.com/labstack/echo"
	"go-ca-webapi/02_cleanarchitecture/usecase/outputport"
)

func NewItem(c echo.Context) outputport.ItemOutputPort {
	return &itemPresenter{c: c}
}

type itemPresenter struct {
	c echo.Context
}
