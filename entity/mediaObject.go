package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/datatypes"
	"time"
	"fmt"

	"github.com/adefilippi/go-demo-api/service/utils"
)

func init() {
	RegisterType("entity.MediaObject", func() interface{} {
		return &MediaObject{}
	})
}

type MediaObject struct {
	ID           *uuid.UUID     `gorm:"type:uuid;primaryKey" json:"id"` // Define custom UUID ID
	Title        *string        `json:"title"`
	Name         *string        `json:"name"`
	OriginalName *string        `json:"original_name"`
	Path         *string        `json:"path"`
	MimeType     *string        `json:"mime_type"`
	Tag          *string        `json:"tag"`
	Size         *int64         `json:"size"`
	Association  *string        `json:"association"`
	Dimensions   datatypes.JSON `gorm:"type:json" json:"dimensions"`
	ModelID      uuid.UUID      `json:"model_id"` // Foreign key for Model
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	Url          *string        `json:"url" gorm:"-"`
}

func (MediaObject) TableName() string {
	return "media_object" // Specify the singular form here
}

func (m *MediaObject) AfterFind(tx *gorm.DB) (err error) {
	m = setURL(m)
	return
}

func (m *MediaObject) AfterSave(tx *gorm.DB) (err error) {
	m = setURL(m)
	return
}

func (m *MediaObject) AfterDelete(tx *gorm.DB) (err error) {
	// Delete associated file if it exists
	err = utils.DeleteFile(*m.Name, *m.Association+"_"+utils.GetAssociationValueId(m, *m.Association))
	if err != nil {
		fmt.Println(err)
	}

	return
}

func setURL(m *MediaObject) *MediaObject {
	// Assuming URL is derived from the Path field, you can set the URL here
	fullUrl := "/" + *m.Name
	m.Url = &fullUrl
	return m
}
