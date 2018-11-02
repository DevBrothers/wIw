package models

import (
	"github.com/jinzhu/gorm"
	"github.com/google/uuid"
)

type Item struct {
	BaseModel
	Name string `json:`
	ShortDescription string `json:`
}

func NewItemModel() *Item {
	var itemModel Item
	return &itemModel
}

func (item Item) BeforeCreate(scope *gorm.Scope) error{
	uid, err := uuid.NewUUID()

	if err != nil {
		return err
	}
	scope.SetColumn("UUID",uid.String())
	return nil
}
