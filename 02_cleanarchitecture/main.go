package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"go-ca-webapi/02_cleanarchitecture/adapter/controller"
	"go-ca-webapi/02_cleanarchitecture/driver"
	"log"
	"os"
)

// 注意: プロダクション品質ではありません。
func main() {
	/*
	 * DBコネクション取得等の初期セットアップ
	 */
	dbConn, closeFunc, err := driver.NewDBConnection(
		os.Getenv("PRJ_GROWUP_USERNAME"),
		os.Getenv("PRJ_GROWUP_PASSWORD"),
		os.Getenv("PRJ_GROWUP_INSTANCE"),
		os.Getenv("PRJ_GROWUP_DBNAME"))
	if err != nil {
		log.Fatal(err)
	}
	defer closeFunc()

	/*
	 * Web-APIサーバの起動とルーティング設定
	 */
	e := echo.New()

	// 「/item」へのCRUDをルーティング
	controller.NewItem().Handle(e)

	e.Logger.Fatal(e.Start(":8080"))
}

// 「商品」を定義
type item struct {
	ID    string `json:"id" gorm:"column:id;primary_key"` // 商品ID
	Name  string `json:"name" gorm:"column:name"`         // 商品名
	Price int    `json:"price" gorm:"column:price"`       // 金額
}

func (i *item) TableName() string {
	return "item"
}
