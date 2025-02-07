package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Aggregation struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	FarmId    uuid.UUID      `json:"farm_id" gorm:"type:uuid;not null"`
	SystemId  uuid.UUID      `json:"system_id" gorm:"type:uuid;not null"`
	Name      string         `json:"name" gorm:"type:varchar;not null"`
	Value     float64        `json:"value" gorm:"type:float6;not null"`
	TimeRange string         `json:"time_range" gorm:"type:varchar;not null"`
	Activity  string         `json:"activity" gorm:"type:varchar;not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
