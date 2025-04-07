package interfaces

import (
	"context"
	"github.com/google/uuid"
	"raspyx/internal/domain/models"
)

//go:generate go run github.com/vektra/mockery/v2@latest --name=GroupRepository
type GroupRepository interface {
	Create(ctx context.Context, group *models.Group) error
	Get(ctx context.Context) ([]*models.Group, error)
	GetByUUID(ctx context.Context, uuid uuid.UUID) (*models.Group, error)
	GetByNumber(ctx context.Context, number string) (*models.Group, error)
	Update(ctx context.Context, group *models.Group) error
	Delete(ctx context.Context, uuid uuid.UUID) error
}
