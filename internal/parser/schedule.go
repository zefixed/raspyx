package parser

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"io"
	"log/slog"
	"net/http"
	"raspyx/config"
	"raspyx/internal/domain/interfaces"
	"raspyx/internal/domain/services"
	"raspyx/internal/dto"
	"raspyx/internal/repository"
	"raspyx/internal/repository/postgres"
	myredis "raspyx/internal/repository/redis"
	"raspyx/internal/usecase"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type ScheduleParser struct {
	client       *http.Client
	conn         *pgxpool.Pool
	log          *slog.Logger
	cfg          config.Parser
	added        *added
	groupRepo    *postgres.GroupRepository
	groupSVC     *services.GroupService
	sbjRepo      *postgres.SubjectRepository
	sbjSVC       *services.SubjectService
	teacherRepo  *postgres.TeacherRepository
	teacherSVC   *services.TeacherService
	roomRepo     *postgres.RoomRepository
	roomSVC      *services.RoomService
	locationRepo *postgres.LocationRepository
	locationSVC  *services.LocationService
	typeRepo     *postgres.SubjectTypeRepository
	typeSVC      *services.SubjectTypeService
	scheduleRepo *postgres.ScheduleRepository
	scheduleSVC  *services.ScheduleService
	repoTToS     interfaces.TeachersToScheduleRepository
	repoRToS     interfaces.RoomsToScheduleRepository
	cache        interfaces.Cache
}

type lesson struct {
	Sbj        string `json:"sbj"`
	Teacher    string `json:"teacher"`
	Dts        string `json:"dts"`
	Df         string `json:"df"`
	Dt         string `json:"dt"`
	Auditories []struct {
		Title string `json:"title"`
		Color string `json:"color"`
	} `json:"auditories"`
	ShortRooms []string `json:"shortRooms"`
	Location   string   `json:"location"`
	Type       string   `json:"type"`
	Week       string   `json:"week"`
	Align      string   `json:"align"`
	ELink      any      `json:"e_link"`
}

type response struct {
	Status string                         `json:"status"`
	Grid   map[string]map[string][]lesson `json:"grid"`
}

type added struct {
	groups    int
	subjects  int
	teachers  int
	rooms     int
	locations int
	types     int
	schedule  int
}

func NewScheduleParser(timeout time.Duration, conn *pgxpool.Pool, redisClient *redis.Client, log *slog.Logger, cfg config.Parser) *ScheduleParser {
	return &ScheduleParser{
		client: &http.Client{Timeout: timeout},
		conn:   conn,
		log:    log,
		cache:  myredis.NewRedisCache(redisClient),
		cfg:    cfg,
	}
}

func (p *ScheduleParser) New(ctx context.Context) {
	p.log = p.log.With(slog.String("module", "ScheduleParser"))

	p.groupRepo = postgres.NewGroupRepository(p.conn)
	p.groupSVC = services.NewGroupService()

	p.sbjRepo = postgres.NewSubjectRepository(p.conn)
	p.sbjSVC = services.NewSubjectService()

	p.teacherRepo = postgres.NewTeacherRepository(p.conn)
	p.teacherSVC = services.NewTeacherService()

	p.roomRepo = postgres.NewRoomRepository(p.conn)
	p.roomSVC = services.NewRoomService()

	p.locationRepo = postgres.NewLocationRepository(p.conn)
	p.locationSVC = services.NewLocationService()

	p.typeRepo = postgres.NewSubjectTypeRepository(p.conn)
	p.typeSVC = services.NewSubjectTypeService()

	p.scheduleRepo = postgres.NewScheduleRepository(p.conn)
	p.scheduleSVC = services.NewScheduleService()
	p.repoTToS = postgres.NewTeachersToScheduleRepository(p.conn)
	p.repoRToS = postgres.NewRoomsToScheduleRepository(p.conn)

	// Init parsing schedule
	err := p.parse(ctx)
	if err != nil {
		p.log.Error(fmt.Sprintf("error parsing schedule: %v", err))
	}

	// Set timeout to 10 minute if it too small
	if p.cfg.Timeout < 1 {
		p.cfg.Timeout = 10
	}

	// Ticker for parsing schedule
	ticker := time.NewTicker(time.Duration(p.cfg.Timeout) * time.Minute)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ctx.Done():
				p.log.Error(fmt.Sprintf("cancel schedule parser"))
				return
			case <-ticker.C:
				err := p.parse(ctx)
				if err != nil {
					p.log.Error(fmt.Sprintf("error parsing schedule: %v", err))
				}
			}
		}
	}()

	<-ctx.Done()
}

