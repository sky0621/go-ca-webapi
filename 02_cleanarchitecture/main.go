package main

import (
	"github.com/labstack/echo"
	"go-ca-webapi/02_cleanarchitecture/driver"
	"log"
	"net/http"
	"os"
)

// 注意: プロダクション品質ではありません。
func main() {
	/*
	 * セットアップ
	 */
	adminKey := os.Getenv("GO_CA_WEBAPI_ADMIN_KEY")
	if adminKey == "" {
		log.Fatal("env GO_CA_WEBAPI_ADMIN_KEY doesn't exist")
	}

	// RDB(MySQL)コネクション取得
	db, closeFunc, err := driver.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer closeFunc()

	// Webサーバセットアップ
	e := echo.New()

	/*
	 * WebAPIルーティング設定
	 */
	// 「商品」を登録
	e.POST("/item", func(c echo.Context) error {
		i := &item{}
		// 本来用途と違うが「管理者のみ商品登録可能」の要件を満たすために暫定実装
		if c.Request().Header.Get("WWW-Authenticate") != adminKey {
			return sendResponse(c, http.StatusUnauthorized)
		}
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
			return err
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
