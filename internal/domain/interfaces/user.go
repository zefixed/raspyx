package interfaces

import (
	"context"
	"github.com/google/uuid"
	"raspyx/internal/domain/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	Get(ctx context.Context) ([]*models.User, error)
	GetByUUID(ctx context.Context, uuid uuid.UUID) (*models.User, error)
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	GetByAccessLevel(ctx context.Context, accessLevel int) ([]*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, uuid uuid.UUID) error
}
