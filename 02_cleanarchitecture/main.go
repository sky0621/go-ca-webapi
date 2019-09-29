package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"go-ca-webapi/02_cleanarchitecture/adapter/controller"
	"go-ca-webapi/02_cleanarchitecture/usecase"

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

	app := Initialize(dbConn, echo.New())
	app.Start()
}

func NewApp(
	e *echo.Echo,
	itemUsecase usecase.Item,
	itemController controller.Item,
) App {
	return App{
		e:              e,
		itemUsecase:    itemUsecase,
		itemController: itemController,
	}
}

type App struct {
	e              *echo.Echo
	itemUsecase    usecase.Item
	itemController controller.Item
}

func (a App) Start() {
	a.itemController.Handle()
	a.e.Logger.Fatal(a.e.Start(":8080"))
}
