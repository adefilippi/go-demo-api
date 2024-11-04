package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gosimple/slug"

	"example/web-service-gin/entity"
)

func GetAllModels() ([]entity.Model, error) {
	if db == nil {
		fmt.Println("db is nil")
	}
	var models []entity.Model

	err := db.Model(&entity.Model{}).Preload("Images").Find(&models).Error // Use the correct field name for association
	return models, err
}

func GetModelById(id uuid.UUID) (entity.Model, error) {
	var model entity.Model
	if result := db.First(&model, id); result.Error != nil {
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
