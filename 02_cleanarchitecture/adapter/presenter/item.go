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
	return i.c.JSON(http.StatusOK, target)
}

func (i itemPresenter) RenderListResult(targets []*usecasemodel.ListItem) error {
	return sendResponse(i.c, http.StatusOK)
}

func (i itemPresenter) RenderFailure(err error) error {
	return sendResponse(i.c, http.StatusInternalServerError)
}
