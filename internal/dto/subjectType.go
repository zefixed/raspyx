package dto

import (
	"github.com/google/uuid"
	"raspyx/internal/domain/models"
)

type CreateSubjectTypeRequest struct {
	Type string `json:"type" example:"Практика" binding:"required"`
}

type CreateSubjectTypeResponse struct {
	UUID uuid.UUID `json:"uuid" example:"c555b9e8-0d7a-11f0-adcd-20114d2008d9"`
}

type GetSubjectTypesResponse struct {
	SubjectTypes []models.SubjectType `json:"subjectTypes" binding:"required"`
}

type UpdateSubjectTypeRequest struct {
	Type string `json:"type" example:"Практика" binding:"required"`
}
