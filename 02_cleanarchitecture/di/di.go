//+build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"go-ca-webapi/02_cleanarchitecture/adapter/controller"
	"go-ca-webapi/02_cleanarchitecture/adapter/gateway"
	"go-ca-webapi/02_cleanarchitecture/domain"
	"go-ca-webapi/02_cleanarchitecture/usecase"
)

func Initialize(e *echo.Echo, dbConn *gorm.DB) string {
	//wire.Build(usecase.NewItem)
	//wire.Build(domain.NewItem)
	//wire.Build(
	//	controller.NewItem,
	//	gateway.NewItem(dbConn),
	//)

	return wire.Build(
		usecase.NewItem,
		domain.NewItem,
		controller.NewItem,
		gateway.NewItem,
	)
}
