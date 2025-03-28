package models

import (
	"github.com/google/uuid"
)

type Location struct {
	UUID uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
}
