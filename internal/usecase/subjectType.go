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

type SubjectTypeUseCase struct {
	repo interfaces.SubjectTypeRepository
	svc  services.SubjectTypeService
}

func NewSubjectTypeUseCase(repo interfaces.SubjectTypeRepository, svc services.SubjectTypeService) *SubjectTypeUseCase {
	return &SubjectTypeUseCase{repo: repo, svc: svc}
}
func (uc *SubjectTypeUseCase) Create(ctx context.Context, subjectTypeDTO *dto.CreateSubjectTypeRequest) (*dto.CreateSubjectTypeResponse, error) {
	const op = "usecase.subjectType.Create"

	// Generating new uuid
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, ErrGeneratingUUID)
	}

	// DTO to model
	subjectType := &models.SubjectType{UUID: newUUID, Type: subjectTypeDTO.Type}

	// Adding subjectType to db
	err = uc.repo.Create(ctx, subjectType)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &dto.CreateSubjectTypeResponse{UUID: subjectType.UUID}, nil
}

func (uc *SubjectTypeUseCase) Get(ctx context.Context) ([]*models.SubjectType, error) {
	const op = "usecase.subjectType.Get"

	// Getting all subjectTypes from db
	subjectTypes, err := uc.repo.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return subjectTypes, nil
}

func (uc *SubjectTypeUseCase) GetByUUID(ctx context.Context, UUID string) (*models.SubjectType, error) {
	const op = "usecase.subjectType.GetByUUID"

	// Parsing subject type uuid
	subjectTypeUUID, err := uuid.Parse(UUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, ErrInvalidUUID)
	}

	// Getting subjectType from db with given uuid
	subjectType, err := uc.repo.GetByUUID(ctx, subjectTypeUUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return subjectType, nil
}

func (uc *SubjectTypeUseCase) GetByType(ctx context.Context, number string) (*models.SubjectType, error) {
	const op = "usecase.subjectType.GetByType"

	// Getting subjectType from db with given number
	subjectType, err := uc.repo.GetByType(ctx, number)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return subjectType, nil
}

func (uc *SubjectTypeUseCase) Update(ctx context.Context, UUID string, subjectTypeDTO *dto.UpdateSubjectTypeRequest) error {
	const op = "usecase.subjectType.Update"

	// Parsing subject type uuid
	subjectTypeUUID, err := uuid.Parse(UUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, ErrInvalidUUID)
	}

	// Updating subjectType in db with given subjectType
	err = uc.repo.Update(ctx, &models.SubjectType{UUID: subjectTypeUUID, Type: subjectTypeDTO.Type})
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (uc *SubjectTypeUseCase) Delete(ctx context.Context, UUID string) error {
	const op = "usecase.subjectType.Delete"

	// Parsing subject type uuid
	subjectTypeUUID, err := uuid.Parse(UUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, ErrInvalidUUID)
	}

	// Deleting subjectType from db with given uuid
	err = uc.repo.Delete(ctx, subjectTypeUUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
