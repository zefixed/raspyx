package models

import (
	"github.com/google/uuid"
)

type Location struct {
	UUID uuid.UUID `json:"uuid" example:"c555b9e8-0d7a-11f0-adcd-20114d2008d9"`
	Name string    `json:"name" example:"Автозаводская"`
}
