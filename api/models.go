package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/datatypes"

	"github.com/adefilippi/go-demo-api/entity"
	"github.com/adefilippi/go-demo-api/repository"
	"github.com/syneido/go-api-core/service/utils"
)

const ASSOCIATION string = "model"

// @Summary		Show all models
// @Description	get all models
// @Tags			Model
// @Accept			json
// @Produce		json
// @Success		200	{object}	entity.Model
// @Failure		400	{object}	ApiError
// @Failure		404	{object}	ApiError
// @Failure		500	{object}	ApiError
// @Router			/models [get]
func GetModels(c *gin.Context) {
	models, err := repository.GetAllModels(handleQuery(c))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.IndentedJSON(http.StatusOK, models)
}

// @Summary		Show an account
// @Description	get string by ID
// @Tags			Model
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"Model ID"
// @Success		200	{object}	entity.Model
// @Failure		400	{object}	ApiError
// @Failure		404	{object}	ApiError
// @Failure		500	{object}	ApiError
// @Router			/models/{id} [get]
func GetModelById(c *gin.Context) {
	model, _ := repository.GetModelById(handleQuery(c))
	c.IndentedJSON(http.StatusOK, model)
}

// @Summary		Add a model
// @Description	Add by json model
// @Tags			Model
// @Accept			json
// @Produce		json
// @Param			model	body		entity.Model	true	"Add account"
// @Success		200		{object}	entity.Model
// @Failure		400		{object}	ApiError
// @Failure		404		{object}	ApiError
// @Failure		500		{object}	ApiError
// @Router			/models [post]
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

// @Summary		Update a Model
// @Description	Update by json Model
// @Tags			Model
// @Accept			json
// @Produce		json
// @Param			id		path		string					true	"Model ID"
// @Param			model	body		entity.Model	true	"Update model"
// @Success		200		{object}	entity.Model
// @Failure		400		{object}	ApiError
// @Failure		404		{object}	ApiError
// @Failure		500		{object}	ApiError
// @Router			/models/{id} [patch]
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

// @Summary		Delete a model
// @Description	Delete by model ID
// @Tags			Model
// @Accept			json
// @Produce		json
// @Param			id	path	string	true	"Model ID"
// @Success		204	{object}	entity.Model
// @Failure		400	{object}	ApiError
// @Failure		404	{object}	ApiError
// @Failure		500	{object}	ApiError
// @Router			/models/{id} [delete]
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

// @Summary		Get a model image
// @Description	Get by model image ID
// @Tags			Model - image
// @Accept			json
// @Produce		application/octet-stream
// @Param			id	path	string	true	"Model ID"
// @Param			image-id	path	string	true	"Model Image ID"
// @Success 200 {file} file "Image or file"
// @Failure		400	{object}	ApiError
// @Failure		404	{object}	ApiError
// @Failure		500	{object}	ApiError
// @Router			/models/{id}/image/{image-id} [get]
func GetMdelImage(c *gin.Context) {
	_, error := repository.GetModelById(handleQuery(c))
	if error != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	image, error := repository.GetMediaObjectById(handleQuery(c))
	if error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	img, err := utils.GetFile(*image.Name, *image.Association+"_"+utils.GetAssociationValueId(image, *image.Association))
	if err != nil {
	}

	c.Data(http.StatusOK, *image.MimeType, img)
}

//		@Summary		Post a model image
//		@Description	Post by model image ID
//		@Tags			Model - image
//		@Accept			multipart/form-data
//		@Produce		json
//		@Param			id	path	string	true	"Model ID"
//	 @Param file formData file true "Image file"
//
// @Param tag formData string true "Tag for the image"
//
//	@Success		200	{object}	entity.MediaObject
//	@Failure		400	{object}	ApiError
//	@Failure		404	{object}	ApiError
//	@Failure		500	{object}	ApiError
//	@Router			/models/{id}/image [post]
func CreateModelImage(c *gin.Context) {

	model, error := repository.GetModelById(handleQuery(c))
	if model.ID == nil {
		c.IndentedJSON(http.StatusNotFound, "Model not found")
		return
	}

	modelId, error := uuid.Parse(c.Param("id"))
	if error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	filepath, filename := utils.HandleFile(c, "model_"+modelId.String())
	mimetype, err := utils.GetFileMimeType(filepath)
	if err != nil {
		c.AbortWithStatus(http.StatusServiceUnavailable)
	}
	size, err := utils.GetFileSize(filepath)
	if err != nil {
		c.AbortWithStatus(http.StatusServiceUnavailable)
	}

	association := ASSOCIATION
	mo := entity.MediaObject{
		ModelID:      *model.ID,
		Name:         &filename,
		OriginalName: &filename,
		MimeType:     &mimetype,
		Size:         &size,
		Association:  &association,
	}

	if utils.IsImageMimeType(mimetype) {
		width, height, err := utils.GetImageDimensions(filepath)
		dimensions := [2]int{width, height}
		if err == nil {
			dimensionsJSON, err := json.Marshal(dimensions)
			if err == nil {
				mo.Dimensions = datatypes.JSON(dimensionsJSON)
			}
		}
	}

	mo, error = repository.CreateMediaObject(mo)
	if error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.IndentedJSON(http.StatusCreated, mo)
}

// @Summary		Delete a model image
// @Description	Delete by  model image ID
// @Tags			Model - image
// @Accept			json
// @Produce		json
// @Param			id	path	string	true	"Model ID"
// @Param			image-id	path	string	true	"Model Image ID"
// @Success		204	{object}	entity.MediaObject
// @Failure		400	{object}	ApiError
// @Failure		404	{object}	ApiError
// @Failure		500	{object}	ApiError
// @Router			/models/{id}/image/{image-id} [delete]
func DeleteModelImage(c *gin.Context) {
	error := repository.DeleteMediaObject(handleQuery(c))
	if error != nil {
		code, messages := HandleError(error)
		c.IndentedJSON(code, messages)
	} else {
		c.AbortWithStatus(http.StatusNoContent)
	}
}
