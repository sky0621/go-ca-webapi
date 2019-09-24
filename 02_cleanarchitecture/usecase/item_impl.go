package usecase

import (
	"go-ca-webapi/02_cleanarchitecture/usecase/inputport"
	"go-ca-webapi/02_cleanarchitecture/usecase/outputport"
)

func NewItem() Item {
	return &item{}
}

type item struct {
}

func (i *item) AddItem(in inputport.Item) error {

	return nil
}

func (i *item) ListItem() ([]outputport.Item, error) {

	return nil, nil
}
