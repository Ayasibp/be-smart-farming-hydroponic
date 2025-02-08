package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GrowthHist struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	FarmId    uuid.UUID      `json:"farm_id" gorm:"type:uuid;not null"`
	SystemId  uuid.UUID      `json:"system_id" gorm:"type:uuid;not null"`
	Ppm       float64        `json:"ppm" gorm:"type:float;not null"`
	Ph        float64        `json:"ph" gorm:"type:float;not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type GrowthHistFilter struct {
	Ppm       float64   `json:"ppm" gorm:"type:float;not null"`
	Ph        float64   `json:"ph" gorm:"type:float;not null"`
	CreatedAt time.Time `json:"created_at"`
}

type GrowthHistAggregate struct {
	TotalPpm  float64 `json:"totalPpm" gorm:"column:totalPpm;type:float;"`
	TotalPh   float64 `json:"totalPh" gorm:"column:totalPh;type:float;"`
	TotalData int64   `json:"totalData" gorm:"column:totalData;type:integer;"`
	MinPpm    float64 `json:"minPpm" gorm:"column:minPpm;type:float;"`
	MaxPpm    float64 `json:"maxPpm" gorm:"column:maxPpm;type:float;"`
	MinPh     float64 `json:"minPh" gorm:"column:minPh;type:float;"`
	MaxPh     float64 `json:"maxPh" gorm:"column:maxPh;type:float;"`
	AvgPpm    float64 `json:"avgPpm" gorm:"column:avgPpm;type:float;"`
	AvgPh     float64 `json:"avgPh" gorm:"column:avgPh;type:float;"`
}
