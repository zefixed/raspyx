package dto

import (
	"github.com/google/uuid"
	"raspyx/internal/domain/models"
)

type CreateRoomRequest struct {
	Number string `json:"number" example:"ав4805" binding:"required"`
}

type CreateRoomResponse struct {
	UUID uuid.UUID `json:"uuid" example:"c555b9e8-0d7a-11f0-adcd-20114d2008d9"`
}

type GetRoomsResponse struct {
	Rooms []models.Room `json:"rooms" binding:"required"`
}

type UpdateRoomRequest struct {
	Number string `json:"number" example:"ав4805" binding:"required"`
}
