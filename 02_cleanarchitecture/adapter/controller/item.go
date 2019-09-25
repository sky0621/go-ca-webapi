package controller

import (
	"github.com/labstack/echo"
	"go-ca-webapi/02_cleanarchitecture/usecase"
	"log"
	"net/http"
)

func NewItem(e *echo.Echo, input usecase.Item) Item {
	return &item{
		e:     e,
		input: input,
	}
}

type Item interface {
	Handle()
}

type item struct {
	e     *echo.Echo
	input usecase.Item
}

func (i *item) Handle() {
	i.e.POST("/item", i.saveItem)
	i.e.GET("/item", listItem)
}

// 「商品」を登録
func (i *item) saveItem(c echo.Context) error {
	r := &saveItemRequest{}
	if err := c.Bind(r); err != nil {
		return sendResponse(c, http.StatusBadRequest)
	}

	i.input.SaveItem(r.convert())

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

// 「商品」を定義
type saveItemRequest struct {
	ID    string `json:"id"`    // 商品ID
	Name  string `json:"name"`  // 商品名
	Price int    `json:"price"` // 金額
}

func (r *saveItemRequest) convert() *usecase.SaveItemRequest {
	return &usecase.SaveItemRequest{
		ID:    r.ID,
		Name:  r.Name,
		Price: r.Price,
	}
}
