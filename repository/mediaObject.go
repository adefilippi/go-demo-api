package repository

import (
	"fmt"
	"github.com/google/uuid"

	"github.com/adefilippi/go-demo-api/entity"
	"github.com/adefilippi/go-demo-api/service/utils"
)

func CreateMediaObject(mediaObject entity.MediaObject) (entity.MediaObject, error) {
	u := uuid.New()
	mediaObject.ID = &u
	transaction := db.Begin()
	if err := db.Create(&mediaObject).Error; err != nil {
		transaction.Rollback()
		return entity.MediaObject{}, err
	}

	transaction.Commit()
	return mediaObject, nil
}

func GetMediaObjectById(parameters map[string]interface{}) (entity.MediaObject, error) {

	id := parameters["path"].(map[string]interface{})["image-id"]
	uid, err := utils.ParseId(id)
	if err != nil {
		return entity.MediaObject{}, err
	}

	var mediaObject entity.MediaObject
	if result := db.First(&mediaObject, uid); result.Error != nil {
		return entity.MediaObject{}, result.Error
	}
	return mediaObject, nil
}

func DeleteMediaObject(parameters map[string]interface{}) error {
	mediaObject, err := GetMediaObjectById(parameters)
	if err == nil {
		return err
	}
	transaction := db.Begin()
	if err := db.Delete(&mediaObject).Error; err != nil {
		transaction.Rollback()
		fmt.Println("Erreur ", err)
		return err
	}

	transaction.Commit()
	return nil
}
