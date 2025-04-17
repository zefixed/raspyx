package postgres

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"raspyx/internal/domain/models"
	"raspyx/internal/repository"
	"strings"
)

type TeachersToScheduleRepository struct {
	db *pgxpool.Pool
}

func NewTeachersToScheduleRepository(db *pgxpool.Pool) *TeachersToScheduleRepository {
	return &TeachersToScheduleRepository{db: db}
}

func (r *TeachersToScheduleRepository) Create(ctx context.Context, teachersToSchedule *models.TeachersToSchedule) error {
	const op = "repository.postgres.TeachersToScheduleRepository.Create"

	query := `INSERT INTO teachers_to_schedule (teacher_uuid, schedule_uuid) 
			  VALUES ($1, $2)`
	_, err := r.db.Exec(ctx, query, teachersToSchedule.TeacherUUID, teachersToSchedule.ScheduleUUID)
	if err != nil {
		if strings.Contains(err.Error(), "23505") {
			return fmt.Errorf("%s: %w", op, repository.ErrExist)
		} else if strings.Contains(err.Error(), "23503") {
			return fmt.Errorf("%s: %w", op, repository.ErrNotExist)
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *TeachersToScheduleRepository) Get(ctx context.Context) ([]*models.TeachersToSchedule, error) {
	const op = "repository.postgres.TeachersToScheduleRepository.Get"

	query := `SELECT teacher_uuid, schedule_uuid
			  FROM teachers_to_schedule`
	rows, err := r.db.Query(ctx, query)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var relations []*models.TeachersToSchedule
	for rows.Next() {
		var relation models.TeachersToSchedule
		err := rows.Scan(&relation.TeacherUUID, &relation.ScheduleUUID)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		relations = append(relations, &relation)
	}

	return relations, nil
}

func (r *TeachersToScheduleRepository) GetByTeacherUUID(ctx context.Context, teacherUUID uuid.UUID) ([]*models.TeachersToSchedule, error) {
	const op = "repository.postgres.TeachersToScheduleRepository.GetByTeacherUUID"
	query := `SELECT teacher_uuid, schedule_uuid
			  FROM teachers_to_schedule
			  WHERE teacher_uuid = $1`
	rows, err := r.db.Query(ctx, query, teacherUUID)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var relations []*models.TeachersToSchedule
	for rows.Next() {
		var relation models.TeachersToSchedule
		err := rows.Scan(&relation.TeacherUUID, &relation.ScheduleUUID)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		relations = append(relations, &relation)
	}

	if len(relations) == 0 {
		return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return relations, nil
}

func (r *TeachersToScheduleRepository) GetByScheduleUUID(ctx context.Context, scheduleUUID uuid.UUID) ([]*models.TeachersToSchedule, error) {
	const op = "repository.postgres.TeachersToScheduleRepository.GetByScheduleUUID"
	query := `SELECT teacher_uuid, schedule_uuid
			  FROM teachers_to_schedule
			  WHERE schedule_uuid = $1`
	rows, err := r.db.Query(ctx, query, scheduleUUID)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var relations []*models.TeachersToSchedule
	for rows.Next() {
		var relation models.TeachersToSchedule
		err := rows.Scan(&relation.TeacherUUID, &relation.ScheduleUUID)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		relations = append(relations, &relation)
	}

	if len(relations) == 0 {
		return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return relations, nil
}

func (r *TeachersToScheduleRepository) Delete(ctx context.Context, teachersToSchedule *models.TeachersToSchedule) error {
	const op = "repository.postgres.TeachersToScheduleRepository.Delete"
	query := `DELETE FROM teachers_to_schedule
			  WHERE teacher_uuid = $1 AND schedule_uuid = $2`
	result, err := r.db.Exec(ctx, query, teachersToSchedule.TeacherUUID, teachersToSchedule.ScheduleUUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return nil
}
