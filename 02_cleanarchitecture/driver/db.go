package driver

import (
	"github.com/jinzhu/gorm"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DBコネクションクローズ用の関数
type closeDBConnectionFunc func()

// RDB(MySQL)コネクション取得
func NewDBConnection() (*gorm.DB, closeDBConnectionFunc, error) {
	db, err := gorm.Open("mysql", "localuser:localpass@tcp(localhost:3306)/localdb?charset=utf8&parseTime=True&loc=Local")
	return db, func() {
		if db != nil {
			if err := db.Close(); err != nil {
				log.Fatal(err)
			}
		}
	}, err
}
