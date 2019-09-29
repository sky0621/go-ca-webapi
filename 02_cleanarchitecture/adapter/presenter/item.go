package presenter

import (
	"github.com/labstack/echo"
	usecasemodel "go-ca-webapi/02_cleanarchitecture/usecase/model"
	"go-ca-webapi/02_cleanarchitecture/usecase/outputport"
	"net/http"
)

func NewItem(c echo.Context) outputport.ItemOutputPort {
	return &itemPresenter{c: c}
}

type itemPresenter struct {
	c echo.Context
}

func (i itemPresenter) RenderSaveResult(target *usecasemodel.SaveItem) error {
	return i.c.JSON(http.StatusOK, convertFromSaveItem(target))
}

func (i itemPresenter) RenderListResult(targets []*usecasemodel.ListItem) error {
	return i.c.JSON(http.StatusOK, convertFromListItems(targets))
}

func (i itemPresenter) RenderFailure(err error) error {
	return sendResponse(i.c, http.StatusInternalServerError)
}

// JSON形式のHTTPレスポンスBody生成用
type saveItemResponse struct {
	ID string `json:"id"` // ID
}

// HTTPリクエストをクライアントに渡すための変換
func convertFromSaveItem(r *usecasemodel.SaveItem) *saveItemResponse {
	return &saveItemResponse{ID: r.ID,}
}

// JSON形式のHTTPレスポンスBody生成用
type listItemResponse struct {
	ID    string `json:"id"`    // ID
	Name  string `json:"name"`  // 商品名
	Price int    `json:"price"` // 金額
}

// HTTPリクエストをクライアントに渡すための変換
func convertFromListItems(listItems []*usecasemodel.ListItem) []*listItemResponse {
	if listItems == nil || len(listItems) == 0 {
		return []*listItemResponse{}
	}
	var res []*listItemResponse
	for _, listItem := range listItems {
		res = append(res, &listItemResponse{
			ID:    listItem.ID,
			Name:  listItem.Name,
			Price: listItem.Price,
		})
	}
	return res
}
