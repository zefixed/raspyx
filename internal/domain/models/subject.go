package models

import "github.com/google/uuid"

type Subject struct {
	UUID uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
}
