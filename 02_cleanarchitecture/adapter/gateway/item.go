package gateway

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go-ca-webapi/02_cleanarchitecture/entity/model"
	"go-ca-webapi/02_cleanarchitecture/entity/repository"
)

func NewItem(dbConn *gorm.DB) repository.Item {
	return &item{dbConn: dbConn}
}

type item struct {
	dbConn *gorm.DB
}

func (i *item) SaveItem(m *model.Item) error {
	if err := i.dbConn.Create(&i).Error; err != nil {
		return errors.Wrap(err, "@repository.item#SaveItem()")
	}
	return nil
}
