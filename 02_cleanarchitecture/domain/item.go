package domain

import (
	"go-ca-webapi/02_cleanarchitecture/domain/model"
	"go-ca-webapi/02_cleanarchitecture/domain/repository"
)

func NewItem(itemRepository repository.Item) Item {
	return &item{itemRepository: itemRepository}
}

type Item interface {
	SaveItem(m *model.Item) error
	ListItem() ([]*model.Item, error)
}

type item struct {
	itemRepository repository.Item
}

func (i *item) SaveItem(m *model.Item) error {
	return i.itemRepository.SaveItem(m)
}

func (i *item) ListItem() ([]*model.Item, error) {
	return i.itemRepository.ListItem()
}
