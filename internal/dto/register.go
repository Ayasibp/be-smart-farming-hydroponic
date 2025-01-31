package dto

import (
	"github.com/google/uuid"
)

type RegisterBody struct {
	UserName string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type RegisterResponse struct {
	UserID   uuid.UUID        `json:"user_id"`
	Username string           `json:"username"`
	Role     string           `json:"role"`
	ProfileResponse *ProfileResponse `json:"profile_response"`
}
