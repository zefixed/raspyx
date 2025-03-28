package models

import "github.com/google/uuid"

type Teacher struct {
	UUID       uuid.UUID `json:"uuid"`
	FirstName  string    `json:"first_name"`
	SecondName string    `json:"second_name"`
	MiddleName string    `json:"middle_name,omitempty"`
}
