package models

import "github.com/google/uuid"

type Room struct {
	UUID   uuid.UUID `json:"uuid"`
	Number string    `json:"number"`
}
