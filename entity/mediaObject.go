package entity

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type MediaObject struct {
	ID           *uuid.UUID     `gorm:"type:uuid;primaryKey" json:"id"` // Define custom UUID ID
	Title        *string        `json:"title"`
	Name         *string        `json:"name"`
	OriginalName *string        `json:"original_name"`
	MimeType     *string        `json:"mime_type"`
	Tag          *string        `json:"tag"`
	Size         *int64         `json:"size"`
	Association  *string        `json:"association"`
	Dimensions   datatypes.JSON `gorm:"type:json" json:"dimensions"`
	ModelID      uuid.UUID      `json:"model_id"` // Foreign key for Model
}

func (MediaObject) TableName() string {
	return "media_object" // Specify the singular form here
}
