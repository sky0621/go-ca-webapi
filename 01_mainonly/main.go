package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

// 注意: プロダクション品質ではありません。
func main() {
	/*
	 * DBコネクション取得等の初期セットアップ
	 */
	db, err := gorm.Open("mysql", "localuser:localpass@tcp(localhost:3306)/localdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if db != nil {
			if err := db.Close(); err != nil {
				log.Fatal(err)
			}
		}
	}()

	/*
	 * Web-APIサーバの起動とルーティング設定
	 */
	e := echo.New()

	// 「商品」を登録
	e.POST("/item", func(c echo.Context) error {
		i := &item{}
		if err := c.Bind(i); err != nil {
			return sendResponse(c, http.StatusBadRequest)
		}
		if err := db.Create(&i).Error; err != nil {
			log.Println(err)
			return sendResponse(c, http.StatusInternalServerError)
		}
		return sendResponse(c, http.StatusOK)
	})

	// 「商品」一覧を返却
	e.GET("/item", func(c echo.Context) error {
		var res []*item
		if err := db.Find(&res).Error; err != nil {
			log.Println(err)
			return sendResponse(c, http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, res)
	})

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

func sendResponse(c echo.Context, code int) error {
	return c.JSON(code, struct {
		Message string `json:"message"`
	}{Message: http.StatusText(code)})
}
