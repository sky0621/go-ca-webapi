package gateway

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go-ca-webapi/02_cleanarchitecture/domain/model"
	"go-ca-webapi/02_cleanarchitecture/domain/repository"
)

func NewItem(dbConn *gorm.DB) repository.Item {
	return &ItemRepository{dbConn: dbConn}
}

// entity.repository層の実装
type ItemRepository struct {
	dbConn *gorm.DB
}

func (i *ItemRepository) SaveItem(m *model.Item) error {
	if err := i.dbConn.Create(convertFrom(m)).Error; err != nil {
		return errors.Wrap(err, "@repository.itemRepository#SaveItem()")
	}
	return nil
}

func (i *ItemRepository) ListItem() ([]*model.Item, error) {
	var items []*itemRecord
	if err := i.dbConn.Find(&items).Error; err != nil {
		return nil, errors.Wrap(err, "@repository.itemRepository#SaveItem()")
	}

	var results []*model.Item
	for _, item := range items {
		results = append(results, item.convertToModel())
	}
	return results, nil
}

type itemRecord struct {
	ID    string `gorm:"column:id;primary_key"` // 商品ID
	Name  string `gorm:"column:name"`           // 商品名
	Price int    `gorm:"column:price"`          // 金額
}

// O/RマッパーにGormを使う上で必要となる「テーブル名」のマッピング
func (i *itemRecord) TableName() string {
	return "item"
}

func (i *itemRecord) convertToModel() *model.Item {
	return &model.Item{
		ID:    i.ID,
		Name:  i.Name,
		Price: i.Price,
	}
}

// entity層のモデルをGorm依存のモデルにマッピング
func convertFrom(itemModel *model.Item) *itemRecord {
	return &itemRecord{
		ID:    itemModel.ID,
		Name:  itemModel.Name,
		Price: itemModel.Price,
	}
}
