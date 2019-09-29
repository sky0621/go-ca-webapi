package usecase

import (
	"go-ca-webapi/02_cleanarchitecture/domain"
	"go-ca-webapi/02_cleanarchitecture/domain/model"
	usecasemodel "go-ca-webapi/02_cleanarchitecture/usecase/model"
	"go-ca-webapi/02_cleanarchitecture/usecase/outputport"
)

func NewItem(itemDomain domain.Item) Item {
	return &item{itemDomain: itemDomain}
}

// adapter/controller層から呼ばれるインプットポート
type Item interface {
	SaveItem(r *SaveItemRequest, o outputport.ItemOutputPort) error
	ListItem(o outputport.ItemOutputPort) error
}

type item struct {
	itemDomain domain.Item
}

func (i *item) SaveItem(r *SaveItemRequest, o outputport.ItemOutputPort) error {
	err := i.itemDomain.SaveItem(convertFrom(r))
	if err == nil {
		return o.RenderSaveResult(&usecasemodel.SaveItem{ID: r.ID})
	} else {
		return o.RenderFailure(err)
	}
}

func (i *item) ListItem(o outputport.ItemOutputPort) error {
	// FIXME:
	//results, err := i.itemDomain.ListItem()
	return nil
}

type SaveItemRequest struct {
	ID    string // 商品ID
	Name  string // 商品名
	Price int    // 金額
}

func convertFrom(r *SaveItemRequest) *model.Item {
	return &model.Item{
		ID:    r.ID,
		Name:  r.Name,
		Price: r.Price,
	}
}
