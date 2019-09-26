package presenter

import (
	"github.com/labstack/echo"
	"net/http"
)

func sendResponse(c echo.Context, code int) error {
	return c.JSON(code, struct {
		Message string `json:"message"`
	}{Message: http.StatusText(code)})
}
