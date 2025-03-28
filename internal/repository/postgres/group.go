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

type GroupRepository struct {
	db *pgx.Conn
}

func NewGroupRepository(db *pgx.Conn) *GroupRepository {
	return &GroupRepository{db: db}
}

func (r *GroupRepository) Create(ctx context.Context, group *models.Group) error {
	const op = "repository.postgres.GroupRepository.Create"

	query := `INSERT INTO groups (uuid, number) 
			  VALUES ($1, $2)`
	_, err := r.db.Exec(ctx, query, group.UUID, group.Number)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
func (r *GroupRepository) GetByUUID(ctx context.Context, uuid uuid.UUID) (*models.Group, error) {
	const op = "repository.postgres.GroupRepository.GetByUUID"

	query := `SELECT uuid, number
			  FROM groups
			  WHERE uuid = $1`
	row := r.db.QueryRow(ctx, query, uuid)

	var group models.Group
	err := row.Scan(&group.UUID, &group.Number)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &group, nil
}
func (r *GroupRepository) GetByNumber(ctx context.Context, number string) (*models.Group, error) {
	const op = "repository.postgres.GroupRepository.GetByNumber"

	query := `SELECT uuid, number
			  FROM groups
			  WHERE number = $1`
	row := r.db.QueryRow(ctx, query, number)

	var group models.Group
	err := row.Scan(&group.UUID, &group.Number)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &group, nil
}
func (r *GroupRepository) Update(ctx context.Context, group *models.Group) error {
	const op = "repository.postgres.GroupRepository.Update"

	query := `UPDATE groups
			  SET number = $1
			  WHERE uuid = $2`
	result, err := r.db.Exec(ctx, query, group.Number, group.UUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowAffected := result.RowsAffected()
	if rowAffected == 0 {
		return fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return nil
}
func (r *GroupRepository) Delete(ctx context.Context, uuid uuid.UUID) error {
	const op = "repository.postgres.GroupRepository.Update"

	query := `DELETE FROM groups
			  WHERE uuid = $1`
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
