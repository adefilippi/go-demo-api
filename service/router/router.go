package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/lunmy/go-demo-api/api"
	docs "github.com/lunmy/go-demo-api/docs"
	"github.com/lunmy/go-demo-api/handler"
	"github.com/lunmy/go-api-core/middleware"
)

func SetupRouter(modelHandler *handler.ModelHandler, mediaObjectHandler *handler.MediaObjectHandler) *gin.Engine {
	router := gin.Default()

	router.HandleMethodNotAllowed = true
	router.SetTrustedProxies([]string{"127.0.0.0/8", "10.0.0.0/8", "172.16.0.0/12", "192.168.0.0/16"})
	router.Use(middleware.DefaultHeadersMiddleware())

	router.GET("/health-check", api.Home)
	router.GET("/models", modelHandler.GetAll)
	router.GET("/models/:id", modelHandler.GetOne)
	router.POST("/models", middleware.SecurityMiddleware(), modelHandler.Create)
	router.PATCH("/models/:id", middleware.SecurityMiddleware(), modelHandler.Update)
	router.DELETE("/models/:id", middleware.SecurityMiddleware(), modelHandler.Delete)

	router.GET("/models/:id/image/:image-id", mediaObjectHandler.GetOne)
	router.POST("/models/:id/image", middleware.SecurityMiddleware(), mediaObjectHandler.Create)
	router.DELETE("/models/:id/image/:image-id", middleware.SecurityMiddleware(), mediaObjectHandler.Delete)

	router.GET("/location/:codeCE", api.GetLocationInfos)

	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	return router
}