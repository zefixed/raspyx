package models

import "github.com/google/uuid"

type Teacher struct {
	UUID       uuid.UUID
	FirstName  string
	SecondName string
	MiddleName string
}