func (p *ScheduleParser) parse(ctx context.Context) error {
	t := time.Now()

	p.added = &added{}

	// Parsing groups
	groups, err := p.parseGroups(ctx)
	if err != nil {
		return err
	}

	// Adding groups to db
	p.addGroupsToDB(ctx, groups)

	// Parsing schedule
	// Not working with goroutines, rasp.dmami.ru not handling 500rps
	// Parsing schedule time ~3m, first(init) ~9m
	for _, group := range groups {
		err = p.parseGroupSchedule(ctx, group)
		if err != nil {
			p.log.Error(fmt.Sprintf("error parsing schedule for %v: %v", group, err))
		}
	}

	//err = p.parseGroupSchedule(ctx, "221-376")
	//err = p.parseGroupSchedule(ctx, "221-352")
	//if err != nil {
	//	p.log.Error(fmt.Sprintf("error parsing schedule for %v: %v", "221-376", err))
	//}

	p.log.Info(
		"schedule parsed",
		slog.String("time_taken", time.Since(t).String()),
		slog.Any("added", map[string]int{
			"schedules": p.added.schedule,
			"groups":    p.added.groups,
			"subjects":  p.added.subjects,
			"teachers":  p.added.teachers,
			"rooms":     p.added.rooms,
			"locations": p.added.locations,
			"types":     p.added.types,
		}),
	)

	return nil
}

func (p *ScheduleParser) parseGroups(ctx context.Context) ([]string, error) {
	// New request to rasp.dmami.ru
	req, err := http.NewRequestWithContext(ctx, "GET", "https://rasp.dmami.ru/", nil)
	if err != nil {
		return nil, err
	}

	// Set referer
	req.Header.Set("Referer", "https://rasp.dmami.ru/")

	// Sending request
	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Reading response
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Collect groups from response
	re := regexp.MustCompile(`\d{2}[0-9a-zA-Zа-яА-Я]-\d{3}(\s{1}[a-zA-Zа-яА-я]{3})?`)
	matches := re.FindAll(raw, -1)

	// Deleting repeats from groups
	gm := make(map[string]int)
	for _, m := range matches {
		if _, ok := gm[string(m)]; !ok {
			gm[string(m)] = 1
		}
	}

	// From map of groups to []string
	groups := make([]string, 0, len(gm))
	for group, _ := range gm {
		groups = append(groups, group)
	}

	return groups, nil
}

func (p *ScheduleParser) addGroupsToDB(ctx context.Context, groups []string) {
	groupUC := usecase.NewGroupUseCase(p.groupRepo, *p.groupSVC)
	for _, group := range groups {
		// Adding group to db
		_, err := groupUC.Create(ctx, &dto.CreateGroupRequest{Group: strings.TrimSpace(group)})

		// If error != group exist
		if err != nil {
			if !strings.Contains(err.Error(), repository.ErrExist.Error()) {
				p.log.Error(fmt.Sprintf("error adding group %v to db: %v", group, err))
			}
		} else {
			p.added.groups++
		}
	}
}

