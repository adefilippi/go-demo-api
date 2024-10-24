package repository

import (
	"github.com/google/uuid"

	"example/web-service-gin/entity"
)

func CreateMediaObject(mediaObject entity.MediaObject) (entity.MediaObject, error) {
	u := uuid.New()
	association := "toto"
	mediaObject.Association = &association
	mediaObject.ID = &u
	transaction := db.Begin()
	if err := db.Create(&mediaObject).Error; err != nil {
		transaction.Rollback()
		return entity.MediaObject{}, err
	}

	transaction.Commit()
	return mediaObject, nil
}
