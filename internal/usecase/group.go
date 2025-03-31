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
		return nil, fmt.Errorf("%s: %w", op, errors.New("internal error"))
	}

	// DTO to model
	group := &models.Group{UUID: newUUID, Number: groupDTO.Group}

	// Validation request group number
	if valid := uc.svc.Validate(group); valid != true {
		return nil, fmt.Errorf("%s: %w", op, errors.New("group is not valid"))
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

	groups, err := uc.repo.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return groups, nil
}

func (uc *GroupUseCase) GetByUUID(ctx context.Context, uuid uuid.UUID) (*models.Group, error) {
	const op = "usecase.group.GetByUUID"

	group, err := uc.repo.GetByUUID(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return group, nil
}

func (uc *GroupUseCase) GetByNumber(ctx context.Context, number string) (*models.Group, error) {
	const op = "usecase.group.GetByNumber"

	valid := uc.svc.Validate(&models.Group{Number: number})
	if !valid {
		return nil, fmt.Errorf("%s: %w", op, errors.New("group is not valid"))
	}

	group, err := uc.repo.GetByNumber(ctx, number)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return group, nil
}

func (uc *GroupUseCase) Update(ctx context.Context, group *models.Group) error {
	const op = "usecase.group.Update"

	valid := uc.svc.Validate(group)
	if !valid {
		return fmt.Errorf("%s: %w", op, errors.New("group is not valid"))
	}

	err := uc.repo.Update(ctx, group)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (uc *GroupUseCase) Delete(ctx context.Context, uuid uuid.UUID) error {
	const op = "usecase.group.Delete"
	err := uc.repo.Delete(ctx, uuid)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
