package apiserver

import (
	"github.com/gin-gonic/gin"
	"watIwant/controllers"
)

func itemRoutes(parentGroup *gin.RouterGroup)  {
	items := parentGroup.Group("/item")
	{
		items.GET("", controllers.NewItemController().Get)
		//items.POST("", controllers.NewItemController().Post)

		item:=items.Group(":item_id")
		{
			item.GET("", controllers.NewItemController().GetById)
		}

	}
}
