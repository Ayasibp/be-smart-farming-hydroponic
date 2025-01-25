package dto

import "github.com/google/uuid"

type CreateProfile struct {
	AccountID uuid.UUID `json:"account_id" binding:"required"`
	Name      string    `json:"name" binding:"required"`
	Address   string    `json:"address" binding:"required"`
}

type ProfileResponse struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name" binding:"required"`
	Address string    `json:"address" binding:"required"`
}
