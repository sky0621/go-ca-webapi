package driver

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DBコネクションクローズ用の関数
type closeDBConnectionFunc func()

// RDB(MySQL)コネクション取得
func NewDBConnection(user, pass, host, db string) (*gorm.DB, closeDBConnectionFunc, error) {
	connStr := "%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local"
	dbConn, err := gorm.Open("mysql", fmt.Sprintf(connStr, user, pass, host, db))
	return dbConn, func() {
		if dbConn != nil {
			if err := dbConn.Close(); err != nil {
				log.Fatal(err)
			}
		}
	}, err
}
