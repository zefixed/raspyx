package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"raspyx/internal/domain/interfaces"
	"raspyx/internal/domain/models"
	"raspyx/internal/domain/services"
	"raspyx/internal/dto"
	"raspyx/internal/repository"
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
	cache        interfaces.Cache
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
	cache interfaces.Cache,
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
		cache:        cache,
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

	// Adding flag IsSession to model
	schedule.IsSession = scheduleDTO.IsSession

	return schedule, nil
}

func makeWeek(schedules []*models.ScheduleData) *dto.Week {
	week := &dto.Week{}

	(*week)["monday"] = &dto.Day{}
	(*week)["tuesday"] = &dto.Day{}
	(*week)["wednesday"] = &dto.Day{}
	(*week)["thursday"] = &dto.Day{}
	(*week)["friday"] = &dto.Day{}
	(*week)["saturday"] = &dto.Day{}

	for _, schedule := range schedules {
		pair := dto.Pair{
			Subject:   schedule.Subject,
			Teachers:  schedule.Teachers,
			StartDate: schedule.StartDate.Format(time.DateOnly),
			EndDate:   schedule.EndDate.Format(time.DateOnly),
			Rooms:     schedule.Rooms,
			Location:  schedule.Location,
			Type:      schedule.Type,
			Link:      schedule.Link,
		}

		wd := numToDay(schedule.Weekday)

		if (*week)[wd] == nil {
			(*week)[wd] = &dto.Day{}
		}

		switch schedule.StartTime.Format("15:04") {
		case "09:00":
			(*week)[wd].First = append((*week)[wd].First, pair)
		case "10:40":
			(*week)[wd].Second = append((*week)[wd].Second, pair)
		case "12:20":
			(*week)[wd].Third = append((*week)[wd].Third, pair)
		case "14:30":
			(*week)[wd].Fourth = append((*week)[wd].Fourth, pair)
		case "16:10":
			(*week)[wd].Fifth = append((*week)[wd].Fifth, pair)
		case "17:50":
			(*week)[wd].Sixth = append((*week)[wd].Sixth, pair)
		case "19:30":
			(*week)[wd].Seventh = append((*week)[wd].Seventh, pair)
		}
	}

	return week
}

func makeSessionWeek(schedules []*models.ScheduleData) *dto.Week {
	week := &dto.Week{}

	for _, schedule := range schedules {
		pair := dto.Pair{
			Subject:   schedule.Subject,
			Teachers:  schedule.Teachers,
			StartDate: schedule.StartDate.Format(time.DateOnly),
			EndDate:   schedule.EndDate.Format(time.DateOnly),
			Rooms:     schedule.Rooms,
			Location:  schedule.Location,
			Type:      schedule.Type,
			Link:      schedule.Link,
		}

		sd := schedule.StartDate.Format(time.DateOnly)
		if (*week)[sd] == nil {
			(*week)[sd] = &dto.Day{}
		}

		switch schedule.StartTime.Format("15:04") {
		case "09:00":
			(*week)[sd].First = append((*week)[sd].First, pair)
		case "10:40":
			(*week)[sd].Second = append((*week)[sd].Second, pair)
		case "12:20":
			(*week)[sd].Third = append((*week)[sd].Third, pair)
		case "14:30":
			(*week)[sd].Fourth = append((*week)[sd].Fourth, pair)
		case "16:10":
			(*week)[sd].Fifth = append((*week)[sd].Fifth, pair)
		case "17:50":
			(*week)[sd].Sixth = append((*week)[sd].Sixth, pair)
		case "19:30":
			(*week)[sd].Seventh = append((*week)[sd].Seventh, pair)
		}
	}

	return week
}

func numToDay(num int) string {
	return map[int]string{
		1: "monday",
		2: "tuesday",
		3: "wednesday",
		4: "thursday",
		5: "friday",
		6: "saturday",
	}[num]
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
		return nil, fmt.Errorf("%s: %w", op, ErrGeneratingUUID)
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

	// Parsing schedule uuid
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

func (uc *ScheduleUseCase) GetByTeacher(ctx context.Context, fn string, isSession bool) (*dto.Week, error) {
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
	schedules, err := uc.repo.GetByTeacher(ctx, fnArr[1], fnArr[0], fnArr[2], isSession)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return makeWeek(schedules), nil
}

func (uc *ScheduleUseCase) GetByTeacherUUID(ctx context.Context, UUID string, isSession bool) (*dto.Week, error) {
	const op = "usecase.schedule.GetByTeacherUUID"

	// Parsing teacher uuid
	teacherUUID, err := uuid.Parse(UUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, ErrInvalidUUID)
	}

	// Getting schedule from db with given teacher uuid
	schedules, err := uc.repo.GetByTeacherUUID(ctx, teacherUUID, isSession)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return makeWeek(schedules), nil
}

func (uc *ScheduleUseCase) GetByGroup(ctx context.Context, groupNumber string, isSession bool) (*dto.Week, error) {
	const op = "usecase.schedule.GetByGroup"
	cacheKey := "schedule:" + groupNumber
	if isSession {
		cacheKey += ":1"
	} else {
		cacheKey += ":0"
	}

	groupNumber = strings.TrimSpace(groupNumber)

	if cached, err := uc.cache.Get(ctx, cacheKey); err == nil {
		var week dto.Week
		_ = json.Unmarshal([]byte(cached), &week)
		return &week, nil
	}

	// Getting schedule from db with given group number
	schedules, err := uc.repo.GetByGroup(ctx, groupNumber, isSession)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var week *dto.Week
	if !isSession {
		week = makeWeek(schedules)
	} else {
		week = makeSessionWeek(schedules)
	}

	data, _ := json.Marshal(week)
	_ = uc.cache.Set(ctx, cacheKey, string(data), 10*time.Second)

	return week, nil
}

