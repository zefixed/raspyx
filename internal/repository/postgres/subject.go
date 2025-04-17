package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"raspyx/internal/domain/models"
	"raspyx/internal/repository"
)

type SubjectRepository struct {
	db *pgxpool.Pool
}

func NewSubjectRepository(db *pgxpool.Pool) *SubjectRepository {
	return &SubjectRepository{db: db}
}

func (r *SubjectRepository) Create(ctx context.Context, subject *models.Subject) error {
	const op = "repository.postgres.SubjectRepository.Create"

	query := `INSERT INTO subjects (uuid, name) 
			  VALUES ($1, $2)`
	_, err := r.db.Exec(ctx, query, subject.UUID, subject.Name)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *SubjectRepository) Get(ctx context.Context) ([]*models.Subject, error) {
	const op = "repository.postgres.SubjectRepository.Get"

	query := `SELECT uuid, name
			  FROM subjects`
	rows, err := r.db.Query(ctx, query)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var subjects []*models.Subject
	for rows.Next() {
		var subject models.Subject
		err := rows.Scan(&subject.UUID, &subject.Name)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		subjects = append(subjects, &subject)
	}

	return subjects, nil
}

func (r *SubjectRepository) GetByUUID(ctx context.Context, uuid uuid.UUID) (*models.Subject, error) {
	const op = "repository.postgres.SubjectRepository.GetByUUID"

	query := `SELECT uuid, name
			  FROM subjects
			  WHERE uuid = $1`
	row := r.db.QueryRow(ctx, query, uuid)
	var subject models.Subject
	err := row.Scan(&subject.UUID, &subject.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &subject, nil
}

func (r *SubjectRepository) GetByName(ctx context.Context, name string) ([]*models.Subject, error) {
	const op = "repository.postgres.SubjectRepository.GetByName"

	query := `SELECT uuid, name
			  FROM subjects
			  WHERE name = $1`
	rows, err := r.db.Query(ctx, query, name)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var subjects []*models.Subject
	for rows.Next() {
		var subject models.Subject
		err := rows.Scan(&subject.UUID, &subject.Name)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		subjects = append(subjects, &subject)
	}

	if len(subjects) == 0 {
		return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return subjects, nil
}

func (r *SubjectRepository) Update(ctx context.Context, subject *models.Subject) error {
	const op = "repository.postgres.SubjectRepository.Update"

	query := `UPDATE subjects
			  SET name = $1
			  WHERE uuid = $2`
	result, err := r.db.Exec(ctx, query, subject.Name, subject.UUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return nil
}

func (r *SubjectRepository) Delete(ctx context.Context, uuid uuid.UUID) error {
	const op = "repository.postgres.SubjectRepository.Delete"

	query := `DELETE FROM subjects WHERE uuid = $1`
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
