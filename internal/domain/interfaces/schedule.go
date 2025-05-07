package interfaces

import (
	"context"
	"github.com/google/uuid"
	"raspyx/internal/domain/models"
	"time"
)

type ScheduleRepository interface {
	Create(ctx context.Context, schedule *models.Schedule) error
	Get(ctx context.Context) ([]*models.ScheduleData, error)
	GetForUpdate(ctx context.Context, uuid uuid.UUID) (*models.Schedule, error)
	GetByUUID(ctx context.Context, uuid uuid.UUID) (*models.ScheduleData, error)
	GetByTeacher(ctx context.Context, firstName, secondName, middleName string, isSession bool) ([]*models.ScheduleData, error)
	GetByTeacherUUID(ctx context.Context, teacherUUID uuid.UUID, isSession bool) ([]*models.ScheduleData, error)
	GetByGroup(ctx context.Context, groupNumber string, isSession bool) ([]*models.ScheduleData, error)
	GetByGroupUUID(ctx context.Context, groupUUID uuid.UUID, isSession bool) ([]*models.ScheduleData, error)
	GetByRoom(ctx context.Context, roomNumber string, isSession bool) ([]*models.ScheduleData, error)
	GetByRoomUUID(ctx context.Context, roomUUID uuid.UUID, isSession bool) ([]*models.ScheduleData, error)
	GetBySubject(ctx context.Context, subjectName string, isSession bool) ([]*models.ScheduleData, error)
	GetBySubjectUUID(ctx context.Context, subjectUUID uuid.UUID, isSession bool) ([]*models.ScheduleData, error)
	GetByLocation(ctx context.Context, locationName string, isSession bool) ([]*models.ScheduleData, error)
	GetByLocationUUID(ctx context.Context, locationUUID uuid.UUID, isSession bool) ([]*models.ScheduleData, error)
	Update(ctx context.Context, schedule *models.Schedule) error
	Delete(ctx context.Context, uuid uuid.UUID) error
	DeletePairsByGroupWeekdayTime(ctx context.Context, group uuid.UUID, weekday int, st, sd time.Time, isSession bool) error
	DeleteByParams(ctx context.Context, params *models.ScheduleData) error
}
