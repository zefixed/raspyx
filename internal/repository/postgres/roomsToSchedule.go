package postgres

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"raspyx/internal/domain/models"
	"raspyx/internal/repository"
	"strings"
)

type RoomsToScheduleRepository struct {
	db *pgx.Conn
}

func NewRoomsToScheduleRepository(db *pgx.Conn) *RoomsToScheduleRepository {
	return &RoomsToScheduleRepository{db: db}
}

func (r *RoomsToScheduleRepository) Create(ctx context.Context, roomsToSchedule *models.RoomsToSchedule) error {
	const op = "repository.postgres.RoomsToScheduleRepository.Create"

	query := `INSERT INTO rooms_to_schedule (room_uuid, schedule_uuid) 
			  VALUES ($1, $2)`
	_, err := r.db.Exec(ctx, query, roomsToSchedule.RoomUUID, roomsToSchedule.ScheduleUUID)
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

func (r *RoomsToScheduleRepository) Get(ctx context.Context) ([]*models.RoomsToSchedule, error) {
	const op = "repository.postgres.RoomsToScheduleRepository.Get"

	query := `SELECT room_uuid, schedule_uuid
			  FROM rooms_to_schedule`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var relations []*models.RoomsToSchedule
	for rows.Next() {
		var relation models.RoomsToSchedule
		err := rows.Scan(&relation.RoomUUID, &relation.ScheduleUUID)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		relations = append(relations, &relation)
	}

	return relations, nil
}

func (r *RoomsToScheduleRepository) GetByRoomUUID(ctx context.Context, roomUUID uuid.UUID) ([]*models.RoomsToSchedule, error) {
	const op = "repository.postgres.RoomsToScheduleRepository.GetByRoomUUID"

	query := `SELECT room_uuid, schedule_uuid
			  FROM rooms_to_schedule
			  WHERE room_uuid = $1`
	rows, err := r.db.Query(ctx, query, roomUUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var relations []*models.RoomsToSchedule
	for rows.Next() {
		var relation models.RoomsToSchedule
		err := rows.Scan(&relation.RoomUUID, &relation.ScheduleUUID)
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

func (r *RoomsToScheduleRepository) GetByScheduleUUID(ctx context.Context, scheduleUUID uuid.UUID) ([]*models.RoomsToSchedule, error) {
	const op = "repository.postgres.RoomsToScheduleRepository.GetByScheduleUUID"

	query := `SELECT room_uuid, schedule_uuid
			  FROM rooms_to_schedule
			  WHERE schedule_uuid = $1`
	rows, err := r.db.Query(ctx, query, scheduleUUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var relations []*models.RoomsToSchedule
	for rows.Next() {
		var relation models.RoomsToSchedule
		err := rows.Scan(&relation.RoomUUID, &relation.ScheduleUUID)
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

func (r *RoomsToScheduleRepository) Delete(ctx context.Context, roomsToSchedule *models.RoomsToSchedule) error {
	const op = "repository.postgres.RoomsToScheduleRepository.Delete"

	query := `DELETE FROM rooms_to_schedule
			  WHERE room_uuid = $1 AND schedule_uuid = $2`
	result, err := r.db.Exec(ctx, query, roomsToSchedule.RoomUUID, roomsToSchedule.ScheduleUUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return nil
}
