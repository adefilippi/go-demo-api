package middleware

import (
	"github.com/adefilippi/go-demo-api/service/security"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SecurityMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if isValidApiKey, _ := security.CheckApiKey(c.GetHeader("X-API-Key")); !isValidApiKey {
			if isValidBearerToken, msg := security.CheckBearer(c.GetHeader("Authorization")); !isValidBearerToken {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "message": msg})
			}
		}
		// If the API key is valid, proceed with the request
		c.Next()
	}
}
