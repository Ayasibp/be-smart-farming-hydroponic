package dto

import "github.com/google/uuid"

type CreateSystemUnit struct {
	FarmID uuid.UUID `json:"farm_id" binding:"required"`
	TankVolume      int    `json:"tank_volume" binding:"required"`
	TankAVolume   int    `json:"tank_a_volume" binding:"required"`
	TankBVolume   int    `json:"tank_b_volume" binding:"required"`
}

type SystemUnitResponse struct {
	ID uuid.UUID `json:"id" binding:"required"`
	TankVolume      int    `json:"tank_volume" binding:"required"`
	TankAVolume   int    `json:"tank_a_volume" binding:"required"`
	TankBVolume   int    `json:"tank_b_volume" binding:"required"`
}
