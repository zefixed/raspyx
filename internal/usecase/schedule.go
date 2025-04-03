package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"raspyx/internal/domain/interfaces"
	"raspyx/internal/domain/models"
	"raspyx/internal/domain/services"
	"raspyx/internal/dto"
	"time"
)

type ScheduleUseCase struct {
	repo         interfaces.ScheduleRepository
	repoGroup    interfaces.GroupRepository
	repoSubject  interfaces.SubjectRepository
	repoType     interfaces.SubjectTypeRepository
	repoLocation interfaces.LocationRepository
	repoTeacher  interfaces.TeacherRepository
	repoRoom     interfaces.RoomRepository
	repoTToS     interfaces.TeachersToScheduleRepository
	repoRToS     interfaces.RoomsToScheduleRepository
	svc          services.ScheduleService
}

func NewScheduleUseCase(
	repo interfaces.ScheduleRepository,
	repoGroup interfaces.GroupRepository,
	repoSubject interfaces.SubjectRepository,
	repoType interfaces.SubjectTypeRepository,
	repoLocation interfaces.LocationRepository,
	repoTeacher interfaces.TeacherRepository,
	repoRoom interfaces.RoomRepository,
	repoTToS interfaces.TeachersToScheduleRepository,
	repoRToS interfaces.RoomsToScheduleRepository,
	svc services.ScheduleService,
) *ScheduleUseCase {
	return &ScheduleUseCase{
		repo:         repo,
		repoGroup:    repoGroup,
		repoSubject:  repoSubject,
		repoType:     repoType,
		repoLocation: repoLocation,
		repoTeacher:  repoTeacher,
		repoRoom:     repoRoom,
		repoTToS:     repoTToS,
		repoRToS:     repoRToS,
		svc:          svc,
	}
}

