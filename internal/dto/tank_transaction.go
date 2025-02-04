package dto

import "github.com/google/uuid"

type TankTransaction struct {
	FarmId   uuid.UUID `json:"farm_id" binding:"required"`
	SystemId uuid.UUID `json:"system_id" binding:"required"`
	WaterVolume      int       `json:"water_volume" binding:"required"`
	AVolume       int       `json:"a_volume" binding:"required"`
	BVolume       int       `json:"b_volume" binding:"required"`
}

type TankTransactionResponse struct {
	ID       uuid.UUID `json:"id" binding:"required"`
	FarmId   uuid.UUID `json:"farm_id" binding:"required"`
	SystemId uuid.UUID `json:"system_id" binding:"required"`
	WaterVolume      int       `json:"water_volume" binding:"required"`
	AVolume       int       `json:"a_volume" binding:"required"`
	BVolume       int       `json:"b_volume" binding:"required"`
}