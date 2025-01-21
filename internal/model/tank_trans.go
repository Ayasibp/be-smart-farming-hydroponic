package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TankTran struct {
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	FarmId      uuid.UUID      `json:"farm_id" gorm:"type:uuid;not null"`
	SystemId    uuid.UUID      `json:"system_id" gorm:"type:uuid;not null"`
	WaterVolume int            `json:"water_volume" gorm:"type:int;not null"`
	AVolume     int            `json:"a_volume" gorm:"type:int;not null"`
	BVolume     int            `json:"b_volume" gorm:"type:int;not null"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
