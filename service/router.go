package service

import (
	"github.com/gin-gonic/gin"

	"example/web-service-gin/api"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/models", api.GetModels)
	router.GET("/models/:id", api.GetModelById)
	router.POST("/models", api.CreateModel)
	router.PATCH("/models/:id", api.UpdateModel)
	router.DELETE("/models/:id", api.DeleteModel)
	return router
}
