package usecase

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"raspyx/internal/domain/interfaces"
	"raspyx/internal/domain/models"
	"raspyx/internal/domain/services"
	"raspyx/internal/dto"
)

type GroupUseCase struct {
	repo interfaces.GroupRepository
	svc  services.GroupService
}

func NewGroupUseCase(repo interfaces.GroupRepository, svc services.GroupService) *GroupUseCase {
	return &GroupUseCase{repo: repo, svc: svc}
}

func (uc *GroupUseCase) Create(ctx context.Context, groupDTO *dto.CreateGroupRequest) (*dto.CreateGroupResponse, error) {
	const op = "usecase.group.Create"

	// Generating new uuid
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, ErrGeneratingUUID)
	}

	// DTO to model
	group := &models.Group{UUID: newUUID, Number: groupDTO.Group}

	// Validation request group number
	if valid := uc.svc.Validate(group); valid != true {
		return nil, fmt.Errorf("%s: %w", op, ErrInvalidGroup)
	}

	// Adding group to db
	err = uc.repo.Create(ctx, group)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &dto.CreateGroupResponse{UUID: group.UUID}, nil
}

func (uc *GroupUseCase) Get(ctx context.Context) ([]*models.Group, error) {
	const op = "usecase.group.Get"

	// Getting all groups from db
	groups, err := uc.repo.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return groups, nil
}

func (uc *GroupUseCase) GetByUUID(ctx context.Context, UUID string) (*models.Group, error) {
	const op = "usecase.group.GetByUUID"

	// Parsing group uuid
	groupUUID, err := uuid.Parse(UUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, ErrInvalidUUID)
	}

	// Getting group from db with given uuid
	group, err := uc.repo.GetByUUID(ctx, groupUUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return group, nil
}

func (uc *GroupUseCase) GetByNumber(ctx context.Context, number string) (*models.Group, error) {
	const op = "usecase.group.GetByNumber"

	// Validating given group number
	valid := uc.svc.Validate(&models.Group{Number: number})
	if !valid {
		return nil, fmt.Errorf("%s: %w", op, ErrInvalidGroup)
	}

	// Getting group from db with given number
	group, err := uc.repo.GetByNumber(ctx, number)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return group, nil
}

func (uc *GroupUseCase) Update(ctx context.Context, UUID string, groupDTO *dto.UpdateGroupRequest) error {
	const op = "usecase.group.Update"

	// Parsing group uuid
	groupUUID, err := uuid.Parse(UUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, ErrInvalidUUID)
	}

	group := &models.Group{UUID: groupUUID, Number: groupDTO.Group}

	// Validating group number
	valid := uc.svc.Validate(group)
	if !valid {
		return fmt.Errorf("%s: %w", op, ErrInvalidGroup)
	}

	// Updating group in db with given group
	err = uc.repo.Update(ctx, group)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (uc *GroupUseCase) Delete(ctx context.Context, UUID string) error {
	const op = "usecase.group.Delete"

	// Parsing group uuid
	groupUUID, err := uuid.Parse(UUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, ErrInvalidUUID)
	}

	// Deleting groups from db with given uuid
	err = uc.repo.Delete(ctx, groupUUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
