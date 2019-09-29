//+build wireinject

package main

import (
	"context"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"go-ca-webapi/02_cleanarchitecture/adapter/controller"
	"go-ca-webapi/02_cleanarchitecture/adapter/gateway"
	"go-ca-webapi/02_cleanarchitecture/domain"
	"go-ca-webapi/02_cleanarchitecture/usecase"
)

var superSet = wire.NewSet(
	gateway.NewItem,
	controller.NewItem,
	domain.NewItem,
	usecase.NewItem,
	NewApp,
)

func Initialize(ctx context.Context, dbConn *gorm.DB, e *echo.Echo) App {
	wire.Build(superSet)
	return App{}
}