func (p *ScheduleParser) parseGroupSchedule(ctx context.Context, group string) error {

	// New request to rasp.dmami.ru
	req, err := http.NewRequestWithContext(
		ctx,
		"GET", fmt.Sprintf("https://rasp.dmami.ru/site/group?group=%v", group),
		nil,
	)
	if err != nil {
		return err
	}

	// Set referer
	req.Header.Set("Referer", "https://rasp.dmami.ru/")

	// Sending request
	resp, err := p.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Reading response
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Unmarshall response
	var r response
	if err := json.Unmarshal(raw, &r); err != nil {
		p.log.Error(fmt.Sprintf("error unmarshling response %v: %v", r, err))
	}

	// Adding subjects to db
	p.addSubjectsToDB(ctx, &r)

	// Adding teachers to db
	p.addTeachersToDB(ctx, &r)

	// Adding rooms to db
	p.addRoomsToDB(ctx, &r)

	// Adding locations to db
	p.addLocationsToDB(ctx, &r)

	// Adding types to db
	p.addTypesToDB(ctx, &r)

	// Adding types to db
	p.addSchedulesToDB(ctx, group, &r)

	return nil
}

func (p *ScheduleParser) addSubjectsToDB(ctx context.Context, r *response) {
	sbjUC := usecase.NewSubjectUseCase(p.sbjRepo, *p.sbjSVC)

	for day := 1; day < 7; day++ {
		for pair := 1; pair < 8; pair++ {
			for _, pairData := range r.Grid[strconv.Itoa(day)][strconv.Itoa(pair)] {
				_, err := sbjUC.GetByName(ctx, strings.TrimSpace(pairData.Sbj))
				if err != nil && strings.Contains(err.Error(), repository.ErrNotFound.Error()) {
					_, err = sbjUC.Create(ctx, &dto.CreateSubjectRequest{Name: strings.TrimSpace(pairData.Sbj)})
					if err != nil {
						p.log.Error(fmt.Sprintf("error adding subject %v to db: %v", pairData.Sbj, err))
					} else {
						p.added.subjects++
					}
				}
			}
		}
	}
}

func (p *ScheduleParser) addTeachersToDB(ctx context.Context, r *response) {
	teacherUC := usecase.NewTeacherUseCase(p.teacherRepo, *p.teacherSVC)

	for day := 1; day < 7; day++ {
		for pair := 1; pair < 8; pair++ {
			for _, pairData := range r.Grid[strconv.Itoa(day)][strconv.Itoa(pair)] {
				teachers := teachersFromString(pairData.Teacher)
				for _, fullname := range teachers {
					// first, last, middle
					flm := strings.Split(fullname, " ")

					flm = append(flm, []string{"", ""}...)

					if len(flm) > 2 {
						_, err := teacherUC.GetByFullName(ctx, strings.TrimSpace(strings.Join(flm, " ")))
						if err != nil && strings.Contains(err.Error(), repository.ErrNotFound.Error()) {
							_, err = teacherUC.Create(ctx, &dto.CreateTeacherRequest{
								FirstName:  flm[1],
								SecondName: flm[0],
								MiddleName: strings.TrimSpace(strings.Join(flm[2:], " ")),
							})
							if err != nil {
								p.log.Error(fmt.Sprintf("error adding teacher %v to db: %v", strings.TrimSpace(strings.Join(flm, " ")), err))
							} else {
								p.added.teachers++
							}
						}
					}
				}

			}
		}
	}
}

func (p *ScheduleParser) addRoomsToDB(ctx context.Context, r *response) {
	roomUC := usecase.NewRoomUseCase(p.roomRepo, *p.roomSVC)

	for day := 1; day < 7; day++ {
		for pair := 1; pair < 8; pair++ {
			for _, pairData := range r.Grid[strconv.Itoa(day)][strconv.Itoa(pair)] {
				for _, room := range pairData.Auditories {
					roomNum := removeHTML(removeEmojis(room.Title))
					_, err := roomUC.GetByNumber(ctx, roomNum)
					if err != nil && strings.Contains(err.Error(), repository.ErrNotFound.Error()) {
						_, err = roomUC.Create(ctx, &dto.CreateRoomRequest{Number: strings.TrimSpace(roomNum)})
						if err != nil {
							p.log.Error(fmt.Sprintf("error adding room %v to db: %v", roomNum, err))
						} else {
							p.added.rooms++
						}
					}
				}

			}
		}
	}
}

