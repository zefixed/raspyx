package interfaces

import (
	"context"
	"github.com/google/uuid"
	"raspyx/internal/domain/models"
)

type TeacherRepository interface {
	Create(ctx context.Context, teacher *models.Teacher) error
	GetByUUID(ctx context.Context, uuid uuid.UUID) (*models.Teacher, error)
	GetByFullName(ctx context.Context, fn string) (*models.Teacher, error)
	Update(ctx context.Context, teacher *models.Teacher) error
	Delete(ctx context.Context, uuid uuid.UUID) error
}
