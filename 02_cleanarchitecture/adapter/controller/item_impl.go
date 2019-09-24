package controller

import (
	"github.com/labstack/echo"
	"go-ca-webapi/02_cleanarchitecture/usecase/inputport"
)

//func NewItem(input inputport.Item) Item {
//	return &item{input: input}
//}
func NewItem(g echo.Group) Item {
	return &item{}
}

type item struct {
	input inputport.Item
}

func (i *item) GetItemByID(c echo.Context) error {
	c.Request()
}
