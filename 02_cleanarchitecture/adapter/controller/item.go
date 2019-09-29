package controller

import (
	"github.com/labstack/echo"
	"go-ca-webapi/02_cleanarchitecture/adapter/presenter"
	"go-ca-webapi/02_cleanarchitecture/usecase"
	"net/http"
)

func NewItem(e *echo.Echo, input usecase.Item) Item {
	c := &itemController{
		e:     e,
		input: input,
	}
	c.handle()
	return c
}

type Item interface {
	handle()
}

type itemController struct {
	e     *echo.Echo
	input usecase.Item
}

func (i *itemController) handle() {
	i.e.POST("/item", i.saveItem)
	i.e.GET("/item", i.listItem)
}

// 「商品」を登録
func (i *itemController) saveItem(c echo.Context) error {
	r := &saveItemRequest{}
	if err := c.Bind(r); err != nil {
		return sendResponse(c, http.StatusBadRequest)
	}

	return i.input.SaveItem(convertFrom(r), presenter.NewItem(c))
}

// 「商品」一覧を返却
func (i *itemController) listItem(c echo.Context) error {
	return i.input.ListItem(presenter.NewItem(c))
}

func sendResponse(c echo.Context, code int) error {
	return c.JSON(code, struct {
		Message string `json:"message"`
	}{Message: http.StatusText(code)})
}

// JSON形式のHTTPリクエストBodyパース用
type saveItemRequest struct {
	Name  string `json:"name"`  // 商品名
	Price int    `json:"price"` // 金額
}

// HTTPリクエストをusecase層に渡すための変換
func convertFrom(r *saveItemRequest) *usecase.SaveItemRequest {
	return &usecase.SaveItemRequest{
		ID:    generateID(),
		Name:  r.Name,
		Price: r.Price,
	}
}
