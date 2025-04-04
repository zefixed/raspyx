package interfaces

import (
	"context"
	"github.com/google/uuid"
	"raspyx/internal/domain/models"
)

type ScheduleRepository interface {
	Create(ctx context.Context, schedule *models.Schedule) error
	Get(ctx context.Context) ([]*models.ScheduleData, error)
	GetForUpdate(ctx context.Context, uuid uuid.UUID) (*models.Schedule, error)
	GetByUUID(ctx context.Context, uuid uuid.UUID) (*models.ScheduleData, error)
	GetByTeacher(ctx context.Context, firstName, secondName, middleName string) ([]*models.ScheduleData, error)
	GetByTeacherUUID(ctx context.Context, teacherUUID uuid.UUID) ([]*models.ScheduleData, error)
	GetByGroup(ctx context.Context, groupNumber string) ([]*models.ScheduleData, error)
	GetByGroupUUID(ctx context.Context, groupUUID uuid.UUID) ([]*models.ScheduleData, error)
	GetByRoom(ctx context.Context, roomNumber string) ([]*models.ScheduleData, error)
	GetByRoomUUID(ctx context.Context, roomUUID uuid.UUID) ([]*models.ScheduleData, error)
	GetBySubject(ctx context.Context, subjectName string) ([]*models.ScheduleData, error)
	GetBySubjectUUID(ctx context.Context, subjectUUID uuid.UUID) ([]*models.ScheduleData, error)
	GetByLocation(ctx context.Context, locationName string) ([]*models.ScheduleData, error)
	GetByLocationUUID(ctx context.Context, locationUUID uuid.UUID) ([]*models.ScheduleData, error)
	Update(ctx context.Context, schedule *models.Schedule) error
	Delete(ctx context.Context, uuid uuid.UUID) error
}
