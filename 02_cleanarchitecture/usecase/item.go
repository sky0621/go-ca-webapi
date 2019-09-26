package usecase

import (
	"go-ca-webapi/02_cleanarchitecture/entity"
	"go-ca-webapi/02_cleanarchitecture/entity/model"
)

func NewItem(itemEntity entity.Item) Item {
	return &item{itemEntity: itemEntity}
}

// adapter/controller層から呼ばれるインプットポート
type Item interface {
	SaveItem(r *SaveItemRequest)
	ListItem()
}

type item struct {
	itemEntity entity.Item
}

func (i *item) SaveItem(r *SaveItemRequest) {
	err := i.itemEntity.SaveItem(convertFrom(r))
	// FIXME: output-portの呼び出し！
}

func (i *item) ListItem() {
	// FIXME:
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
