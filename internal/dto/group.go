package dto

import (
	"github.com/google/uuid"
)

type CreateGroupRequest struct {
	Group string `json:"group" example:"221-352" binding:"required"`
}

type CreateGroupResponse struct {
	UUID uuid.UUID `json:"uuid" example:"954921ef-0d5c-11f0-91ca-20114d2008d9"`
}
