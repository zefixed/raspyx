package dto

import "github.com/google/uuid"

type CreateGroupRequest struct {
	Group string `json:"group" binding:"required"`
}

type CreateGroupResponse struct {
	UUID uuid.UUID `json:"uuid"`
}
