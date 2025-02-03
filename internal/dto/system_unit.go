package dto

import "github.com/google/uuid"

type CreateSystemUnit struct {
	FarmID      uuid.UUID `json:"farm_id" binding:"required"`
	UnitKey     uuid.UUID `json:"unit_key" binding:"required"`
	TankVolume  int       `json:"tank_volume" binding:"required"`
	TankAVolume int       `json:"tank_a_volume" binding:"required"`
	TankBVolume int       `json:"tank_b_volume" binding:"required"`
}

type CreateSystemUnitResponse struct {
	ID          uuid.UUID `json:"id" binding:"required"`
	TankVolume  int       `json:"tank_volume" binding:"required"`
	TankAVolume int       `json:"tank_a_volume" binding:"required"`
	TankBVolume int       `json:"tank_b_volume" binding:"required"`
}
type SystemUnitResponse struct {
	ID          uuid.UUID `json:"id" binding:"required"`
	UnitKey     uuid.UUID `json:"unit_key" binding:"required"`
	FarmID      uuid.UUID `json:"farm_id" binding:"required"`
	FarmName      string `json:"farm_name" binding:"required"`
	TankVolume  int       `json:"tank_volume" binding:"required"`
	TankAVolume int       `json:"tank_a_volume" binding:"required"`
	TankBVolume int       `json:"tank_b_volume" binding:"required"`
}