func removeEmojis(text string) string {
	emojiRegex := regexp.MustCompile(`[\x{1F600}-\x{1F64F}]|[\x{1F300}-\x{1F5FF}]|[\x{1F680}-\x{1F6FF}]|[\x{2600}-\x{26FF}]|[\x{2700}-\x{27BF}]`)
	return strings.TrimSpace(emojiRegex.ReplaceAllString(text, ""))
}

func removeHTML(text string) string {
	htmlRegex := regexp.MustCompile(`>.*<`)
	newText := htmlRegex.FindString(text)
	if newText == "" {
		return text
	}
	return strings.TrimSpace(newText[1 : len(newText)-1])
}

func (p *ScheduleParser) addLocationsToDB(ctx context.Context, r *response) {
	locationUC := usecase.NewLocationUseCase(p.locationRepo, *p.locationSVC)

	for day := 1; day < 7; day++ {
		for pair := 1; pair < 8; pair++ {
			for _, pairData := range r.Grid[strconv.Itoa(day)][strconv.Itoa(pair)] {
				_, err := locationUC.GetByName(ctx, strings.TrimSpace(pairData.Location))
				if err != nil && strings.Contains(err.Error(), repository.ErrNotFound.Error()) {
					_, err = locationUC.Create(ctx, &dto.CreateLocationRequest{Name: strings.TrimSpace(pairData.Location)})
					if err != nil {
						p.log.Error(fmt.Sprintf("error adding location %v to db: %v", pairData.Location, err))
					} else {
						p.added.locations++
					}
				}

			}
		}
	}
}

func (p *ScheduleParser) addTypesToDB(ctx context.Context, r *response) {
	typeUC := usecase.NewSubjectTypeUseCase(p.typeRepo, *p.typeSVC)

	for day := 1; day < 7; day++ {
		for pair := 1; pair < 8; pair++ {
			for _, pairData := range r.Grid[strconv.Itoa(day)][strconv.Itoa(pair)] {
				_, err := typeUC.GetByType(ctx, strings.TrimSpace(pairData.Type))
				if err != nil && strings.Contains(err.Error(), repository.ErrNotFound.Error()) {
					_, err = typeUC.Create(ctx, &dto.CreateSubjectTypeRequest{Type: strings.TrimSpace(pairData.Type)})
					if err != nil {
						p.log.Error(fmt.Sprintf("error adding type %v to db: %v", pairData.Type, err))
					} else {
						p.added.types++
					}
				}
			}
		}
	}
}

