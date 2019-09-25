package presenter

import (
	"github.com/labstack/echo"
	"go-ca-webapi/02_cleanarchitecture/usecase/outputport"
	"net/http"
)

func NewItem(c echo.Context) outputport.ItemOutputPort {
	return &item{c: c}
}

type item struct {
	c echo.Context
}

func sendResponse(c echo.Context, code int) error {
	return c.JSON(code, struct {
		Message string `json:"message"`
	}{Message: http.StatusText(code)})
}
