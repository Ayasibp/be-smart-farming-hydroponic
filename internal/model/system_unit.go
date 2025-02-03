package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SystemUnit struct {
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	FarmId      uuid.UUID      `json:"farm_id" gorm:"type:uuid;not null"`
	UnitKey     uuid.UUID      `json:"unit_key" gorm:"type:uuid;not null"`
	TankVolume  int            `json:"tank_volume" gorm:"type:int;not null"`
	TankAVolume int            `json:"tank_a_volume" gorm:"type:int;not null"`
	TankBVolume int            `json:"tank_b_volume" gorm:"type:int;not null"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

type SystemUnitJoined struct {
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	FarmId      uuid.UUID      `json:"farm_id" gorm:"type:uuid;not null"`
	FarmName      string      `json:"farm_name" gorm:"type:varchar;not null"`
	UnitKey     uuid.UUID      `json:"unit_key" gorm:"type:uuid;not null"`
	TankVolume  int            `json:"tank_volume" gorm:"type:int;not null"`
	TankAVolume int            `json:"tank_a_volume" gorm:"type:int;not null"`
	TankBVolume int            `json:"tank_b_volume" gorm:"type:int;not null"`
}
