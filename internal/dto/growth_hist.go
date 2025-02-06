package dto

import (
	"time"

	"github.com/google/uuid"
)

type GrowthHist struct {
	FarmId   uuid.UUID `json:"farm_id" binding:"required"`
	SystemId uuid.UUID `json:"system_id" binding:"required"`
	Ppm      float64   `json:"ppm" binding:"required"`
	Ph       float64   `json:"ph" binding:"required"`
}
type GrowthHistResponse struct {
	ID       uuid.UUID `json:"id" binding:"required"`
	FarmId   uuid.UUID `json:"farm_id" binding:"required"`
	SystemId uuid.UUID `json:"system_id" binding:"required"`
	Ppm      float64   `json:"ppm" binding:"required"`
	Ph       float64   `json:"ph" binding:"required"`
}
type GetGrowthFilter struct {
	Period   string `json:"period" binding:"required"`
	StartDate time.Time `json:"start_date" binding:"required"`
	EndDate time.Time `json:"end_date" binding:"required"`
}