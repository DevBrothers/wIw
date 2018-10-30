package apiserver

import (
	"github.com/gin-gonic/gin"
)

func Handlers() *gin.Engine{
	engine := gin.Default()

	apiV1 := engine.Group("/v1")
	{
		itemRoutes(apiV1)
	}
	return engine
}
