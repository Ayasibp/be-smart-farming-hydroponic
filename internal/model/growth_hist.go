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
