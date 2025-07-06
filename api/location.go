package api

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/lunmy/go-demo-api/repository"
)

func GetLocationInfos(c *gin.Context) {
	locations, err := repository.GetAllLocations(handleQuery(c))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.IndentedJSON(http.StatusOK, locations)
}
