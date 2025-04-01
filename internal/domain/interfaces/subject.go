package interfaces

import (
	"context"
	"github.com/google/uuid"
	"raspyx/internal/domain/models"
)

type SubjectRepository interface {
	Create(ctx context.Context, subject *models.Subject) error
	Get(ctx context.Context) ([]*models.Subject, error)
	GetByUUID(ctx context.Context, uuid uuid.UUID) (*models.Subject, error)
	GetByName(ctx context.Context, name string) ([]*models.Subject, error)
	Update(ctx context.Context, subject *models.Subject) error
	Delete(ctx context.Context, uuid uuid.UUID) error
}
