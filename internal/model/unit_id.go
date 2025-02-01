package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UnitId struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