func (uc *ScheduleUseCase) Create(ctx context.Context, scheduleDTO *dto.CreateScheduleRequest) (*dto.CreateScheduleResponse, error) {
	const op = "usecase.schedule.Create"

	// Generating new uuid
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, errors.New("internal error"))
	}

	// DTO to model
	schedule := &models.Schedule{UUID: newUUID}
	// Adding groupUUID to model
	groupUUID, err := uc.repoGroup.GetByNumber(ctx, scheduleDTO.Group)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	schedule.GroupUUID = groupUUID.UUID

	// Adding subjectUUID to model
	subjectUUID, err := uuid.Parse(scheduleDTO.SubjectUUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	schedule.SubjectUUID = subjectUUID

	// Adding typeUUID to model
	typeUUID, err := uc.repoType.GetByType(ctx, scheduleDTO.Type)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	schedule.TypeUUID = typeUUID.UUID

	// Adding locationUUID to model
	locationUUID, err := uc.repoLocation.GetByName(ctx, scheduleDTO.Location)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	schedule.LocationUUID = locationUUID.UUID

	// Adding startTime to model
	startTime, err := time.Parse("15:04:05", scheduleDTO.StartTime)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, errors.New("invalid start time"))
	}
	schedule.StartTime = startTime

	// Adding endTime to model
	endTime, err := time.Parse("15:04:05", scheduleDTO.EndTime)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, errors.New("invalid end time"))
	}
	schedule.EndTime = endTime

	// Adding startDate to model
	startDate, err := time.Parse("2006-01-02", scheduleDTO.StartDate)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, errors.New("invalid start date"))
	}
	schedule.StartDate = startDate

	// Adding endDate to model
	endDate, err := time.Parse("2006-01-02", scheduleDTO.EndDate)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, errors.New("invalid end date"))
	}
	schedule.EndDate = endDate

	// Adding weekday to model
	if scheduleDTO.Weekday < 1 || scheduleDTO.Weekday > 7 {
		return nil, fmt.Errorf("%s: %w", op, errors.New("invalid weekday"))
	}
	schedule.Weekday = scheduleDTO.Weekday

	// Adding link to model
	schedule.Link = scheduleDTO.Link

	// Adding schedule to db
	err = uc.repo.Create(ctx, schedule)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var teachers []*models.Teacher
	for _, UUID := range scheduleDTO.TeachersUUID {
		teacherUUID, err := uuid.Parse(UUID)
		if err != nil {
			_ = uc.repo.Delete(ctx, schedule.UUID)
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		teacher, err := uc.repoTeacher.GetByUUID(ctx, teacherUUID)
		if err != nil {
			_ = uc.repo.Delete(ctx, schedule.UUID)
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		teachers = append(teachers, teacher)
	}

	for _, teacher := range teachers {
		err = uc.repoTToS.Create(ctx, &models.TeachersToSchedule{
			TeacherUUID: teacher.UUID, ScheduleUUID: schedule.UUID,
		})
		if err != nil {
			_ = uc.repo.Delete(ctx, schedule.UUID)
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	var rooms []*models.Room
	for _, roomNumber := range scheduleDTO.Rooms {
		roomUUID, err := uc.repoRoom.GetByNumber(ctx, roomNumber)
		if err != nil {
			_ = uc.repo.Delete(ctx, schedule.UUID)
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		room, err := uc.repoRoom.GetByUUID(ctx, roomUUID.UUID)
		if err != nil {
			_ = uc.repo.Delete(ctx, schedule.UUID)
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		rooms = append(rooms, room)
	}

	for _, room := range rooms {
		err = uc.repoRToS.Create(ctx, &models.RoomsToSchedule{
			RoomUUID: room.UUID, ScheduleUUID: schedule.UUID,
		})
		if err != nil {
			_ = uc.repo.Delete(ctx, schedule.UUID)
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	return &dto.CreateScheduleResponse{UUID: schedule.UUID}, nil
}

func makeWeek(schedules []*models.ScheduleData) *dto.Week {
	week := &dto.Week{}
	days := map[int]*dto.Day{
		1: &week.Monday, 2: &week.Tuesday, 3: &week.Wednesday,
		4: &week.Thursday, 5: &week.Friday, 6: &week.Saturday,
	}

	for _, schedule := range schedules {
		pair := dto.Pair{
			Subject:   schedule.Subject,
			Teachers:  schedule.Teachers,
			StartDate: schedule.StartDate.Format(time.DateOnly),
			EndDate:   schedule.EndDate.Format(time.DateOnly),
			Rooms:     schedule.Rooms,
			Location:  schedule.Location,
			Type:      schedule.Type,
		}

		if days[schedule.Weekday] == nil {
			days[schedule.Weekday] = &dto.Day{}
		}

		switch schedule.StartTime.Format("15:04") {
		case "09:00":
			days[schedule.Weekday].First = append(days[schedule.Weekday].First, pair)
		case "10:40":
			days[schedule.Weekday].Second = append(days[schedule.Weekday].Second, pair)
		case "12:20":
			days[schedule.Weekday].Third = append(days[schedule.Weekday].Third, pair)
		case "14:30":
			days[schedule.Weekday].Fourth = append(days[schedule.Weekday].Fourth, pair)
		case "16:10":
			days[schedule.Weekday].Fifth = append(days[schedule.Weekday].Fifth, pair)
		case "17:50":
			days[schedule.Weekday].Sixth = append(days[schedule.Weekday].Sixth, pair)
		case "19:30":
			days[schedule.Weekday].Seventh = append(days[schedule.Weekday].Seventh, pair)
		}
	}

	return week
}

func (uc *ScheduleUseCase) Get(ctx context.Context) (*dto.Week, error) {
	const op = "usecase.schedule.Get"

	// Getting all schedules from db
	schedules, err := uc.repo.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return makeWeek(schedules), nil
}

func (uc *ScheduleUseCase) GetByUUID(ctx context.Context, UUID string) (*dto.Week, error) {
	const op = "usecase.schedule.GetByUUID"

	scheduleUUID, err := uuid.Parse(UUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, ErrInvalidUUID)
	}

	// Getting schedule from db with given uuid
	schedules, err := uc.repo.GetByUUID(ctx, scheduleUUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return makeWeek([]*models.ScheduleData{schedules}), nil
}

//func (uc *ScheduleUseCase) GetByUUID(ctx context.Context, uuid uuid.UUID) (*models.ScheduleData, error) {
//	const op = "usecase.schedule.GetByUUID"
//
//	// Getting schedule from db with given uuid
//	schedule, err := uc.repo.GetByUUID(ctx, uuid)
//	if err != nil {
//		return nil, fmt.Errorf("%s: %w", op, err)
//	}
//
//	return schedule, nil
//}
//
//func (uc *ScheduleUseCase) Update(ctx context.Context, schedule *models.Schedule) error {
//	const op = "usecase.schedule.Update"
//
//	// Updating schedule in db with given schedule
//	err := uc.repo.Update(ctx, schedule)
//	if err != nil {
//		return fmt.Errorf("%s: %w", op, err)
//	}
//
//	return nil
//}
//
//func (uc *ScheduleUseCase) Delete(ctx context.Context, uuid uuid.UUID) error {
//	const op = "usecase.schedule.Delete"
//
//	// Deleting schedule from db with given uuid
//	err := uc.repo.Delete(ctx, uuid)
//	if err != nil {
//		return fmt.Errorf("%s: %w", op, err)
//	}
//	return nil
//}
