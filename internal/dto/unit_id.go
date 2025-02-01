package dto

import "github.com/google/uuid"

type UnitIdResponse struct {
	ID uuid.UUID `json:"id" binding:"required"`
}
