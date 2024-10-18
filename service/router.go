package service

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"example/web-service-gin/api"
	"example/web-service-gin/service/env"
)

var allowedOrigins = strings.Split(env.GetEnvVariable("CORS_ALLOW_ORIGIN"), ",")

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func SetDefaultHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !contains(allowedOrigins, c.Request.Host) {
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
	router.Use(SetDefaultHeaders())

	router.GET("/models", api.GetModels)
	router.GET("/models/:id", api.GetModelById)
	router.POST("/models", api.CreateModel)
	router.PATCH("/models/:id", api.UpdateModel)
	router.DELETE("/models/:id", api.DeleteModel)

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
