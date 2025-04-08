package interfaces

import (
	"context"
	"github.com/google/uuid"
	"raspyx/internal/domain/models"
)

type TeachersToScheduleRepository interface {
	Create(ctx context.Context, roomsToSchedule *models.TeachersToSchedule) error
	Get(ctx context.Context) ([]*models.TeachersToSchedule, error)
	GetByTeacherUUID(ctx context.Context, teacherUUID uuid.UUID) ([]*models.TeachersToSchedule, error)
	GetByScheduleUUID(ctx context.Context, scheduleUUID uuid.UUID) ([]*models.TeachersToSchedule, error)
	Delete(ctx context.Context, roomsToSchedule *models.TeachersToSchedule) error
}
