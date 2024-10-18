package entity

import (
	"github.com/google/uuid"
)

type Model struct {
	ID              *uuid.UUID `json:"id,omitempty"`
	Name            string     `json:"name"`
	Title           *string    `json:"title,omitempty"`
	SubTitle        *string    `json:"sub_title,omitempty"`
	Description     *string    `json:"description,omitempty"`
	IsNew           bool       `json:"is_new"`
	Encrypt         *string    `json:"encrypt,omitempty"`
	Random          *string    `json:"random,omitempty"`
	SettingsLink    *string    `json:"settings_link,omitempty"`
	AltImage        *string    `json:"alt_image,omitempty"`
	Position        *int       `json:"position,omitempty"`
	Price           float64    `json:"price"`
	Slug            string     `json:"slug"`
	BodyType        *string    `json:"body_type,omitempty"`
	CargoVolume     *string    `json:"cargo_volume,omitempty"`
	EmissionCO2     *string    `json:"emission_co2,omitempty"`
	FuelConsumption *string    `json:"fuel_consumption,omitempty"`
	FuelType        *string    `json:"fuel_type,omitempty"`
	VehicleEngine   *string    `json:"vehicle_engine,omitempty"`
	SeatCapacity    *int       `json:"seat_capacity,omitempty"`
	EngineType      *string    `json:"engine_type,omitempty"`
	HybridSystem    *string    `json:"hybrid_system,omitempty"`
	Displayable     bool       `json:"displayable"`
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
