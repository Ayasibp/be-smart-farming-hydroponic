package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GrowthPlan struct {
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	FarmId      uuid.UUID      `json:"farm_id" gorm:"type:uuid;not null"`
	SystemId    uuid.UUID      `json:"system_id" gorm:"type:uuid;not null"`
	SeedName    string         `json:"seed_name" gorm:"type:varchar;not null"`
	SeedSource  string         `json:"seed_source" gorm:"type:varchar;not null"`
	SeedQty     string         `json:"seed_qty" gorm:"type:int;not null"`
	StartPlan   time.Time      `json:"start_plan" gorm:"type:date;not null"`
	HarvestPlan time.Time      `json:"harvest_plan" gorm:"type:date;not null"`
	StartAct    time.Time      `json:"start_act" gorm:"type:date;not null"`
	HarvesAct   time.Time      `json:"harves_act" gorm:"type:date;not null"`
	HarvesQty   int            `json:"harvest_qty" gorm:"type:int;not null"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
