package dao

import (
	"watIwant/config"
	"watIwant/models"
)

type ItemDAO struct {
	database *config.Database
}


func NewItemDAO() ItemDAO{
	var db = config.GetDatabase()

	if !db.HasTable(models.Item{}){
		db.Table("item").CreateTable(models.Item{})
	}
	return ItemDAO{database:db}
}

func (itemDAO ItemDAO) ReadOne (itemId string) (models.Item, error){
	var item models.Item
	err := itemDAO.database.Find(&item).Where(&models.Item{Name: itemId}).Error
	return item, err
}
func (itemDAO ItemDAO) ReadAll() ([]models.Item, error) {
	var collection []models.Item
	err := itemDAO.database.Find(&collection).Error
	return collection, err
}