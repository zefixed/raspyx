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
	"strings"
)

type LocationRepository struct {
	db *pgxpool.Pool
}

func NewLocationRepository(db *pgxpool.Pool) *LocationRepository {
	return &LocationRepository{db: db}
}

func (r *LocationRepository) Create(ctx context.Context, location *models.Location) error {
	const op = "repository.postgres.LocationRepository.Create"

	query := `INSERT INTO locations (uuid, name) 
			  VALUES ($1, $2)`
	_, err := r.db.Exec(ctx, query, location.UUID, location.Name)

	if err != nil {
		if strings.Contains(err.Error(), "23505") {
			return fmt.Errorf("%s: %w", op, repository.ErrExist)
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil

}

func (r *LocationRepository) Get(ctx context.Context) ([]*models.Location, error) {
	const op = "repository.postgres.LocationRepository.Get"

	query := `SELECT uuid, name
			  FROM locations`
	rows, err := r.db.Query(ctx, query)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var locations []*models.Location
	for rows.Next() {
		var location models.Location
		err := rows.Scan(&location.UUID, &location.Name)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		locations = append(locations, &location)
	}

	return locations, nil
}

func (r *LocationRepository) GetByUUID(ctx context.Context, uuid uuid.UUID) (*models.Location, error) {
	const op = "repository.postgres.LocationRepository.GetByUUID"

	query := `SELECT uuid, name
			  FROM locations
			  WHERE uuid = $1`

	row := r.db.QueryRow(ctx, query, uuid)
	var location models.Location
	err := row.Scan(&location.UUID, &location.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &location, nil
}

func (r *LocationRepository) GetByName(ctx context.Context, name string) (*models.Location, error) {
	const op = "repository.postgres.LocationRepository.GetByName"

	query := `SELECT uuid, name
			  FROM locations
			  WHERE name = $1`

	row := r.db.QueryRow(ctx, query, name)
	var location models.Location
	err := row.Scan(&location.UUID, &location.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &location, nil
}

func (r *LocationRepository) Update(ctx context.Context, location *models.Location) error {
	const op = "repository.postgres.LocationRepository.Update"

	query := `UPDATE locations
			  SET name = $1
			  WHERE uuid = $2`

	result, err := r.db.Exec(ctx, query, location.Name, location.UUID)
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

func (r *LocationRepository) Delete(ctx context.Context, uuid uuid.UUID) error {
	const op = "repository.postgres.LocationRepository.Delete"

	query := `DELETE FROM locations WHERE uuid = $1`
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
