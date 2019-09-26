package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"go-ca-webapi/02_cleanarchitecture/adapter/controller"
	"go-ca-webapi/02_cleanarchitecture/adapter/gateway"
	"go-ca-webapi/02_cleanarchitecture/driver"
	"go-ca-webapi/02_cleanarchitecture/entity"
	"go-ca-webapi/02_cleanarchitecture/usecase"
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

	// MEMO: カスタムロガー、クラウドサービス接続クライアント、その他外部APIクライアント等の生成もここで。

	/*
	 * Web-APIサーバの起動とルーティング設定
	 */
	e := echo.New()

	/*
	 * DI（後ほど、google/wire にて取りまとめる
	 */
	itemRepository := gateway.NewItem(dbConn)
	itemEntity := entity.NewItem(itemRepository)
	itemUsecase := usecase.NewItem(itemEntity)
	itemController := controller.NewItem(e, itemUsecase)

	// FIXME:
	//itemPresenter := presenter.NewItem()

	// 「/item」へのCRUDをルーティング
	itemController.Handle()

	e.Logger.Fatal(e.Start(":8080"))
}
