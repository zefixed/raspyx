package interfaces

import (
	"context"
	"github.com/google/uuid"
	"raspyx/internal/domain/models"
)

type ScheduleRepository interface {
	Create(ctx context.Context, schedule *models.Schedule) error
	GetByUUID(ctx context.Context, uuid uuid.UUID) (*models.Schedule, error)
	GetByTeacher(ctx context.Context, firstName, secondName, middleName string) ([]*models.Schedule, error)
	GetByTeacherUUID(ctx context.Context, teacherUUID uuid.UUID) ([]*models.Schedule, error)
	GetByGroup(ctx context.Context, groupNumber string) ([]*models.Schedule, error)
	GetByGroupUUID(ctx context.Context, groupUUID uuid.UUID) ([]*models.Schedule, error)
	GetByRoom(ctx context.Context, roomNumber string) ([]*models.Schedule, error)
	GetByRoomUUID(ctx context.Context, roomUUID uuid.UUID) ([]*models.Schedule, error)
	GetBySubject(ctx context.Context, subjectName string) ([]*models.Schedule, error)
	GetBySubjectUUID(ctx context.Context, subjectUUID uuid.UUID) ([]*models.Schedule, error)
	GetByLocation(ctx context.Context, locationName string) ([]*models.Schedule, error)
	GetByLocationUUID(ctx context.Context, locationUUID uuid.UUID) ([]*models.Schedule, error)
	Update(ctx context.Context, schedule *models.Schedule) error
	Delete(ctx context.Context, uuid uuid.UUID) error
}
