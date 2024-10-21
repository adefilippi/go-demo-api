package middleware

import (
	"strings"
	"github.com/gin-gonic/gin"
	"net/http"
	"example/web-service-gin/service/env"
	"example/web-service-gin/service/utils"
	"fmt"
)

var allowedOrigins = strings.Split(env.GetEnvVariable("CORS_ALLOW_ORIGIN"), ",")

func DefaultHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Host : ", c.Request.Host)
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
