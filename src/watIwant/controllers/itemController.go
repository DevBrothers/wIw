package controllers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"watIwant/models"
)

type ItemController struct {
	itemCollection []*models.ItemModel
}

func NewItemController() *ItemController  {
	var itemController ItemController

	for i := 0; i < 10; i++ {
		itemController.itemCollection = append(itemController.itemCollection, models.NewItemModelWithName("Test "+strconv.Itoa(i)))
	}
	return &itemController
}

func (controller ItemController) Get(context *gin.Context){
	context.JSON(200, controller.itemCollection)
}

func (controller ItemController) GetById(ctx *gin.Context){

	item_id, _ := strconv.Atoi(ctx.Param("item_id"))
	ctx.JSON(200,
		controller.itemCollection[item_id])
}