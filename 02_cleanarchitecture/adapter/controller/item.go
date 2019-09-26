package controller

import (
	"github.com/labstack/echo"
	"go-ca-webapi/02_cleanarchitecture/usecase"
	"net/http"
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
	r := &saveItemRequest{}
	if err := c.Bind(r); err != nil {
		return sendResponse(c, http.StatusBadRequest)
	}

	// adapter層(controller)はusecase層を呼ぶだけ（レスポンスはusecase層が(output-portを実装した)adapter層(controller)を呼ぶことで実現）
	i.input.SaveItem(convertFrom(r))
	return nil
}

// 「商品」一覧を返却
func (i *itemController) listItem(c echo.Context) error {
	// adapter層(controller)はusecase層を呼ぶだけ（レスポンスはusecase層が(output-portを実装した)adapter層(controller)を呼ぶことで実現）
	i.input.ListItem()
	return nil
}

func sendResponse(c echo.Context, code int) error {
	return c.JSON(code, struct {
		Message string `json:"message"`
	}{Message: http.StatusText(code)})
}

// JSON形式のHTTPリクエストBodyパース用
type saveItemRequest struct {
	ID    string `json:"id"`    // 商品ID
	Name  string `json:"name"`  // 商品名
	Price int    `json:"price"` // 金額
}

// HTTPリクエストをusecase層に渡すための変換
func convertFrom(r *saveItemRequest) *usecase.SaveItemRequest {
	return &usecase.SaveItemRequest{
		ID:    r.ID,
		Name:  r.Name,
		Price: r.Price,
	}
}
