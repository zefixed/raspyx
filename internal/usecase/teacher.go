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

type TeacherUseCase struct {
	repo interfaces.TeacherRepository
	svc  services.TeacherService
}

func NewTeacherUseCase(repo interfaces.TeacherRepository, svc services.TeacherService) *TeacherUseCase {
	return &TeacherUseCase{repo: repo, svc: svc}
}
func (uc *TeacherUseCase) Create(ctx context.Context, teacherDTO *dto.CreateTeacherRequest) (*dto.CreateTeacherResponse, error) {
	const op = "usecase.teacher.Create"

	// Generating new uuid
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, ErrGeneratingUUID)
	}

	// DTO to model
	teacher := &models.Teacher{
		UUID:       newUUID,
		FirstName:  teacherDTO.FirstName,
		SecondName: teacherDTO.SecondName,
		MiddleName: teacherDTO.MiddleName,
	}

	// Adding teacher to db
	err = uc.repo.Create(ctx, teacher)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &dto.CreateTeacherResponse{UUID: teacher.UUID}, nil
}

func (uc *TeacherUseCase) Get(ctx context.Context) ([]*dto.TeacherDTO, error) {
	const op = "usecase.teacher.Get"

	// Getting all teachers from db
	teachers, err := uc.repo.Get(ctx)
	var teachersDTO []*dto.TeacherDTO
	for _, teacher := range teachers {
		teachersDTO = append(teachersDTO, &dto.TeacherDTO{
			Name: fmt.Sprintf("%v %v %v",
				teacher.SecondName,
				teacher.FirstName,
				teacher.MiddleName,
			)})
	}
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return teachersDTO, nil
}

func (uc *TeacherUseCase) GetByUUID(ctx context.Context, UUID string) (*models.Teacher, error) {
	const op = "usecase.teacher.GetByUUID"

	// Parsing teacher uuid
	teacherUUID, err := uuid.Parse(UUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, ErrInvalidUUID)
	}

	// Getting teacher from db with given uuid
	teacher, err := uc.repo.GetByUUID(ctx, teacherUUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return teacher, nil
}

func (uc *TeacherUseCase) GetByFullName(ctx context.Context, fullname string) ([]*models.Teacher, error) {
	const op = "usecase.teacher.GetByFullName"

	// Getting teacher from db with given fullname
	teachers, err := uc.repo.GetByFullName(ctx, fullname)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return teachers, nil
}

func (uc *TeacherUseCase) Update(ctx context.Context, UUID string, teacherDTO *dto.UpdateTeacherRequest) error {
	const op = "usecase.teacher.Update"

	// Parsing teacher uuid
	teacherUUID, err := uuid.Parse(UUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, ErrInvalidUUID)
	}

	// Updating teacher in db with given teacher
	err = uc.repo.Update(ctx, &models.Teacher{
		UUID:       teacherUUID,
		FirstName:  teacherDTO.FirstName,
		SecondName: teacherDTO.SecondName,
		MiddleName: teacherDTO.MiddleName,
	})
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (uc *TeacherUseCase) Delete(ctx context.Context, UUID string) error {
	const op = "usecase.teacher.Delete"

	// Parsing teacher uuid
	teacherUUID, err := uuid.Parse(UUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, ErrInvalidUUID)
	}

	// Deleting teacher from db with given uuid
	err = uc.repo.Delete(ctx, teacherUUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
