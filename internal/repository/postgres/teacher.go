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

type TeacherRepository struct {
	db *pgx.Conn
}

func NewTeacherRepository(db *pgx.Conn) *TeacherRepository {
	return &TeacherRepository{db: db}
}

func (r *TeacherRepository) Create(ctx context.Context, teacher *models.Teacher) error {
	const op = "repository.postgres.TeacherRepository.Create"

	query := `INSERT INTO teachers (uuid, first_name, second_name, middle_name) 
			  VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(ctx, query,
		teacher.UUID,
		teacher.FirstName,
		teacher.SecondName,
		teacher.MiddleName,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *TeacherRepository) Get(ctx context.Context) ([]*models.Teacher, error) {
	const op = "repository.postgres.TeacherRepository.Get"

	query := `SELECT uuid, first_name, second_name, middle_name 
			  FROM teachers`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var teachers []*models.Teacher
	for rows.Next() {
		var teacher models.Teacher
		err := rows.Scan(&teacher.UUID, &teacher.FirstName, &teacher.SecondName, &teacher.MiddleName)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		teachers = append(teachers, &teacher)
	}

	return teachers, nil
}

func (r *TeacherRepository) GetByUUID(ctx context.Context, uuid uuid.UUID) (*models.Teacher, error) {
	const op = "repository.postgres.TeacherRepository.GetByUUID"

	query := `SELECT uuid, first_name, second_name, middle_name 
			  FROM teachers 
			  WHERE uuid = $1`

	row := r.db.QueryRow(ctx, query, uuid)

	var teacher models.Teacher
	err := row.Scan(&teacher.UUID, &teacher.FirstName, &teacher.SecondName, &teacher.MiddleName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &teacher, nil
}

func (r *TeacherRepository) GetByFullName(ctx context.Context, fn string) ([]*models.Teacher, error) {
	const op = "repository.postgres.TeacherRepository.GetByFullName"

	query := `SELECT uuid, first_name, second_name, middle_name 
			  FROM teachers 
			  WHERE CONCAT(second_name, ' ', first_name, ' ', middle_name) = $1`

	rows, err := r.db.Query(ctx, query, fn)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var teachers []*models.Teacher
	for rows.Next() {
		var teacher models.Teacher
		err := rows.Scan(&teacher.UUID, &teacher.FirstName, &teacher.SecondName, &teacher.MiddleName)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
			}
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		teachers = append(teachers, &teacher)
	}

	if len(teachers) == 0 {
		return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return teachers, nil
}
func (r *TeacherRepository) Update(ctx context.Context, teacher *models.Teacher) error {
	const op = "repository.postgres.TeacherRepository.Update"

	query := `UPDATE teachers 
	          SET first_name = $1, second_name = $2, middle_name = $3 
	          WHERE uuid = $4`

	result, err := r.db.Exec(ctx, query, teacher.FirstName, teacher.SecondName, teacher.MiddleName, teacher.UUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return nil
}

func (r *TeacherRepository) Delete(ctx context.Context, uuid uuid.UUID) error {
	const op = "repository.postgres.TeacherRepository.Delete"

	query := `DELETE FROM teachers WHERE uuid = $1`

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
