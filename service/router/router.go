package router

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"example/web-service-gin/api"
	docs "example/web-service-gin/docs"
	"example/web-service-gin/middleware"
	"example/web-service-gin/service/utils"
	"example/web-service-gin/service/env"
)

var allowedOrigins = strings.Split(env.GetEnvVariable("CORS_ALLOW_ORIGIN"), ",")

func DefaultHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !utils.Contains(allowedOrigins, c.Request.Host) {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.Header("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PATCH, DELETE")
		c.Header("Access-control-allow-origin", c.Request.Header.Get("Origin"))
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Cache-control", "no-cache, private, max-age=0")
		c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
		c.Header("Content-type", "application/json")
		c.Header("Referrer-Policy", "strict-origin")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("X-Content-Type-Options", "nosniff")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		// Add more headers as needed
		c.Next()
	}
}

func SetupRouter() *gin.Engine {

	router := gin.Default()
	router.HandleMethodNotAllowed = true
	router.SetTrustedProxies([]string{"127.0.0.0/8", "10.0.0.0/8", "172.16.0.0/12", "192.168.0.0/16"})
	router.Use(DefaultHeadersMiddleware())

	router.GET("/models", api.GetModels)
	router.GET("/models/:id", api.GetModelById)
	router.POST("/models", middleware.SecurityMiddleware(), api.CreateModel)
	router.PATCH("/models/:id", middleware.SecurityMiddleware(), api.UpdateModel)
	router.DELETE("/models/:id", middleware.SecurityMiddleware(), api.DeleteModel)

	router.GET("/models/:id/image/:image-id", api.GetMdelImage)
	router.POST("/models/:id/image", middleware.SecurityMiddleware(), api.CreateModelImage)
	router.DELETE("/models/:id/image/:image-id", middleware.SecurityMiddleware(), api.DeleteModelImage)

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
