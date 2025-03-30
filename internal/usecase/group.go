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

func (uc *GroupUseCase) Delete(ctx context.Context, uuid uuid.UUID) error {
	const op = "usecase.group.Delete"
	err := uc.repo.Delete(ctx, uuid)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
