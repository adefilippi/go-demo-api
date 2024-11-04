package entity

import (
	"github.com/google/uuid"
	"time"
)

type Model struct {
	ID              *uuid.UUID    `gorm:"type:uuid;primaryKey" json:"id,omitempty" groups:"model"`
	Name            string        `json:"name,omitempty" groups:"model"`
	Title           *string       `json:"title,omitempty" groups:"model"`
	SubTitle        *string       `json:"sub_title,omitempty" groups:"model"`
	Description     *string       `json:"description,omitempty" groups:"model, model:write"`
	IsNew           bool          `json:"is_new,omitempty"`
	Encrypt         *string       `json:"encrypt,omitempty"`
	Random          *string       `json:"random,omitempty"`
	SettingsLink    *string       `json:"settings_link,omitempty"`
	AltImage        *string       `json:"alt_image,omitempty"`
	Position        *int          `json:"position,omitempty"`
	Price           float64       `json:"price,omitempty"`
	Slug            string        `json:"slug,omitempty" group:"model"`
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
	CreatedAt       time.Time     `json:"created_at,omitempty" groups:"timestampable"`
	UpdatedAt       time.Time     `json:"updated_at,omitempty" groups:"timestampable"`
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
