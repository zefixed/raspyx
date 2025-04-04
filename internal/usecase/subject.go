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

type SubjectUseCase struct {
	repo interfaces.SubjectRepository
	svc  services.SubjectService
}

func NewSubjectUseCase(repo interfaces.SubjectRepository, svc services.SubjectService) *SubjectUseCase {
	return &SubjectUseCase{repo: repo, svc: svc}
}
func (uc *SubjectUseCase) Create(ctx context.Context, SubjectDTO *dto.CreateSubjectRequest) (*dto.CreateSubjectResponse, error) {
	const op = "usecase.subject.Create"

	// Generating new uuid
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, errors.New("internal error"))
	}

	// DTO to model
	subject := &models.Subject{UUID: newUUID, Name: SubjectDTO.Name}

	// Adding subject to db
	err = uc.repo.Create(ctx, subject)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &dto.CreateSubjectResponse{UUID: subject.UUID}, nil
}

func (uc *SubjectUseCase) Get(ctx context.Context) ([]*models.Subject, error) {
	const op = "usecase.subject.Get"

	// Getting all subjects from db
	subjects, err := uc.repo.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return subjects, nil
}

func (uc *SubjectUseCase) GetByUUID(ctx context.Context, UUID string) (*models.Subject, error) {
	const op = "usecase.subject.GetByUUID"

	// Parsing subject uuid
	subjectUUID, err := uuid.Parse(UUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, ErrInvalidUUID)
	}

	// Getting subject from db with given uuid
	subject, err := uc.repo.GetByUUID(ctx, subjectUUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return subject, nil
}

func (uc *SubjectUseCase) GetByName(ctx context.Context, name string) ([]*models.Subject, error) {
	const op = "usecase.subject.GetByName"

	// Getting subject from db with given name
	subjects, err := uc.repo.GetByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return subjects, nil
}

func (uc *SubjectUseCase) Update(ctx context.Context, UUID string, subjectDTO *dto.UpdateSubjectRequest) error {
	const op = "usecase.subject.Update"

	// Parsing subject uuid
	subjectUUID, err := uuid.Parse(UUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, ErrInvalidUUID)
	}

	// Updating subject in db with given subject
	err = uc.repo.Update(ctx, &models.Subject{UUID: subjectUUID, Name: subjectDTO.Name})
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (uc *SubjectUseCase) Delete(ctx context.Context, UUID string) error {
	const op = "usecase.subject.Delete"

	// Parsing subject uuid
	subjectUUID, err := uuid.Parse(UUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, ErrInvalidUUID)
	}
	
	// Deleting subject from db with given uuid
	err = uc.repo.Delete(ctx, subjectUUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
