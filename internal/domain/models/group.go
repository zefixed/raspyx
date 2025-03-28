package models

import "github.com/google/uuid"

type Group struct {
	UUID   uuid.UUID `json:"uuid"`
	Number string    `json:"number"`
}
