package controller

import "github.com/labstack/echo"

func NewItem(g *echo.Group) Item {
	return &item{g: g}
}

type Item interface {
	HandleItem()
}

type item struct {
	g *echo.Group
}

func (i *item) HandleItem() {
	i.g.GET("/notices", listNotice)
}
