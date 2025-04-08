package models

import "github.com/google/uuid"

type Teacher struct {
	UUID       uuid.UUID `json:"uuid" example:"c555b9e8-0d7a-11f0-adcd-20114d2008d9"`
	FirstName  string    `json:"first_name" example:"Имя"`
	SecondName string    `json:"second_name" example:"Фамилия"`
	MiddleName string    `json:"middle_name,omitempty" example:"Отчество"`
}
