package usecase

import (
	"go-ca-webapi/02_cleanarchitecture/entity"
	"go-ca-webapi/02_cleanarchitecture/usecase/inputport"
)

func NewItem(entity *entity.Item) inputport.ItemInputPort {
	return &item{entity: entity}
}

type item struct {
	entity *entity.Item
}

func (i *item) SaveItem(r *inputport.SaveItemRequest) error {

	return nil
}

func (i *item) ListItem() error {

	return nil
}
