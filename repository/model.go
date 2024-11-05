package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gosimple/slug"

	"example/web-service-gin/entity"
	"example/web-service-gin/service/utils"
)

func GetAllModels(parameters map[string]interface{}) ([]entity.Model, error) {
	var models []entity.Model
	criterias := entity.Model{}
	var filters map[string]interface{}
	var ok bool

	if parameters["filters"] != nil {
		filters, ok = parameters["filters"].(map[string]interface{})
		if !ok {
			fmt.Println("expected filters to be of type map[string]interface{}, got ", parameters["filters"])
			return models, fmt.Errorf("expected filters to be of type map[string]interface{}, got %T", parameters["filters"])
		}
	}

	where, params := utils.FiltersToWhereQuery(filters, &criterias)
	err := db.Model(&entity.Model{}).Preload("Images").Limit(parameters["itemsPerPage"].(int)).Offset((parameters["page"].(int)-1)*parameters["itemsPerPage"].(int)).Where(where, params...).Find(&models).Error // Use the correct field name for association
	return models, err
}

func GetModelById(parameters map[string]interface{}) (entity.Model, error) {
	id := parameters["path"].(map[string]interface{})["id"]
	uid, err := utils.ParseId(id)
	if err != nil {
		return entity.Model{}, err
	}
	var model entity.Model
	if result := db.First(&model, uid); result.Error != nil {
		return entity.Model{}, result.Error
	}
	return model, nil
}

func CreateModel(model entity.Model) (entity.Model, error) {
	u := uuid.New()
	model.ID = &u
	model.Slug = slug.Make(model.Name)
	transaction := db.Begin()
	if err := db.Create(&model).Error; err != nil {
		transaction.Rollback()
		return entity.Model{}, err
	}

	transaction.Commit()
	return model, nil
}

func UpdateModel(model entity.Model) (entity.Model, error) {
	transaction := db.Begin()
	model.Slug = slug.Make(model.Name)
	if err := db.Save(&model).Error; err != nil {
		transaction.Rollback()
		return entity.Model{}, err
	}
	transaction.Commit()
	return model, nil
}

func DeleteModel(id uuid.UUID) error {
	var model entity.Model
	transaction := db.Begin()
	if err := db.Where("id = ?", id).Delete(&model).Error; err != nil {
		transaction.Rollback()
		return err
	}
	transaction.Commit()
	return nil
}
