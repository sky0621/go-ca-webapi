package controller

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func NewItem(e *echo.Echo) Item {
	return &item{
		e: e,
	}
}

type Item interface {
	Handle()
}

type item struct {
	e *echo.Echo
}

func (i *item) Handle() {
	i.e.POST("/item", saveItem)
	i.e.GET("/item", listItem)
}

// 「商品」を登録
func saveItem(c echo.Context) error {
	i := &item{}
	if err := c.Bind(i); err != nil {
		return sendResponse(c, http.StatusBadRequest)
	}
	if err := db.Create(&i).Error; err != nil {
		log.Println(err)
		return sendResponse(c, http.StatusInternalServerError)
	}
	return sendResponse(c, http.StatusOK)

}

// 「商品」一覧を返却
func listItem(c echo.Context) error {
	var res []*item
	if err := db.Find(&res).Error; err != nil {
		log.Println(err)
		return sendResponse(c, http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, res)
}

func sendResponse(c echo.Context, code int) error {
	return c.JSON(code, struct {
		Message string `json:"message"`
	}{Message: http.StatusText(code)})
}
