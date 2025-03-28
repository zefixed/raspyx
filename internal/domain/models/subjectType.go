package models

import "github.com/google/uuid"

type SubjectType struct {
	UUID uuid.UUID `json:"uuid"`
	Type string    `json:"type"`
}
