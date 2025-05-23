package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"raspyx/internal/domain/models"
	"raspyx/internal/repository"
	"strings"
	"time"
)

type ScheduleRepository struct {
	db *pgxpool.Pool
}

func NewScheduleRepository(db *pgxpool.Pool) *ScheduleRepository {
	return &ScheduleRepository{db: db}
}

func parseSchedule(rows *pgx.Rows) ([]*models.ScheduleData, error) {
	const op = "repository.postgres.ScheduleRepository.parseSchedule"

	var schedules []*models.ScheduleData
	err := pgxscan.ScanAll(&schedules, *rows)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return schedules, nil
}

var (
	baseSelectStatement = `
		SELECT schedule.uuid AS "uuid",
			groups.number AS "group_number",
       		ARRAY_AGG(DISTINCT(TRIM(CONCAT(second_name, ' ', first_name, ' ', COALESCE(middle_name, ''))))) AS "teachers",
       		ARRAY_AGG(DISTINCT COALESCE(rooms.number, '')) AS "rooms",
			subjects.name AS "subject_name",
			subj_types.type AS "subject_type",
			locations.name AS "location",
			schedule.start_time AS "start_time",
			schedule.end_time AS "end_time",
			schedule.start_date AS "start_date",
			schedule.end_date AS "end_date",
			schedule.weekday AS "weekday",
			COALESCE(schedule.link, '') AS "link",
			schedule.is_session AS "is_session"
		FROM schedule
			LEFT JOIN groups ON schedule.group_uuid = groups.uuid
			LEFT JOIN subjects ON schedule.subject_uuid = subjects.uuid
			LEFT JOIN subj_types ON schedule.type_uuid = subj_types.uuid
			LEFT JOIN locations ON schedule.location_uuid = locations.uuid
			LEFT JOIN teachers_to_schedule ON schedule.uuid = teachers_to_schedule.schedule_uuid
			LEFT JOIN teachers ON teachers_to_schedule.teacher_uuid = teachers.uuid
			LEFT JOIN rooms_to_schedule ON schedule.uuid = rooms_to_schedule.schedule_uuid
			LEFT JOIN rooms ON rooms_to_schedule.room_uuid = rooms.uuid`
	baseGroupByStatement = `
		GROUP BY schedule.uuid, groups.number, subjects.name, subj_types.type, locations.name,
			schedule.start_time, schedule.end_time, schedule.start_date, schedule.end_date,
			schedule.weekday, schedule.link`
)

