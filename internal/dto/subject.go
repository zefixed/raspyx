package dto

import (
	"github.com/google/uuid"
	"raspyx/internal/domain/models"
)

type CreateSubjectRequest struct {
	Name string `json:"name" example:"Иностранный язык" binding:"required"`
}

type CreateSubjectResponse struct {
	UUID uuid.UUID `json:"uuid" example:"c555b9e8-0d7a-11f0-adcd-20114d2008d9"`
}

type GetSubjectsResponse struct {
	Subjects []models.Subject `json:"subjects" binding:"required"`
}

type UpdateSubjectRequest struct {
	Name string `json:"name" example:"Иностранный язык" binding:"required"`
}
