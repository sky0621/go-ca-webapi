package repository

import (
	"go-ca-webapi/02_cleanarchitecture/domain/model"
)

type Item interface {
	SaveItem(m *model.Item) error
	ListItem() ([]*model.Item, error)
}
