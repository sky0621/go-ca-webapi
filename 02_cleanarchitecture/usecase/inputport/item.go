package inputport

type ItemInputPort interface {
	SaveItem(r *SaveItemRequest) error
	ListItem() error
}

type SaveItemRequest struct {
	ID    string // 商品ID
	Name  string // 商品名
	Price int    // 金額
}
