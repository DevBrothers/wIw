package models

type ItemModel struct {
	Name string `json:"name"`
}

func NewItemModel() *ItemModel{
	var itemModel ItemModel
	return &itemModel
}

func NewItemModelWithName(name string) *ItemModel{
	itemModel := NewItemModel()
	itemModel.Name = name
	return itemModel
}