func (p *ScheduleParser) addSchedulesToDB(ctx context.Context, group string, r *response) {
	scheduleUC := usecase.NewScheduleUseCase(
		p.scheduleRepo, p.groupRepo, p.sbjRepo, p.typeRepo,
		p.locationRepo, p.teacherRepo, p.roomRepo, p.repoTToS,
		p.repoRToS, *p.scheduleSVC, p.cache)

	week, err := scheduleUC.GetByGroup(ctx, group)

	if err != nil && strings.Contains(err.Error(), repository.ErrNotFound.Error()) {
		week = &dto.Week{}
	}

	for day := 1; day < 7; day++ {
		for pair := 1; pair < 8; pair++ {
			parsedPairs := r.Grid[strconv.Itoa(day)][strconv.Itoa(pair)]
			dbPairs := getPairs(getDay(week, numToDay(day)), numToPair(pair))
			if len(parsedPairs) == 0 && len(dbPairs) == 0 {
				continue
			} else if len(parsedPairs) == 0 && len(dbPairs) > 0 {
				err = scheduleUC.DeletePairsByGroupWeekdayTime(ctx, &dto.DeletePBGWTRequest{Group: group, Weekday: day, PairNum: pair})
				if err != nil && !strings.Contains(err.Error(), repository.ErrNotFound.Error()) {
					p.log.Error(fmt.Sprintf(
						"error deleting schedule for the group %v on %v at %v: %v",
						group, numToDay(day), numToPair(pair), err,
					))
				}
			} else {
				var parsedPairsDTO []dto.Pair
				for _, pairData := range parsedPairs {
					// Removing trash from rooms
					var rooms []string
					for _, room := range pairData.Auditories {
						rooms = append(rooms, removeHTML(removeEmojis(room.Title)))
					}

					// Sorting teachers and rooms in parsed pair
					teachers := teachersFromString(pairData.Teacher)
					sort.Slice(teachers, func(i, j int) bool { return strings.ToLower(teachers[i]) < strings.ToLower(teachers[j]) })
					sort.Slice(rooms, func(i, j int) bool { return strings.ToLower(rooms[i]) < strings.ToLower(rooms[j]) })

					// Mapping parsed pair to dto
					pairDataDTO := &dto.Pair{
						Subject:   strings.TrimSpace(pairData.Sbj),
						Teachers:  teachers,
						StartDate: strings.TrimSpace(pairData.Df),
						EndDate:   strings.TrimSpace(pairData.Dt),
						Rooms:     rooms,
						Location:  strings.TrimSpace(pairData.Location),
						Type:      strings.TrimSpace(pairData.Type),
						Link:      getLinkFromHTML(pairData.Auditories[0].Title),
					}
					parsedPairsDTO = append(parsedPairsDTO, *pairDataDTO)
				}

				// Sorting teachers and rooms in pair from db
				for _, pairData := range dbPairs {
					sort.Slice(pairData.Teachers, func(i, j int) bool {
						return strings.ToLower(pairData.Teachers[i]) < strings.ToLower(pairData.Teachers[j])
					})
					sort.Slice(pairData.Rooms, func(i, j int) bool { return strings.ToLower(pairData.Rooms[i]) < strings.ToLower(pairData.Rooms[j]) })
				}

				if !cmp.Equal(parsedPairsDTO, dbPairs) {
					// Deleting pair from db
					err = scheduleUC.DeletePairsByGroupWeekdayTime(ctx, &dto.DeletePBGWTRequest{Group: group, Weekday: day, PairNum: pair})
					if err != nil && !strings.Contains(err.Error(), repository.ErrNotFound.Error()) {
						p.log.Error(fmt.Sprintf(
							"error deleting schedule for the group %v on %v at %v: %v",
							group, numToDay(day), numToPair(pair), err,
						))
					}

					// Adding new pair to db
					for _, pairData := range parsedPairs {
						// Removing trash from rooms
						var rooms []string
						for _, room := range pairData.Auditories {
							rooms = append(rooms, removeHTML(removeEmojis(room.Title)))
						}

						// Getting start and end times from pair num
						st, et := pairNumToSTET(pair)

						// Getting teachers uuid
						teacherUC := usecase.NewTeacherUseCase(p.teacherRepo, *p.teacherSVC)
						var teachersUUID []string
						if pairData.Teacher != "" {
							teachersUUID, err = teachersToUUID(ctx, strings.Split(pairData.Teacher, ", "), teacherUC)
							if err != nil {
								p.log.Error(fmt.Sprintf("error getting teachers uuid %v: %v", pairData.Teacher, err))
							}
						}

						// Getting subject uuid
						subjUC := usecase.NewSubjectUseCase(p.sbjRepo, *p.sbjSVC)
						subjUUID, err := subjectToUUID(ctx, pairData.Sbj, subjUC)
						if err != nil {
							p.log.Error(fmt.Sprintf("error getting subject %v uuid: %v", pairData.Sbj, err))
						}

						// Mapping parsed pair to dto
						pairDataDTO := &dto.ScheduleRequest{
							Group:        group,
							TeachersUUID: teachersUUID,
							Rooms:        rooms,
							SubjectUUID:  subjUUID,
							Type:         strings.TrimSpace(pairData.Type),
							Location:     strings.TrimSpace(pairData.Location),
							StartTime:    st,
							EndTime:      et,
							StartDate:    strings.TrimSpace(pairData.Df),
							EndDate:      strings.TrimSpace(pairData.Dt),
							Weekday:      day,
							Link:         getLinkFromHTML(pairData.Auditories[0].Title),
						}

						_, err = scheduleUC.Create(ctx, pairDataDTO)
						if err != nil {
							p.log.Error(
								fmt.Sprintf("error adding pair to db: %v", err),
								slog.Any("pairDataDTO", pairDataDTO),
							)
						} else {
							p.added.schedule++
						}
					}
				}
			}
		}
	}
}

