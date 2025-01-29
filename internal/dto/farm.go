package dto

import "github.com/google/uuid"

type CreateFarm struct {
	ProfileID uuid.UUID `json:"profile_id" binding:"required"`
	Name      string    `json:"name" binding:"required"`
	Address   string    `json:"address" binding:"required"`
}

type FarmResponse struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name" binding:"required"`
	Address string    `json:"address" binding:"required"`
}

type UpdateFarm struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
}
