package dto

import (
	"github.com/google/uuid"
)

type CreateTeacherRequest struct {
	FirstName  string `json:"first_name" example:"Имя" binding:"required"`
	SecondName string `json:"second_name" example:"Фамилия" binding:"required"`
	MiddleName string `json:"middle_name" example:"Отчество"`
}

type CreateTeacherResponse struct {
	UUID uuid.UUID `json:"uuid" example:"c555b9e8-0d7a-11f0-adcd-20114d2008d9"`
}

type TeacherDTO struct {
	Name string `json:"name" example:"Фамилия Имя Отчество" binding:"required"`
}

type GetTeachersResponse struct {
	Teachers []TeacherDTO `json:"teachers" binding:"required"`
}

type UpdateTeacherRequest struct {
	FirstName  string `json:"first_name" example:"Имя" binding:"required"`
	SecondName string `json:"second_name" example:"Фамилия" binding:"required"`
	MiddleName string `json:"middle_name" example:"Отчество"`
}
