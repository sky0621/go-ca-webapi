package usecase

import (
	"go-ca-webapi/02_cleanarchitecture/usecase/inputport"
	"go-ca-webapi/02_cleanarchitecture/usecase/outputport"
)

type Item interface {
	// 新規ユーザを追加するユースケース
	AddItem(in inputport.Item) error

	// 既存ユーザの一覧を取得するユースケース
	ListItem() ([]outputport.Item, error)
}
