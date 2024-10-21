package entity

import (
	"github.com/google/uuid"
)

type Model struct {
	ID              *uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v6();primaryKey" json:"id"`
	Name            string        `json:"name"`
	Title           *string       `json:"title"`
	SubTitle        *string       `json:"sub_title"`
	Description     *string       `json:"description"`
	IsNew           bool          `json:"is_new"`
	Encrypt         *string       `json:"encrypt"`
	Random          *string       `json:"random"`
	SettingsLink    *string       `json:"settings_link"`
	AltImage        *string       `json:"alt_image"`
	Position        *int          `json:"position"`
	Price           float64       `json:"price"`
	Slug            string        `json:"slug"`
	BodyType        *string       `json:"body_type"`
	CargoVolume     *string       `json:"cargo_volume"`
	EmissionCO2     *string       `json:"emission_co2"`
	FuelConsumption *string       `json:"fuel_consumption"`
	FuelType        *string       `json:"fuel_type"`
	VehicleEngine   *string       `json:"vehicle_engine"`
	SeatCapacity    *int          `json:"seat_capacity"`
	EngineType      *string       `json:"engine_type"`
	HybridSystem    *string       `json:"hybrid_system"`
	Displayable     bool          `json:"displayable"`
	Images          []MediaObject `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"images"` // Correct association
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
