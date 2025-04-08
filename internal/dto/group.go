package dto

import (
	"github.com/google/uuid"
	"raspyx/internal/domain/models"
)

type CreateGroupRequest struct {
	Group string `json:"group" example:"221-352" binding:"required"`
}

type CreateGroupResponse struct {
	UUID uuid.UUID `json:"uuid" example:"c555b9e8-0d7a-11f0-adcd-20114d2008d9"`
}

type GetGroupResponse struct {
	Group string `json:"group" example:"221-352" binding:"required"`
}

type GetGroupsResponse struct {
	Groups []models.Group `json:"groups" binding:"required"`
}

type UpdateGroupRequest struct {
	Group string `json:"group" example:"221-352" binding:"required"`
}
