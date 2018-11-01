package controllers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"watIwant/dao"
	"watIwant/models"
)

type ItemController struct {
	itemCollection []*models.Item
	itemDAO dao.ItemDAO
}

func NewItemController() *ItemController  {
	var itemController ItemController

	itemController.itemDAO = dao.NewItemDAO()

	for i := 0; i < 10; i++ {
		itemController.itemCollection = append(itemController.itemCollection, models.NewItemModelWithName("Test "+strconv.Itoa(i)))
	}
	return &itemController
}

func (controller ItemController) Get(context *gin.Context){
	var itemCollection []models.Item

	itemCollection, err := controller.itemDAO.ReadAll()

	if err !=nil{
		context.JSON(500, err)
	}

	context.JSON(200, itemCollection)
}

func (controller ItemController) GetById(ctx *gin.Context){

	item_id, _ := strconv.Atoi(ctx.Param("item_id"))
	ctx.JSON(200,
		controller.itemCollection[item_id])
}