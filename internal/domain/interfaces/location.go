package interfaces

import (
	"context"
	"github.com/google/uuid"
	"raspyx/internal/domain/models"
)

type LocationRepository interface {
	Create(ctx context.Context, location *models.Location) error
	GetByUUID(ctx context.Context, uuid uuid.UUID) (*models.Location, error)
	GetByName(ctx context.Context, name string) (*models.Location, error)
	Update(ctx context.Context, location *models.Location) error
	Delete(ctx context.Context, uuid uuid.UUID) error
}
