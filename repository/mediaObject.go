package repository

import (
	"fmt"
	"github.com/google/uuid"

	"example/web-service-gin/entity"
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

func GetMediaObjectById(id uuid.UUID) (entity.MediaObject, error) {
	var mediaObject entity.MediaObject
	if result := db.First(&mediaObject, id); result.Error != nil {
		return entity.MediaObject{}, result.Error
	}
	return mediaObject, nil
}

func DeleteMediaObject(id uuid.UUID) error {
	mediaObject, err := GetMediaObjectById(id)
	if err == nil {
		return err
	}

	transaction := db.Begin()

	if err := db.Delete(&mediaObject).Error; err != nil {
		transaction.Rollback()
		fmt.Println("Erreur ", err)
		return err
	}

	fmt.Println("mediaObject (Before Delete)", mediaObject)
	transaction.Commit()
	return nil
}