func (r *ScheduleRepository) Create(ctx context.Context, schedule *models.Schedule) error {
	const op = "repository.postgres.ScheduleRepository.Create"

	query := `INSERT INTO schedule (uuid, group_uuid, subject_uuid, type_uuid,
                      				location_uuid, start_time, end_time, start_date,
                      				end_date, weekday, link, is_session)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	_, err := r.db.Exec(
		ctx, query, schedule.UUID, schedule.GroupUUID, schedule.SubjectUUID,
		schedule.TypeUUID, schedule.LocationUUID, schedule.StartTime, schedule.EndTime,
		schedule.StartDate, schedule.EndDate, schedule.Weekday, schedule.Link, schedule.IsSession,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *ScheduleRepository) Get(ctx context.Context) ([]*models.ScheduleData, error) {
	const op = "repository.postgres.ScheduleRepository.GetByTeacherUUID"

	query := baseSelectStatement + baseGroupByStatement
	rows, err := r.db.Query(ctx, query)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	schedules, err := parseSchedule(&rows)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if len(schedules) == 0 {
		return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return schedules, nil
}

func (r *ScheduleRepository) GetForUpdate(ctx context.Context, uuid uuid.UUID) (*models.Schedule, error) {
	const op = "repository.postgres.ScheduleRepository.GetForUpdate"

	query := `SELECT uuid, group_uuid, subject_uuid, type_uuid, location_uuid, start_time,
					 end_time, start_date, end_date, weekday, link, is_session
			  FROM schedule
			  WHERE uuid = $1`
	row := r.db.QueryRow(ctx, query, uuid)

	var schedule models.Schedule
	err := row.Scan(
		&schedule.UUID, &schedule.GroupUUID, &schedule.SubjectUUID, &schedule.TypeUUID,
		&schedule.LocationUUID, &schedule.StartTime, &schedule.EndTime, &schedule.StartDate,
		&schedule.EndDate, &schedule.Weekday, &schedule.Link, &schedule.IsSession,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &schedule, nil
}

func (r *ScheduleRepository) GetByUUID(ctx context.Context, uuid uuid.UUID) (*models.ScheduleData, error) {
	const op = "repository.postgres.ScheduleRepository.GetByUUID"

	query := baseSelectStatement + ` WHERE schedule.uuid = $1 ` + baseGroupByStatement
	row := r.db.QueryRow(ctx, query, uuid)
	var schedule models.ScheduleData
	err := row.Scan(
		&schedule.UUID, &schedule.Group, &schedule.Teachers,
		&schedule.Rooms, &schedule.Subject, &schedule.Type,
		&schedule.Location, &schedule.StartTime, &schedule.EndTime,
		&schedule.StartDate, &schedule.EndDate, &schedule.Weekday, &schedule.Link,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &schedule, nil
}

func (r *ScheduleRepository) GetByTeacher(ctx context.Context, firstName, secondName, middleName string, isSession bool) ([]*models.ScheduleData, error) {
	const op = "repository.postgres.ScheduleRepository.GetByTeacher"

	query := `SELECT uuid
			  FROM teachers
			  WHERE first_name = $1 AND second_name = $2`
	var row pgx.Row
	if middleName != "" {
		query += ` AND middle_name = $3`
		row = r.db.QueryRow(ctx, query, firstName, secondName, middleName)
	} else {
		query += ` AND middle_name IS NULL`
		row = r.db.QueryRow(ctx, query, firstName, secondName)
	}

	var teacherUUID uuid.UUID
	err := row.Scan(&teacherUUID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	schedule, err := r.GetByTeacherUUID(ctx, teacherUUID, isSession)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return schedule, nil
}

func (r *ScheduleRepository) GetByTeacherUUID(ctx context.Context, teacherUUID uuid.UUID, isSession bool) ([]*models.ScheduleData, error) {
	const op = "repository.postgres.ScheduleRepository.GetByTeacherUUID"

	query := baseSelectStatement + `
		WHERE schedule.uuid IN (
			SELECT schedule_uuid
			FROM teachers_to_schedule
			WHERE teacher_uuid = $1
		) AND is_session = $2 ` + baseGroupByStatement
	rows, err := r.db.Query(ctx, query, teacherUUID, isSession)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	schedules, err := parseSchedule(&rows)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if len(schedules) == 0 {
		return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return schedules, nil
}

func (r *ScheduleRepository) GetByGroup(ctx context.Context, groupNumber string, isSession bool) ([]*models.ScheduleData, error) {
	const op = "repository.postgres.ScheduleRepository.GetByGroup"

	query := `SELECT uuid
			  FROM groups
			  WHERE number = $1`

	row := r.db.QueryRow(ctx, query, groupNumber)

	var groupUUID uuid.UUID
	err := row.Scan(&groupUUID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	schedule, err := r.GetByGroupUUID(ctx, groupUUID, isSession)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return schedule, nil
}

func (r *ScheduleRepository) GetByGroupUUID(ctx context.Context, groupUUID uuid.UUID, isSession bool) ([]*models.ScheduleData, error) {
	const op = "repository.postgres.ScheduleRepository.GetByGroupUUID"

	query := `SELECT uuid, teachers,
			   rooms, subject_name, subject_type,
			   location, start_time, end_time,
			   start_date, end_date, weekday,
			   link, is_session
			  FROM (` + baseSelectStatement + ` WHERE groups.uuid = $1 AND is_session = $2 ` + baseGroupByStatement + ")"
	rows, err := r.db.Query(ctx, query, groupUUID, isSession)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	schedules, err := parseSchedule(&rows)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if len(schedules) == 0 {
		return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return schedules, nil
}

func (r *ScheduleRepository) GetByRoom(ctx context.Context, roomNumber string, isSession bool) ([]*models.ScheduleData, error) {
	const op = "repository.postgres.ScheduleRepository.GetByRoom"

	query := `SELECT uuid
			  FROM rooms
			  WHERE number = $1`

	row := r.db.QueryRow(ctx, query, roomNumber)

	var roomUUID uuid.UUID
	err := row.Scan(&roomUUID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	schedule, err := r.GetByRoomUUID(ctx, roomUUID, isSession)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return schedule, nil
}

func (r *ScheduleRepository) GetByRoomUUID(ctx context.Context, roomUUID uuid.UUID, isSession bool) ([]*models.ScheduleData, error) {
	const op = "repository.postgres.ScheduleRepository.GetByRoomUUID"

	query := baseSelectStatement + `
		WHERE schedule.uuid IN (
			SELECT schedule_uuid
			FROM rooms_to_schedule
			WHERE room_uuid = $1
		) AND is_session = $2 ` + baseGroupByStatement
	rows, err := r.db.Query(ctx, query, roomUUID, isSession)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	schedules, err := parseSchedule(&rows)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if len(schedules) == 0 {
		return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return schedules, nil
}

func (r *ScheduleRepository) GetBySubject(ctx context.Context, subjectName string, isSession bool) ([]*models.ScheduleData, error) {
	const op = "repository.postgres.ScheduleRepository.GetBySubject"

	query := `SELECT uuid
			  FROM subjects
			  WHERE name = $1`

	row := r.db.QueryRow(ctx, query, subjectName)

	var subjectUUID uuid.UUID
	err := row.Scan(&subjectUUID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	schedule, err := r.GetBySubjectUUID(ctx, subjectUUID, isSession)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return schedule, nil
}

func (r *ScheduleRepository) GetBySubjectUUID(ctx context.Context, subjectUUID uuid.UUID, isSession bool) ([]*models.ScheduleData, error) {
	const op = "repository.postgres.ScheduleRepository.GetBySubjectUUID"

	query := baseSelectStatement + ` WHERE subjects.uuid = $1 AND is_session = $2 ` + baseGroupByStatement
	rows, err := r.db.Query(ctx, query, subjectUUID, isSession)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	schedules, err := parseSchedule(&rows)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if len(schedules) == 0 {
		return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return schedules, nil
}

func (r *ScheduleRepository) GetByLocation(ctx context.Context, locationName string, isSession bool) ([]*models.ScheduleData, error) {
	const op = "repository.postgres.ScheduleRepository.GetByLocation"

	query := `SELECT uuid
			  FROM locations
			  WHERE name = $1`

	row := r.db.QueryRow(ctx, query, locationName)

	var locationUUID uuid.UUID
	err := row.Scan(&locationUUID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	schedule, err := r.GetByLocationUUID(ctx, locationUUID, isSession)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return schedule, nil
}

func (r *ScheduleRepository) GetByLocationUUID(ctx context.Context, locationUUID uuid.UUID, isSession bool) ([]*models.ScheduleData, error) {
	const op = "repository.postgres.ScheduleRepository.GetByLocationUUID"

	query := baseSelectStatement + ` WHERE locations.uuid = $1 AND is_session = $2 ` + baseGroupByStatement
	rows, err := r.db.Query(ctx, query, locationUUID, isSession)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	schedules, err := parseSchedule(&rows)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if len(schedules) == 0 {
		return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return schedules, nil
}

func (r *ScheduleRepository) Update(ctx context.Context, schedule *models.Schedule) error {
	const op = "repository.postgres.ScheduleRepository.Update"

	query := `UPDATE schedule
			  SET group_uuid = $2, subject_uuid = $3, type_uuid = $4,
			      location_uuid = $5, start_time = $6, end_time = $7,
			      start_date = $8, end_date = $9, weekday = $10, link = $11
			  WHERE uuid = $1`
	result, err := r.db.Exec(
		ctx, query, schedule.UUID, schedule.GroupUUID, schedule.SubjectUUID,
		schedule.TypeUUID, schedule.LocationUUID, schedule.StartTime, schedule.EndTime,
		schedule.StartDate, schedule.EndDate, schedule.Weekday, schedule.Link,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return nil
}

func (r *ScheduleRepository) Delete(ctx context.Context, uuid uuid.UUID) error {
	const op = "repository.postgres.ScheduleRepository.Delete"

	query := `DELETE FROM schedule WHERE uuid = $1`
	result, err := r.db.Exec(ctx, query, uuid)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return nil
}

func (r *ScheduleRepository) DeletePairsByGroupWeekdayTime(ctx context.Context, groupUUID uuid.UUID, weekday int, st, sd time.Time, isSession bool) error {
	const op = "repository.postgres.ScheduleRepository.DeletePairsByGroupWeekdayTime"

	query := `DELETE FROM schedule 
       		  WHERE group_uuid = $1 AND weekday = $2 AND start_time = $3 AND start_date = $4 AND is_session = $5`
	result, err := r.db.Exec(ctx, query, groupUUID, weekday, st, sd, isSession)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return nil
}

func (r *ScheduleRepository) DeleteByParams(ctx context.Context, params *models.ScheduleData) error {
	const op = "repository.postgres.ScheduleRepository.DeleteByParams"

	base := `DELETE FROM schedule
			 USING groups g, subjects s, subj_types t, locations l
			 WHERE schedule.group_uuid   = g.uuid
				AND schedule.subject_uuid = s.uuid
				AND schedule.type_uuid    = t.uuid
				AND schedule.location_uuid= l.uuid`

	var (
		conds []string
		args  []interface{}
		idx   = 1
	)

	add := func(expr string, val interface{}) {
		conds = append(conds, fmt.Sprintf(expr, idx))
		args = append(args, val)
		idx++
	}

	if params.Group != "" {
		add("g.number = $%d", params.Group)
	}
	if params.Subject != "" {
		add("s.name = $%d", params.Subject)
	}
	if params.Type != "" {
		add("t.type = $%d", params.Type)
	}
	if params.Location != "" {
		add("l.name = $%d", params.Location)
	}
	if !params.StartTime.IsZero() {
		add("schedule.start_time = $%d", params.StartTime)
	}
	if !params.EndTime.IsZero() {
		add("schedule.end_time = $%d", params.EndTime)
	}
	if !params.StartDate.IsZero() {
		add("schedule.start_date = $%d", params.StartDate)
	}
	if !params.EndDate.IsZero() {
		add("schedule.end_date = $%d", params.EndDate)
	}
	if params.Weekday != 0 {
		add("schedule.weekday = $%d", params.Weekday)
	}
	add("schedule.is_session = $%d", params.IsSession)

	if len(conds) == 0 {
		return fmt.Errorf("%s: no parameters provided for deletion", op)
	}

	query := base + " AND " + strings.Join(conds, " AND ")

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}
	return nil
}
