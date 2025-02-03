package dto

import "github.com/google/uuid"

type GrowthHist struct {
	FarmId   uuid.UUID `json:"farm_id" binding:"required"`
	SystemId uuid.UUID `json:"system_id" binding:"required"`
	Ppm      int       `json:"ppm" binding:"required"`
	Ph       int       `json:"ph" binding:"required"`
}
type GrowthHistResponse struct {
	ID       uuid.UUID `json:"id" binding:"required"`
	FarmId   uuid.UUID `json:"farm_id" binding:"required"`
	SystemId uuid.UUID `json:"system_id" binding:"required"`
	Ppm      int       `json:"ppm" binding:"required"`
	Ph       int       `json:"ph" binding:"required"`
}
