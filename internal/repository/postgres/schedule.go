package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"raspyx/internal/domain/models"
	"raspyx/internal/repository"
)

type ScheduleRepository struct {
	db *pgx.Conn
}

func NewScheduleRepository(db *pgx.Conn) *ScheduleRepository {
	return &ScheduleRepository{db: db}
}

func (r *ScheduleRepository) Create(ctx context.Context, schedule *models.Schedule) error {
	const op = "repository.postgres.ScheduleRepository.Create"

	query := `INSERT INTO schedule (uuid, teacher_uuid, group_uuid,
                      				room_uuid, subject_uuid, type_uuid,
                      				location_uuid, start_time, end_time,
                      				start_date, end_date, weekday, link) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`
	_, err := r.db.Exec(
		ctx, query, schedule.UUID, schedule.TeacherUUID, schedule.GroupUUID,
		schedule.RoomUUID, schedule.SubjectUUID, schedule.TypeUUID,
		schedule.LocationUUID, schedule.StartTime, schedule.EndTime,
		schedule.StartDate, schedule.EndDate, schedule.Weekday, schedule.Link,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *ScheduleRepository) GetByUUID(ctx context.Context, uuid uuid.UUID) (*models.Schedule, error) {
	const op = "repository.postgres.ScheduleRepository.GetByUUID"

	query := `SELECT uuid, teacher_uuid, group_uuid,
					 room_uuid, subject_uuid, type_uuid,
					 location_uuid, start_time, end_time,
					 start_date, end_date, weekday, link
			  FROM schedule
			  WHERE uuid = $1`
	row := r.db.QueryRow(ctx, query, uuid)
	var schedule models.Schedule
	err := row.Scan(
		&schedule.UUID, &schedule.TeacherUUID, &schedule.GroupUUID,
		&schedule.RoomUUID, &schedule.SubjectUUID, &schedule.TypeUUID,
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

func (r *ScheduleRepository) GetByTeacher(ctx context.Context, firstName, secondName, middleName string) ([]*models.Schedule, error) {
	const op = "repository.postgres.ScheduleRepository.GetByTeacher"

	query := `SELECT uuid
			  FROM teachers
			  WHERE first_name = $1 AND second_name = $2 AND middle_name = $3`

	row := r.db.QueryRow(ctx, query, firstName, secondName, middleName)

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

func (r *ScheduleRepository) GetByTeacherUUID(ctx context.Context, teacherUUID uuid.UUID) ([]*models.Schedule, error) {
	const op = "repository.postgres.ScheduleRepository.GetByTeacherUUID"

	query := `SELECT uuid, teacher_uuid, group_uuid,
					 room_uuid, subject_uuid, type_uuid,
					 location_uuid, start_time, end_time,
					 start_date, end_date, weekday, link
			  FROM schedule
			  WHERE teacher_uuid = $1`
	rows, err := r.db.Query(ctx, query, teacherUUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	schedules, err := parseSchedule(&rows)

	if len(schedules) == 0 {
		return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return schedules, nil
}

func (r *ScheduleRepository) GetByGroup(ctx context.Context, groupNumber string) ([]*models.Schedule, error) {
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

func (r *ScheduleRepository) GetByGroupUUID(ctx context.Context, groupUUID uuid.UUID) ([]*models.Schedule, error) {
	const op = "repository.postgres.ScheduleRepository.GetByGroupUUID"

	query := `SELECT uuid, teacher_uuid, group_uuid,
					 room_uuid, subject_uuid, type_uuid,
					 location_uuid, start_time, end_time,
					 start_date, end_date, weekday, link
			  FROM schedule
			  WHERE group_uuid = $1`
	rows, err := r.db.Query(ctx, query, groupUUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	schedules, err := parseSchedule(&rows)

	if len(schedules) == 0 {
		return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return schedules, nil
}

func (r *ScheduleRepository) GetByRoom(ctx context.Context, roomNumber string) ([]*models.Schedule, error) {
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

func (r *ScheduleRepository) GetByRoomUUID(ctx context.Context, roomUUID uuid.UUID) ([]*models.Schedule, error) {
	const op = "repository.postgres.ScheduleRepository.GetByRoomUUID"

	query := `SELECT uuid, teacher_uuid, group_uuid,
					 room_uuid, subject_uuid, type_uuid,
					 location_uuid, start_time, end_time,
					 start_date, end_date, weekday, link
			  FROM schedule
			  WHERE room_uuid = $1`
	rows, err := r.db.Query(ctx, query, roomUUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	schedules, err := parseSchedule(&rows)

	if len(schedules) == 0 {
		return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return schedules, nil
}

func (r *ScheduleRepository) GetBySubject(ctx context.Context, subjectName string) ([]*models.Schedule, error) {
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

func (r *ScheduleRepository) GetBySubjectUUID(ctx context.Context, subjectUUID uuid.UUID) ([]*models.Schedule, error) {
	const op = "repository.postgres.ScheduleRepository.GetBySubjectUUID"

	query := `SELECT uuid, teacher_uuid, group_uuid,
					 room_uuid, subject_uuid, type_uuid,
					 location_uuid, start_time, end_time,
					 start_date, end_date, weekday, link
			  FROM schedule
			  WHERE subject_uuid = $1`
	rows, err := r.db.Query(ctx, query, subjectUUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	schedules, err := parseSchedule(&rows)

	if len(schedules) == 0 {
		return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return schedules, nil
}

func (r *ScheduleRepository) GetByLocation(ctx context.Context, locationName string) ([]*models.Schedule, error) {
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

func (r *ScheduleRepository) GetByLocationUUID(ctx context.Context, locationUUID uuid.UUID) ([]*models.Schedule, error) {
	const op = "repository.postgres.ScheduleRepository.GetByLocationUUID"

	query := `SELECT uuid, teacher_uuid, group_uuid,
					 room_uuid, subject_uuid, type_uuid,
					 location_uuid, start_time, end_time,
					 start_date, end_date, weekday, link
			  FROM schedule
			  WHERE location_uuid = $1`
	rows, err := r.db.Query(ctx, query, locationUUID)
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
			  SET teacher_uuid = $2, group_uuid = $3, room_uuid = $4,
			      subject_uuid = $5, type_uuid = $6, location_uuid = $7,
			      start_time = $8, end_time = $9, start_date = $10,
			      end_date = $11, weekday = $12, link = $13
			  WHERE uuid = $1`
	result, err := r.db.Exec(
		ctx, query, schedule.UUID, schedule.TeacherUUID, schedule.GroupUUID,
		schedule.RoomUUID, schedule.SubjectUUID, schedule.TypeUUID,
		schedule.LocationUUID, schedule.StartTime, schedule.EndTime,
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

func parseSchedule(rows *pgx.Rows) ([]*models.Schedule, error) {
	const op = "repository.postgres.ScheduleRepository.parseSchedule"

	var schedules []*models.Schedule
	for (*rows).Next() {
		var schedule models.Schedule
		err := (*rows).Scan(
			&schedule.UUID, &schedule.TeacherUUID, &schedule.GroupUUID,
			&schedule.RoomUUID, &schedule.SubjectUUID, &schedule.TypeUUID,
			&schedule.LocationUUID, &schedule.StartTime, &schedule.EndTime,
			&schedule.StartDate, &schedule.EndDate, &schedule.Weekday, &schedule.Link,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		schedules = append(schedules, &schedule)
	}

	return schedules, nil
}
