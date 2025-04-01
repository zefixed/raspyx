package interfaces

import (
	"context"
	"github.com/google/uuid"
	"raspyx/internal/domain/models"
)

type SubjectTypeRepository interface {
	Create(ctx context.Context, subjectType *models.SubjectType) error
	Get(ctx context.Context) ([]*models.SubjectType, error)
	GetByUUID(ctx context.Context, uuid uuid.UUID) (*models.SubjectType, error)
	GetByType(ctx context.Context, subjectType string) (*models.SubjectType, error)
	Update(ctx context.Context, subjectType *models.SubjectType) error
	Delete(ctx context.Context, uuid uuid.UUID) error
}
