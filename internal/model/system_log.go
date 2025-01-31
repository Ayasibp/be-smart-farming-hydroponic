package model

import (
	"time"

	"github.com/google/uuid"
)

type SystemLog struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Message   string            `json:"message" gorm:"type:varchar;not null"`
	CreatedAt time.Time      `json:"created_at"`
	
}
