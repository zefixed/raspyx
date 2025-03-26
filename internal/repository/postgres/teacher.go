package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"raspyx/internal/domain/models"
)

type TeacherRepository struct {
	db *pgx.Conn
}

func NewTeacherRepository(db *pgx.Conn) *TeacherRepository {
	return &TeacherRepository{db: db}
}

func (r *TeacherRepository) Create(ctx context.Context, teacher *models.Teacher) error {
	query := `INSERT INTO teachers (uuid, first_name, second_name, middle_name) 
			  VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(ctx, query,
		teacher.UUID,
		teacher.FirstName,
		teacher.SecondName,
		teacher.MiddleName,
	)
	return err
}

func (r *TeacherRepository) GetByUUID(ctx context.Context, uuid uuid.UUID) (*models.Teacher, error) {
	query := `SELECT uuid, first_name, second_name, middle_name 
			  FROM teachers 
			  WHERE uuid = $1`

	row := r.db.QueryRow(ctx, query, uuid)

	var teacher models.Teacher
	err := row.Scan(&teacher.UUID, &teacher.FirstName, &teacher.SecondName, &teacher.MiddleName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("teacher with uuid: %v not found", uuid)
		}
		return nil, err
	}

	return &teacher, nil
}

func (r *TeacherRepository) GetByFullName(ctx context.Context, fn string) (*models.Teacher, error) {
	query := `SELECT uuid, first_name, second_name, middle_name 
			  FROM teachers 
			  WHERE CONCAT(first_name, ' ', second_name, ' ', middle_name) = $1`

	row := r.db.QueryRow(ctx, query, fn)

	var teacher models.Teacher
	err := row.Scan(&teacher.UUID, &teacher.FirstName, &teacher.SecondName, &teacher.MiddleName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("teacher with name: %v not found", fn)
		}
		return nil, err
	}

	return &teacher, nil
}
func (r *TeacherRepository) Update(ctx context.Context, teacher *models.Teacher) error {
	query := `UPDATE teachers 
	          SET first_name = $1, second_name = $2, middle_name = $3 
	          WHERE uuid = $4`

	result, err := r.db.Exec(ctx, query, teacher.FirstName, teacher.SecondName, teacher.MiddleName, teacher.UUID)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		return fmt.Errorf("teacher with UUID %s not found", teacher.UUID)
	}

	return nil
}

func (r *TeacherRepository) Delete(ctx context.Context, uuid uuid.UUID) error {
	query := `DELETE FROM teachers WHERE uuid = $1`

	result, err := r.db.Exec(ctx, query, uuid)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		return fmt.Errorf("teacher with UUID %s not found", uuid)
	}

	return nil
}
