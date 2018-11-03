package models

type Item struct {
	BaseModel
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
}

/*func NewItemModel() *Item {
	var itemModel Item
	return &itemModel
}*/
