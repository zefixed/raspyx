package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"raspyx/internal/domain/models"
	"raspyx/internal/repository"
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
	for (*rows).Next() {
		var schedule models.ScheduleData
		err := (*rows).Scan(
			&schedule.UUID, &schedule.Group, &schedule.Teachers,
			&schedule.Rooms, &schedule.Subject, &schedule.Type,
			&schedule.Location, &schedule.StartTime, &schedule.EndTime,
			&schedule.StartDate, &schedule.EndDate, &schedule.Weekday, &schedule.Link,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		schedules = append(schedules, &schedule)
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
			COALESCE(schedule.link, '') AS "link"
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

	query := `INSERT INTO schedule (uuid, group_uuid, subject_uuid,
                      				type_uuid, location_uuid, start_time,
                      				end_time, start_date, end_date, weekday, link)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := r.db.Exec(
		ctx, query, schedule.UUID, schedule.GroupUUID, schedule.SubjectUUID,
		schedule.TypeUUID, schedule.LocationUUID, schedule.StartTime, schedule.EndTime,
		schedule.StartDate, schedule.EndDate, schedule.Weekday, schedule.Link,
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

	if len(schedules) == 0 {
		return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return schedules, nil
}

func (r *ScheduleRepository) GetForUpdate(ctx context.Context, uuid uuid.UUID) (*models.Schedule, error) {
	const op = "repository.postgres.ScheduleRepository.GetForUpdate"

	query := `SELECT uuid, group_uuid, subject_uuid, type_uuid, location_uuid, start_time,
					 end_time, start_date, end_date, weekday, link
			  FROM schedule
			  WHERE uuid = $1`
	row := r.db.QueryRow(ctx, query, uuid)

	var schedule models.Schedule
	err := row.Scan(
		&schedule.UUID, &schedule.GroupUUID, &schedule.SubjectUUID, &schedule.TypeUUID,
		&schedule.LocationUUID, &schedule.StartTime, &schedule.EndTime,
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

func (r *ScheduleRepository) GetByTeacher(ctx context.Context, firstName, secondName, middleName string) ([]*models.ScheduleData, error) {
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

	schedule, err := r.GetByTeacherUUID(ctx, teacherUUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return schedule, nil
}

func (r *ScheduleRepository) GetByTeacherUUID(ctx context.Context, teacherUUID uuid.UUID) ([]*models.ScheduleData, error) {
	const op = "repository.postgres.ScheduleRepository.GetByTeacherUUID"

	query := baseSelectStatement + `
		WHERE schedule.uuid IN (
			SELECT schedule_uuid
			FROM teachers_to_schedule
			WHERE teacher_uuid = $1
		) ` + baseGroupByStatement
	rows, err := r.db.Query(ctx, query, teacherUUID)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	schedules, err := parseSchedule(&rows)

	if len(schedules) == 0 {
		return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return schedules, nil
}

func (r *ScheduleRepository) GetByGroup(ctx context.Context, groupNumber string) ([]*models.ScheduleData, error) {
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

	schedule, err := r.GetByGroupUUID(ctx, groupUUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return schedule, nil
}

func (r *ScheduleRepository) GetByGroupUUID(ctx context.Context, groupUUID uuid.UUID) ([]*models.ScheduleData, error) {
	const op = "repository.postgres.ScheduleRepository.GetByGroupUUID"

	query := baseSelectStatement + ` WHERE groups.uuid = $1 ` + baseGroupByStatement
	rows, err := r.db.Query(ctx, query, groupUUID)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	schedules, err := parseSchedule(&rows)

	if len(schedules) == 0 {
		return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return schedules, nil
}

func (r *ScheduleRepository) GetByRoom(ctx context.Context, roomNumber string) ([]*models.ScheduleData, error) {
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

	schedule, err := r.GetByRoomUUID(ctx, roomUUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return schedule, nil
}

func (r *ScheduleRepository) GetByRoomUUID(ctx context.Context, roomUUID uuid.UUID) ([]*models.ScheduleData, error) {
	const op = "repository.postgres.ScheduleRepository.GetByRoomUUID"

	query := baseSelectStatement + `
		WHERE schedule.uuid IN (
			SELECT schedule_uuid
			FROM rooms_to_schedule
			WHERE room_uuid = $1
		) ` + baseGroupByStatement
	rows, err := r.db.Query(ctx, query, roomUUID)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	schedules, err := parseSchedule(&rows)

	if len(schedules) == 0 {
		return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return schedules, nil
}

func (r *ScheduleRepository) GetBySubject(ctx context.Context, subjectName string) ([]*models.ScheduleData, error) {
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

	schedule, err := r.GetBySubjectUUID(ctx, subjectUUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return schedule, nil
}

func (r *ScheduleRepository) GetBySubjectUUID(ctx context.Context, subjectUUID uuid.UUID) ([]*models.ScheduleData, error) {
	const op = "repository.postgres.ScheduleRepository.GetBySubjectUUID"

	query := baseSelectStatement + ` WHERE subjects.uuid = $1 ` + baseGroupByStatement
	rows, err := r.db.Query(ctx, query, subjectUUID)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	schedules, err := parseSchedule(&rows)

	if len(schedules) == 0 {
		return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return schedules, nil
}

func (r *ScheduleRepository) GetByLocation(ctx context.Context, locationName string) ([]*models.ScheduleData, error) {
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

	schedule, err := r.GetByLocationUUID(ctx, locationUUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return schedule, nil
}

func (r *ScheduleRepository) GetByLocationUUID(ctx context.Context, locationUUID uuid.UUID) ([]*models.ScheduleData, error) {
	const op = "repository.postgres.ScheduleRepository.GetByLocationUUID"

	query := baseSelectStatement + ` WHERE locations.uuid = $1 ` + baseGroupByStatement
	rows, err := r.db.Query(ctx, query, locationUUID)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	schedules, err := parseSchedule(&rows)

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

func (r *ScheduleRepository) DeletePairsByGroupWeekdayTime(ctx context.Context, groupUUID uuid.UUID, weekday int, t time.Time) error {
	const op = "repository.postgres.ScheduleRepository.DeletePairsByGroupWeekdayTime"

	query := `DELETE FROM schedule 
       		  WHERE group_uuid = $1 AND weekday = $2 AND start_time = $3`
	result, err := r.db.Exec(ctx, query, groupUUID, weekday, t)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return nil
}
