package interfaces

import (
	"context"
	"github.com/google/uuid"
	"raspyx/internal/domain/models"
)

type RoomsToScheduleRepository interface {
	Create(ctx context.Context, roomsToSchedule *models.RoomsToSchedule) error
	Get(ctx context.Context) ([]*models.RoomsToSchedule, error)
	GetByRoomUUID(ctx context.Context, roomUUID uuid.UUID) ([]*models.RoomsToSchedule, error)
	GetByScheduleUUID(ctx context.Context, scheduleUUID uuid.UUID) ([]*models.RoomsToSchedule, error)
	Delete(ctx context.Context, roomsToSchedule *models.RoomsToSchedule) error
}
