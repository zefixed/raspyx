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
	"strings"
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

func (uc *ScheduleUseCase) scheduleDTOToScheduleModel(ctx context.Context, scheduleDTO *dto.ScheduleRequest) (*models.Schedule, error) {
	const op = "usecase.schedule.scheduleDTOToScheduleModel"

	schedule := &models.Schedule{}

	// Adding groupUUID to model
	groupUUID, err := uc.repoGroup.GetByNumber(ctx, scheduleDTO.Group)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	schedule.GroupUUID = groupUUID.UUID

	// Adding subjectUUID to model
	subjectUUID, err := uuid.Parse(scheduleDTO.SubjectUUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, ErrInvalidUUID)
	}
	_, err = uc.repoSubject.GetByUUID(ctx, subjectUUID)
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
	if scheduleDTO.Weekday < 1 || scheduleDTO.Weekday > 6 {
		return nil, fmt.Errorf("%s: %w", op, errors.New("invalid weekday"))
	}
	schedule.Weekday = scheduleDTO.Weekday

	// Adding link to model
	schedule.Link = scheduleDTO.Link

	return schedule, nil
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

func (uc *ScheduleUseCase) Create(ctx context.Context, scheduleDTO *dto.ScheduleRequest) (*dto.CreateScheduleResponse, error) {
	const op = "usecase.schedule.Create"

	// DTO to model
	schedule, err := uc.scheduleDTOToScheduleModel(ctx, scheduleDTO)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Generating new uuid
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, errors.New("internal error"))
	}
	schedule.UUID = newUUID

	// Adding schedule to db
	err = uc.repo.Create(ctx, schedule)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Adding teachers to created schedule
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

		err = uc.repoTToS.Create(ctx, &models.TeachersToSchedule{
			TeacherUUID: teacher.UUID, ScheduleUUID: schedule.UUID,
		})
		if err != nil {
			_ = uc.repo.Delete(ctx, schedule.UUID)
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	// Adding rooms to created schedule
	for _, roomNumber := range scheduleDTO.Rooms {
		room, err := uc.repoRoom.GetByNumber(ctx, roomNumber)
		if err != nil {
			_ = uc.repo.Delete(ctx, schedule.UUID)
			return nil, fmt.Errorf("%s: %w", op, err)
		}

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

func (uc *ScheduleUseCase) GetByTeacher(ctx context.Context, fn string) (*dto.Week, error) {
	const op = "usecase.schedule.GetByTeacher"

	fnArr := strings.Split(strings.TrimSpace(fn), " ")
	for i := 0; i < len(fnArr); i++ {
		if fnArr[i] == "" {
			fnArr = append(fnArr[:i], fnArr[i+1:]...)
			i--
		}
	}
	if len(fnArr) < 2 || len(fnArr) > 3 {
		return nil, fmt.Errorf("%s: %w", op, errors.New("invalid fullname"))
	}

	// Getting schedule from db with given teacher fullname
	fnArr = append(fnArr, "")
	schedules, err := uc.repo.GetByTeacher(ctx, fnArr[1], fnArr[0], fnArr[2])
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return makeWeek(schedules), nil
}

func (uc *ScheduleUseCase) GetByTeacherUUID(ctx context.Context, UUID string) (*dto.Week, error) {
	const op = "usecase.schedule.GetByTeacherUUID"

	teacherUUID, err := uuid.Parse(UUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, ErrInvalidUUID)
	}

	// Getting schedule from db with given teacher uuid
	schedules, err := uc.repo.GetByTeacherUUID(ctx, teacherUUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return makeWeek(schedules), nil
}

func (uc *ScheduleUseCase) GetByGroup(ctx context.Context, groupNumber string) (*dto.Week, error) {
	const op = "usecase.schedule.GetByGroup"

	groupNumber = strings.TrimSpace(groupNumber)

	// Getting schedule from db with given group number
	schedules, err := uc.repo.GetByGroup(ctx, groupNumber)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return makeWeek(schedules), nil
}

func (uc *ScheduleUseCase) GetByGroupUUID(ctx context.Context, UUID string) (*dto.Week, error) {
	const op = "usecase.schedule.GetByGroupUUID"

	groupUUID, err := uuid.Parse(UUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, ErrInvalidUUID)
	}

	// Getting schedule from db with given group uuid
	schedules, err := uc.repo.GetByGroupUUID(ctx, groupUUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return makeWeek(schedules), nil
}

func (uc *ScheduleUseCase) GetByRoom(ctx context.Context, roomNumber string) (*dto.Week, error) {
	const op = "usecase.schedule.GetByRoom"

	roomNumber = strings.TrimSpace(roomNumber)

	// Getting schedule from db with given room number
	schedules, err := uc.repo.GetByRoom(ctx, roomNumber)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return makeWeek(schedules), nil
}

func (uc *ScheduleUseCase) GetByRoomUUID(ctx context.Context, UUID string) (*dto.Week, error) {
	const op = "usecase.schedule.GetByRoomUUID"

	roomUUID, err := uuid.Parse(UUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, ErrInvalidUUID)
	}

	// Getting schedule from db with given room uuid
	schedules, err := uc.repo.GetByRoomUUID(ctx, roomUUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return makeWeek(schedules), nil
}

func (uc *ScheduleUseCase) GetBySubject(ctx context.Context, subjectName string) (*dto.Week, error) {
	const op = "usecase.schedule.GetBySubject"

	subjectName = strings.TrimSpace(subjectName)

	// Getting schedule from db with given subject name
	schedules, err := uc.repo.GetBySubject(ctx, subjectName)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return makeWeek(schedules), nil
}

func (uc *ScheduleUseCase) GetBySubjectUUID(ctx context.Context, UUID string) (*dto.Week, error) {
	const op = "usecase.schedule.GetBySubjectUUID"

	subjectUUID, err := uuid.Parse(UUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, ErrInvalidUUID)
	}

	// Getting schedule from db with given subject uuid
	schedules, err := uc.repo.GetBySubjectUUID(ctx, subjectUUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return makeWeek(schedules), nil
}

func (uc *ScheduleUseCase) GetByLocation(ctx context.Context, locationName string) (*dto.Week, error) {
	const op = "usecase.schedule.GetByLocation"

	locationName = strings.TrimSpace(locationName)

	// Getting schedule from db with given location name
	schedules, err := uc.repo.GetByLocation(ctx, locationName)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return makeWeek(schedules), nil
}

func (uc *ScheduleUseCase) GetByLocationUUID(ctx context.Context, UUID string) (*dto.Week, error) {
	const op = "usecase.schedule.GetByLocationUUID"

	locationUUID, err := uuid.Parse(UUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, ErrInvalidUUID)
	}

	// Getting schedule from db with given location uuid
	schedules, err := uc.repo.GetByLocationUUID(ctx, locationUUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return makeWeek(schedules), nil
}

func (uc *ScheduleUseCase) rollbackUpdate(
	ctx context.Context,
	schedule *models.Schedule,
	TToS []*models.TeachersToSchedule,
	RToS []*models.RoomsToSchedule,
) error {
	const op = "usecase.schedule.rollbackUpdate"

	err := uc.repo.Delete(ctx, schedule.UUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = uc.repo.Create(ctx, schedule)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	for _, t := range TToS {
		err = uc.repoTToS.Create(ctx, t)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	for _, r := range RToS {
		err = uc.repoRToS.Create(ctx, r)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	return nil
}

func (uc *ScheduleUseCase) Update(ctx context.Context, UUID string, scheduleDTO *dto.ScheduleRequest) error {
	const op = "usecase.schedule.Update"

	// Parsing given UUID
	scheduleUUID, err := uuid.Parse(UUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, ErrInvalidUUID)
	}

	// Getting old schedule
	oldSchedule, err := uc.repo.GetForUpdate(ctx, scheduleUUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	// Getting old TToS
	oldTTos, err := uc.repoTToS.GetByScheduleUUID(ctx, scheduleUUID)
	if err != nil && !strings.Contains(err.Error(), "not found") {
		return fmt.Errorf("%s: %w", op, err)
	}

	// Getting old RToS
	oldRToS, err := uc.repoRToS.GetByScheduleUUID(ctx, scheduleUUID)
	if err != nil && !strings.Contains(err.Error(), "not found") {
		return fmt.Errorf("%s: %w", op, err)
	}

	// DTO to model
	newSchedule, err := uc.scheduleDTOToScheduleModel(ctx, scheduleDTO)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	newSchedule.UUID = scheduleUUID

	// Updating schedule
	err = uc.repo.Delete(ctx, scheduleUUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = uc.repo.Create(ctx, newSchedule)
	if err != nil {
		er := uc.rollbackUpdate(ctx, oldSchedule, oldTTos, oldRToS)
		if er != nil {
			return fmt.Errorf("%s: %w, %w", op, err, er)
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	// Adding teachers to created schedule
	for _, UUID := range scheduleDTO.TeachersUUID {
		teacherUUID, err := uuid.Parse(UUID)
		if err != nil {
			er := uc.rollbackUpdate(ctx, oldSchedule, oldTTos, oldRToS)
			if er != nil {
				return fmt.Errorf("%s: %w, %w", op, err, er)
			}
			return fmt.Errorf("%s: %w", op, err)
		}

		teacher, err := uc.repoTeacher.GetByUUID(ctx, teacherUUID)
		if err != nil {
			er := uc.rollbackUpdate(ctx, oldSchedule, oldTTos, oldRToS)
			if er != nil {
				return fmt.Errorf("%s: %w, %w", op, err, er)
			}
			return fmt.Errorf("%s: %w", op, err)
		}

		err = uc.repoTToS.Create(ctx, &models.TeachersToSchedule{
			TeacherUUID: teacher.UUID, ScheduleUUID: scheduleUUID,
		})
		if err != nil {
			er := uc.rollbackUpdate(ctx, oldSchedule, oldTTos, oldRToS)
			if er != nil {
				return fmt.Errorf("%s: %w, %w", op, err, er)
			}
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	// Adding rooms to created schedule
	for _, roomNumber := range scheduleDTO.Rooms {
		room, err := uc.repoRoom.GetByNumber(ctx, roomNumber)
		if err != nil {
			er := uc.rollbackUpdate(ctx, oldSchedule, oldTTos, oldRToS)
			if er != nil {
				return fmt.Errorf("%s: %w, %w", op, err, er)
			}
			return fmt.Errorf("%s: %w", op, err)
		}

		err = uc.repoRToS.Create(ctx, &models.RoomsToSchedule{
			RoomUUID: room.UUID, ScheduleUUID: scheduleUUID,
		})
		if err != nil {
			er := uc.rollbackUpdate(ctx, oldSchedule, oldTTos, oldRToS)
			if er != nil {
				return fmt.Errorf("%s: %w, %w", op, err, er)
			}
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	return nil
}

func (uc *ScheduleUseCase) Delete(ctx context.Context, UUID string) error {
	const op = "usecase.schedule.Delete"

	scheduleUUID, err := uuid.Parse(UUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, ErrInvalidUUID)
	}

	// Deleting schedule from db with given uuid
	err = uc.repo.Delete(ctx, scheduleUUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
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
