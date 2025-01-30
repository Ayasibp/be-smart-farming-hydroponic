package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Farm struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	ProfileId uuid.UUID      `json:"profile_id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name      string         `json:"name" gorm:"type:varchar;not null;unique"`
	Address   string         `json:"address" gorm:"type:varchar;not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
