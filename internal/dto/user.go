package dto

import (
	"github.com/google/uuid"
	"raspyx/internal/domain/models"
)

type RegisterUserRequest struct {
	Username string `json:"username" example:"username" binding:"required"`
	Password string `json:"password" example:"password" binding:"required"`
}

type RegisterUserResponse struct {
	UUID uuid.UUID `json:"uuid" example:"c555b9e8-0d7a-11f0-adcd-20114d2008d9"`
}

type LoginUserRequest struct {
	Username string `json:"username" example:"username" binding:"required"`
	Password string `json:"password" example:"password" binding:"required"`
}
type LoginUserResponse struct {
	Token     string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDQwMzk0NjcsInJvbGUiOiJ1c2VyIiwic3ViIjoidXNlcm5hbWUifQ.Y_ql2u5Ir8EMJ_InCd-cO_OM4CFWFAEzJQzG_1-D3yE" binding:"required"`
	TokenType string `json:"token_type" example:"bearer" binding:"required"`
}

type GetUsersResponse struct {
	Users []models.User `json:"users" binding:"required"`
}

type UpdateUserRequest struct {
	AccessLevel int `json:"access_level" example:"0" binding:"required"`
}

type UserDTO struct {
	UUID        uuid.UUID `json:"uuid" example:"c555b9e8-0d7a-11f0-adcd-20114d2008d9"`
	Username    string    `json:"username" example:"username" binding:"required"`
	AccessLevel int       `json:"access_level" example:"0" binding:"required"`
}