func numToDay(num int) string {
	return map[int]string{
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
	}[num]
}

func numToPair(num int) string {
	return map[int]string{
		1: "First",
		2: "Second",
		3: "Third",
		4: "Fourth",
		5: "Fifth",
		6: "Sixth",
		7: "Seventh",
	}[num]
}

func getPairs(day *dto.Day, fieldName string) []dto.Pair {
	v := reflect.ValueOf(*day)
	field := v.FieldByName(fieldName)
	pairs, _ := field.Interface().([]dto.Pair)
	return pairs
}

func getDay(week *dto.Week, dayName string) *dto.Day {
	v := reflect.ValueOf(*week)
	field := v.FieldByName(dayName)
	day, _ := field.Interface().(dto.Day)
	return &day
}

func pairNumToSTET(pair int) (string, string) {
	p := map[int]string{
		1: "09:00:00 10:30:00",
		2: "10:40:00 12:10:00",
		3: "12:20:00 13:50:00",
		4: "14:30:00 16:00:00",
		5: "16:10:00 17:40:00",
		6: "17:50:00 19:20:00",
		7: "19:30:00 21:00:00",
	}
	times := strings.Split(p[pair], " ")
	return times[0], times[1]
}

func getLinkFromHTML(html string) string {
	htmlRegex := regexp.MustCompile(`['"].*?['"]`)
	links := htmlRegex.FindAllString(html, -1)

	for _, link := range links {
		if strings.Contains(link, "http") {
			return link[1 : len(link)-1]
		}
	}

	return ""
}

func teachersToUUID(ctx context.Context, teachers []string, teacherUC *usecase.TeacherUseCase) ([]string, error) {
	var teachersUUID []string
	for _, fn := range teachers {
		// Deleting extra spaces from names
		var flm []string
		for _, str := range strings.Split(fn, " ") {
			if str != "" && str != " " {
				flm = append(flm, str)
			}
		}

		teacher, err := teacherUC.GetByFullName(ctx, strings.TrimSpace(strings.Join(flm, " ")))
		if err != nil {
			return nil, err
		}
		teachersUUID = append(teachersUUID, teacher[0].UUID.String())
	}

	return teachersUUID, nil
}

func subjectToUUID(ctx context.Context, subject string, subjUC *usecase.SubjectUseCase) (string, error) {
	subjectUUID, err := subjUC.GetByName(ctx, strings.TrimSpace(subject))
	if err != nil {
		return "", err
	}

	return subjectUUID[0].UUID.String(), nil
}

func teachersFromString(str string) []string {
	var res []string
	teachers := strings.Split(str, ", ")
	for _, fullname := range teachers {
		// first, last, middle
		var flm []string

		// Deleting extra spaces from names
		for _, str := range strings.Split(fullname, " ") {
			if str != "" && str != " " {
				flm = append(flm, str)
			}
		}
		res = append(res, strings.Join(flm, " "))
	}

	return res
}
