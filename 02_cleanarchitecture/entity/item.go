package entity

func NewItem() Item {
	return &item{}
}

type Item interface {
}

type item struct {
}
