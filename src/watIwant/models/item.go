package models

type Item struct {
	Name string `json:"name"`
}

func NewItemModel() *Item {
	var itemModel Item
	return &itemModel
}

func NewItemModelWithName(name string) *Item {
	itemModel := NewItemModel()
	itemModel.Name = name
	return itemModel
}
