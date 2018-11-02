package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
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

func (controller ItemController) Post(ctx *gin.Context){
	var newItem models.Item
	var err error
	var uid string
	err = ctx.BindJSON(&newItem)

	if err != nil {
		log.Println(err)
	}

	uid, err = controller.itemDAO.Insert(newItem)
	if err != nil {
		ctx.JSON(500, err)
	}
	ctx.Header("Location", ctx.Request.RequestURI+"/"+uid)
	ctx.JSON(204, nil)

}