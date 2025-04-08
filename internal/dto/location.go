package dto

import (
	"github.com/google/uuid"
	"raspyx/internal/domain/models"
)

type CreateLocationRequest struct {
	Name string `json:"name" example:"Автозаводская" binding:"required"`
}

type CreateLocationResponse struct {
	UUID uuid.UUID `json:"uuid" example:"c555b9e8-0d7a-11f0-adcd-20114d2008d9"`
}

type GetLocationsResponse struct {
	Locations []models.Location `json:"locations" binding:"required"`
}

type UpdateLocationRequest struct {
	Name string `json:"name" example:"Автозаводская" binding:"required"`
}
