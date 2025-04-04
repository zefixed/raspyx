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

type LocationUseCase struct {
	repo interfaces.LocationRepository
	svc  services.LocationService
}

func NewLocationUseCase(repo interfaces.LocationRepository, svc services.LocationService) *LocationUseCase {
	return &LocationUseCase{repo: repo, svc: svc}
}
func (uc *LocationUseCase) Create(ctx context.Context, locationDTO *dto.CreateLocationRequest) (*dto.CreateLocationResponse, error) {
	const op = "usecase.location.Create"

	// Generating new uuid
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, errors.New("internal error"))
	}

	// DTO to model
	location := &models.Location{UUID: newUUID, Name: locationDTO.Name}

	// Adding location to db
	err = uc.repo.Create(ctx, location)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &dto.CreateLocationResponse{UUID: location.UUID}, nil
}

func (uc *LocationUseCase) Get(ctx context.Context) ([]*models.Location, error) {
	const op = "usecase.location.Get"

	// Getting all locations from db
	locations, err := uc.repo.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return locations, nil
}

func (uc *LocationUseCase) GetByUUID(ctx context.Context, UUID string) (*models.Location, error) {
	const op = "usecase.location.GetByUUID"

	// Parsing location uuid
	locationUUID, err := uuid.Parse(UUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, ErrInvalidUUID)
	}

	// Getting location from db with given uuid
	location, err := uc.repo.GetByUUID(ctx, locationUUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return location, nil
}

func (uc *LocationUseCase) GetByName(ctx context.Context, name string) (*models.Location, error) {
	const op = "usecase.location.GetByName"

	// Getting location from db with given name
	location, err := uc.repo.GetByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return location, nil
}

func (uc *LocationUseCase) Update(ctx context.Context, UUID string, locationDTO *dto.UpdateLocationRequest) error {
	const op = "usecase.location.Update"

	// Parsing location uuid
	locationUUID, err := uuid.Parse(UUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, ErrInvalidUUID)
	}

	// Updating location in db with given location
	err = uc.repo.Update(ctx, &models.Location{UUID: locationUUID, Name: locationDTO.Name})
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (uc *LocationUseCase) Delete(ctx context.Context, UUID string) error {
	const op = "usecase.location.Delete"

	// Parsing location uuid
	locationUUID, err := uuid.Parse(UUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, ErrInvalidUUID)
	}

	// Deleting location from db with given uuid
	err = uc.repo.Delete(ctx, locationUUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
