package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"

	"example/web-service-gin/entity"
	"example/web-service-gin/repository"
)

func GetModels(c *gin.Context) {
	var models []entity.Model = repository.GetAllModels()
	c.IndentedJSON(http.StatusOK, models)
}

func GetModelById(c *gin.Context) {
	uuid, error := uuid.Parse(c.Param("id"))
	if error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	var model entity.Model = repository.GetModelById(uuid)
	c.IndentedJSON(http.StatusOK, model)
}

func CreateModel(c *gin.Context) {
	var model entity.Model
	c.BindJSON(&model)
	result, error := repository.CreateModel(model)
	if error != nil {

		code, messages := HandleError(error)
		c.IndentedJSON(code, messages)
	} else {
		c.IndentedJSON(http.StatusCreated, result)
	}
}

func UpdateModel(c *gin.Context) {
	var model entity.Model
	c.BindJSON(&model)
	result, error := repository.UpdateModel(model)
	if error != nil {
		code, messages := HandleError(error)
		c.IndentedJSON(code, messages)
	} else {
		c.IndentedJSON(http.StatusOK, result)
	}
}

func DeleteModel(c *gin.Context) {
	uuid, error := uuid.Parse(c.Param("id"))
	if error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	error = repository.DeleteModel(uuid)
	if error != nil {
		code, messages := HandleError(error)
		c.IndentedJSON(code, messages)
	} else {
		c.AbortWithStatus(http.StatusNoContent)
	}
}
