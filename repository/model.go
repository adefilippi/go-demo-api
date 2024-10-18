package repository

import (
	"github.com/google/uuid"

	"example/web-service-gin/entity"
)

func GetAllModels() []entity.Model {
	var models []entity.Model
	//db.Model(&Model{}).Limit(10).Find(&Model{})
	db.Find(&models)

	return models
}

func GetModelById(id uuid.UUID) entity.Model {
	var model entity.Model
	db.First(&model, id)
	return model
}

func CreateModel(model entity.Model) (entity.Model, error) {
	u := uuid.New()
	model.ID = &u
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