func (uc *ScheduleUseCase) GetByGroupUUID(ctx context.Context, UUID string, isSession bool) (*dto.Week, error) {
	const op = "usecase.schedule.GetByGroupUUID"

	// Parsing group uuid
	groupUUID, err := uuid.Parse(UUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, ErrInvalidUUID)
	}

	// Getting schedule from db with given group uuid
	schedules, err := uc.repo.GetByGroupUUID(ctx, groupUUID, isSession)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return makeWeek(schedules), nil
}

func (uc *ScheduleUseCase) GetByRoom(ctx context.Context, roomNumber string, isSession bool) (*dto.Week, error) {
	const op = "usecase.schedule.GetByRoom"

	roomNumber = strings.TrimSpace(roomNumber)

	// Getting schedule from db with given room number
	schedules, err := uc.repo.GetByRoom(ctx, roomNumber, isSession)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return makeWeek(schedules), nil
}

func (uc *ScheduleUseCase) GetByRoomUUID(ctx context.Context, UUID string, isSession bool) (*dto.Week, error) {
	const op = "usecase.schedule.GetByRoomUUID"

	// Parsing room uuid
	roomUUID, err := uuid.Parse(UUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, ErrInvalidUUID)
	}

	// Getting schedule from db with given room uuid
	schedules, err := uc.repo.GetByRoomUUID(ctx, roomUUID, isSession)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return makeWeek(schedules), nil
}

func (uc *ScheduleUseCase) GetBySubject(ctx context.Context, subjectName string, isSession bool) (*dto.Week, error) {
	const op = "usecase.schedule.GetBySubject"

	subjectName = strings.TrimSpace(subjectName)

	// Getting schedule from db with given subject name
	schedules, err := uc.repo.GetBySubject(ctx, subjectName, isSession)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return makeWeek(schedules), nil
}

func (uc *ScheduleUseCase) GetBySubjectUUID(ctx context.Context, UUID string, isSession bool) (*dto.Week, error) {
	const op = "usecase.schedule.GetBySubjectUUID"

	// Parsing subject uuid
	subjectUUID, err := uuid.Parse(UUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, ErrInvalidUUID)
	}

	// Getting schedule from db with given subject uuid
	schedules, err := uc.repo.GetBySubjectUUID(ctx, subjectUUID, isSession)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return makeWeek(schedules), nil
}

func (uc *ScheduleUseCase) GetByLocation(ctx context.Context, locationName string, isSession bool) (*dto.Week, error) {
	const op = "usecase.schedule.GetByLocation"

	locationName = strings.TrimSpace(locationName)

	// Getting schedule from db with given location name
	schedules, err := uc.repo.GetByLocation(ctx, locationName, isSession)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return makeWeek(schedules), nil
}

func (uc *ScheduleUseCase) GetByLocationUUID(ctx context.Context, UUID string, isSession bool) (*dto.Week, error) {
	const op = "usecase.schedule.GetByLocationUUID"

	// Parsing location uuid
	locationUUID, err := uuid.Parse(UUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, ErrInvalidUUID)
	}

	// Getting schedule from db with given location uuid
	schedules, err := uc.repo.GetByLocationUUID(ctx, locationUUID, isSession)
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

	// Parsing schedule UUID
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

	// Parsing schedule uuid
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

func (uc *ScheduleUseCase) DeletePairsByGroupWeekdayTime(ctx context.Context, data *dto.DeletePBGWTRequest, isSession bool) error {
	const op = "usecase.schedule.DeletePairsByGroupWeekdayTime"

	// Getting group from db by given group number
	group, err := uc.repoGroup.GetByNumber(ctx, data.Group)
	if err != nil {
		if strings.Contains(err.Error(), repository.ErrNotFound.Error()) {
			return fmt.Errorf("%s: group %v", op, repository.ErrNotFound)
		}
		return fmt.Errorf("%s: %v", op, err)
	}

	// Validation weekday
	if data.Weekday < 1 || data.Weekday > 6 {
		return fmt.Errorf("%s: invalid weekday", op)
	}

	// Parsing pair start time
	var startTime time.Time
	switch data.PairNum {
	case 1:
		startTime, _ = time.Parse(time.TimeOnly, "09:00:00")
	case 2:
		startTime, _ = time.Parse(time.TimeOnly, "10:40:00")
	case 3:
		startTime, _ = time.Parse(time.TimeOnly, "12:20:00")
	case 4:
		startTime, _ = time.Parse(time.TimeOnly, "14:30:00")
	case 5:
		startTime, _ = time.Parse(time.TimeOnly, "16:10:00")
	case 6:
		startTime, _ = time.Parse(time.TimeOnly, "17:50:00")
	case 7:
		startTime, _ = time.Parse(time.TimeOnly, "19:30:00")
	default:
		return fmt.Errorf("%s: invalid pair num", op)
	}

	// Deleting schedule from db with given data
	err = uc.repo.DeletePairsByGroupWeekdayTime(ctx, group.UUID, data.Weekday, startTime, data.StartDate, isSession)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
