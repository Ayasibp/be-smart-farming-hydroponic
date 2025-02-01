package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SuperUser struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Username  string         `json:"username" gorm:"type:varchar;not null; unique"`
	Password  string         `json:"password" gorm:"type:varchar; not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
