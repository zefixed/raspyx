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
	"strings"
)

type SubjectTypeRepository struct {
	db *pgx.Conn
}

func NewSubjectTypeRepository(db *pgx.Conn) *SubjectTypeRepository {
	return &SubjectTypeRepository{db: db}
}

func (r *SubjectTypeRepository) Create(ctx context.Context, subjectType *models.SubjectType) error {
	const op = "repository.postgres.SubjectTypeRepository.Create"

	query := `INSERT INTO subj_types (uuid, type) 
			  VALUES ($1, $2)`
	_, err := r.db.Exec(ctx, query, subjectType.UUID, subjectType.Type)
	if err != nil {
		if strings.Contains(err.Error(), "23505") {
			return fmt.Errorf("%s: %w", op, repository.ErrExist)
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *SubjectTypeRepository) Get(ctx context.Context) ([]*models.SubjectType, error) {
	const op = "repository.postgres.SubjectTypeRepository.GetByUUID"

	query := `SELECT uuid, type
			  FROM subj_types`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var subjTypes []*models.SubjectType
	for rows.Next() {
		var subjType models.SubjectType
		err := rows.Scan(&subjType.UUID, &subjType.Type)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		subjTypes = append(subjTypes, &subjType)
	}

	return subjTypes, nil
}

func (r *SubjectTypeRepository) GetByUUID(ctx context.Context, uuid uuid.UUID) (*models.SubjectType, error) {
	const op = "repository.postgres.SubjectTypeRepository.GetByUUID"

	query := `SELECT uuid, type
			  FROM subj_types
			  WHERE uuid = $1`
	row := r.db.QueryRow(ctx, query, uuid)
	var subjType models.SubjectType
	err := row.Scan(&subjType.UUID, &subjType.Type)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &subjType, nil
}

func (r *SubjectTypeRepository) GetByType(ctx context.Context, subjectType string) (*models.SubjectType, error) {
	const op = "repository.postgres.SubjectTypeRepository.GetByType"

	query := `SELECT uuid, type
			  FROM subj_types
			  WHERE type = $1`
	row := r.db.QueryRow(ctx, query, subjectType)
	var subjType models.SubjectType
	err := row.Scan(&subjType.UUID, &subjType.Type)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &subjType, nil
}

func (r *SubjectTypeRepository) Update(ctx context.Context, subjectType *models.SubjectType) error {
	const op = "repository.postgres.SubjectTypeRepository.Update"

	query := `UPDATE subj_types
			  SET type = $1
			  WHERE uuid = $2`
	result, err := r.db.Exec(ctx, query, subjectType.Type, subjectType.UUID)
	if err != nil {
		if strings.Contains(err.Error(), "23505") {
			return fmt.Errorf("%s: %w", op, repository.ErrExist)
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return nil
}

func (r *SubjectTypeRepository) Delete(ctx context.Context, uuid uuid.UUID) error {
	const op = "repository.postgres.SubjectTypeRepository.Delete"

	query := `DELETE FROM subj_types WHERE uuid = $1`
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
