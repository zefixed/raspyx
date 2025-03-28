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

type RoomRepository struct {
	db *pgx.Conn
}

func NewRoomRepository(db *pgx.Conn) *RoomRepository {
	return &RoomRepository{db: db}
}

func (r *RoomRepository) Create(ctx context.Context, room *models.Room) error {
	const op = "repository.postgres.RoomRepository.Create"

	query := `INSERT INTO rooms (uuid, number)
			  VALUES ($1, $2)`
	_, err := r.db.Exec(ctx, query, room.UUID, room.Number)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
func (r *RoomRepository) GetByUUID(ctx context.Context, uuid uuid.UUID) (*models.Room, error) {
	const op = "repository.postgres.RoomRepository.GetByUUID"

	query := `SELECT uuid, number
			  FROM rooms
			  WHERE uuid = $1`
	row := r.db.QueryRow(ctx, query, uuid)
	var room models.Room
	err := row.Scan(&room.UUID, &room.Number)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &room, nil
}
func (r *RoomRepository) GetByNumber(ctx context.Context, number string) (*models.Room, error) {
	const op = "repository.postgres.RoomRepository.GetByNumber"

	query := `SELECT uuid, number
			  FROM rooms
			  WHERE number = $1`
	row := r.db.QueryRow(ctx, query, number)
	var room models.Room
	err := row.Scan(&room.UUID, &room.Number)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &room, nil
}
func (r *RoomRepository) Update(ctx context.Context, room *models.Room) error {
	const op = "repository.postgres.RoomRepository.Update"

	query := `UPDATE rooms
			  SET number = $1
			  WHERE uuid = $2`
	result, err := r.db.Exec(ctx, query, room.Number, room.UUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return nil
}
func (r *RoomRepository) Delete(ctx context.Context, uuid uuid.UUID) error {
	const op = "repository.postgres.RoomRepository.Delete"

	query := `DELETE FROM rooms WHERE uuid = $1`
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
