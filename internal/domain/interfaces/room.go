package interfaces

import (
	"context"
	"github.com/google/uuid"
	"raspyx/internal/domain/models"
)

type RoomRepository interface {
	Create(ctx context.Context, room *models.Room) error
	GetByUUID(ctx context.Context, uuid uuid.UUID) (*models.Room, error)
	GetByNumber(ctx context.Context, number string) (*models.Room, error)
	Update(ctx context.Context, room *models.Room) error
	Delete(ctx context.Context, uuid uuid.UUID) error
}
