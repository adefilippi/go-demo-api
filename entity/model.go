package entity

import (
	"github.com/google/uuid"
	"time"
)

func init() {
	RegisterType("entity.Model", func() interface{} {
		return &Model{}
	})
}

type Model struct {
	ID              *uuid.UUID    `gorm:"type:uuid;primaryKey" json:"id,omitempty" filter:"id"`
	Name            string        `json:"name,omitempty" filter:"name"`
	Title           *string       `json:"title,omitempty" filter:"title"`
	SubTitle        *string       `json:"sub_title,omitempty"`
	Description     *string       `json:"description,omitempty"`
	IsNew           bool          `json:"is_new,omitempty" filter:"isNew"`
	Encrypt         *string       `json:"encrypt,omitempty"`
	Random          *string       `json:"random,omitempty"`
	SettingsLink    *string       `json:"settings_link,omitempty"`
	AltImage        *string       `json:"alt_image,omitempty"`
	Position        *int          `json:"position,omitempty" filter:"position"`
	Price           float64       `json:"price,omitempty" filter:"price"`
	Slug            string        `json:"slug,omitempty" filter:"slug"`
	BodyType        *string       `json:"body_type,omitempty"`
	CargoVolume     *string       `json:"cargo_volume,omitempty"`
	EmissionCO2     *string       `json:"emission_co2,omitempty"`
	FuelConsumption *string       `json:"fuel_consumption,omitempty"`
	FuelType        *string       `json:"fuel_type,omitempty"`
	VehicleEngine   *string       `json:"vehicle_engine,omitempty"`
	SeatCapacity    *int          `json:"seat_capacity,omitempty"`
	EngineType      *string       `json:"engine_type,omitempty"`
	HybridSystem    *string       `json:"hybrid_system,omitempty"`
	Displayable     bool          `json:"displayable,omitempty"`
	Images          []MediaObject `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"images,omitempty" groups:"model" not_selectable:"true"` // Correct association
	CreatedAt       time.Time     `json:"created_at,omitempty" filter:"createdAt"`
	UpdatedAt       time.Time     `json:"updated_at,omitempty" filter:"updatedAt"`
}

func (Model) TableName() string {
	return "model" // Specify the singular form here
}

func NewModel(name string, price float64, slug string) Model {
	return Model{
		ID:          nil,
		Name:        name,
		IsNew:       true,
		Price:       price,
		Slug:        slug,
		Displayable: true,
	}
}
