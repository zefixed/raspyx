package interfaces

import (
	"context"
	"github.com/google/uuid"
	"raspyx/internal/domain/models"
)

type GroupRepository interface {
	Create(ctx context.Context, group *models.Group) error
	GetByUUID(ctx context.Context, uuid uuid.UUID) (*models.Group, error)
	GetByNumber(ctx context.Context, number string) (*models.Group, error)
	Update(ctx context.Context, group *models.Group) error
	Delete(ctx context.Context, uuid uuid.UUID) error
}
