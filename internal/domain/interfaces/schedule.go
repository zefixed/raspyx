package interfaces

import (
	"context"
	"github.com/google/uuid"
	"raspyx/internal/domain/models"
)

type ScheduleRepository interface {
	Create(ctx context.Context, schedule *models.Schedule) error
	GetByUUID(ctx context.Context, uuid uuid.UUID) (*models.Schedule, error)
	GetByTeacher(ctx context.Context, teacherName string) (*models.Schedule, error)
	GetByTeacherUUID(ctx context.Context, teacherUUID uuid.UUID) (*models.Schedule, error)
	GetByGroup(ctx context.Context, groupName string) (*models.Schedule, error)
	GetByGroupUUID(ctx context.Context, groupUUID uuid.UUID) (*models.Schedule, error)
	GetByRoom(ctx context.Context, room string) (*models.Schedule, error)
	GetByRoomUUID(ctx context.Context, roomUUID uuid.UUID) (*models.Schedule, error)
	GetBySubject(ctx context.Context, subject string) (*models.Schedule, error)
	GetBySubjectUUID(ctx context.Context, subjectUUID uuid.UUID) (*models.Schedule, error)
	GetByLocation(ctx context.Context, location string) (*models.Schedule, error)
	GetByLocationUUID(ctx context.Context, location uuid.UUID) (*models.Schedule, error)
	Update(ctx context.Context, schedule *models.Schedule) error
	Delete(ctx context.Context, uuid uuid.UUID) error
}
