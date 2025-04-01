package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"raspyx/internal/domain/interfaces"
	"raspyx/internal/domain/models"
	"raspyx/internal/domain/services"
	"raspyx/internal/dto"
)

type RoomUseCase struct {
	repo interfaces.RoomRepository
	svc  services.RoomService
}

func NewRoomUseCase(repo interfaces.RoomRepository, svc services.RoomService) *RoomUseCase {
	return &RoomUseCase{repo: repo, svc: svc}
}
func (uc *RoomUseCase) Create(ctx context.Context, roomDTO *dto.CreateRoomRequest) (*dto.CreateRoomResponse, error) {
	const op = "usecase.room.Create"

	// Generating new uuid
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, errors.New("internal error"))
	}

	// DTO to model
	room := &models.Room{UUID: newUUID, Number: roomDTO.Number}

	// Adding room to db
	err = uc.repo.Create(ctx, room)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &dto.CreateRoomResponse{UUID: room.UUID}, nil
}

func (uc *RoomUseCase) Get(ctx context.Context) ([]*models.Room, error) {
	const op = "usecase.room.Get"

	// Getting all rooms from db
	rooms, err := uc.repo.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return rooms, nil
}

func (uc *RoomUseCase) GetByUUID(ctx context.Context, uuid uuid.UUID) (*models.Room, error) {
	const op = "usecase.room.GetByUUID"

	// Getting room from db with given uuid
	room, err := uc.repo.GetByUUID(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return room, nil
}

func (uc *RoomUseCase) GetByNumber(ctx context.Context, number string) (*models.Room, error) {
	const op = "usecase.room.GetByNumber"

	// Getting room from db with given number
	room, err := uc.repo.GetByNumber(ctx, number)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return room, nil
}

func (uc *RoomUseCase) Update(ctx context.Context, room *models.Room) error {
	const op = "usecase.room.Update"

	// Updating room in db with given room
	err := uc.repo.Update(ctx, room)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (uc *RoomUseCase) Delete(ctx context.Context, uuid uuid.UUID) error {
	const op = "usecase.room.Delete"

	// Deleting room from db with given uuid
	err := uc.repo.Delete(ctx, uuid)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
