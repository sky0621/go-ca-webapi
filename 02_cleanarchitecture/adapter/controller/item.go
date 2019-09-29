package controller

import (
	"github.com/labstack/echo"
	"go-ca-webapi/02_cleanarchitecture/adapter/presenter"
	"go-ca-webapi/02_cleanarchitecture/usecase"
)

func NewItem(e *echo.Echo, input usecase.Item) Item {
	return &itemController{
		e:     e,
		input: input,
	}
}

type Item interface {
	Handle()
}

type itemController struct {
	e     *echo.Echo
	input usecase.Item
}

func (i *itemController) Handle() {
	i.e.POST("/item", i.saveItem)
	i.e.GET("/item", i.listItem)
}

// 「商品」を登録
func (i *itemController) saveItem(c echo.Context) error {
	o := presenter.NewItem(c)
	r := &saveItemRequest{}
	if err := c.Bind(r); err != nil {
		return o.RenderFailure(err)
	}

	return i.input.SaveItem(convertFrom(r), o)
}

// 「商品」一覧を返却
func (i *itemController) listItem(c echo.Context) error {
	return i.input.ListItem(presenter.NewItem(c))
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
