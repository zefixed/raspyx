package usecase

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"raspyx/config"
	"raspyx/internal/domain/interfaces"
	"raspyx/internal/domain/models"
	"raspyx/internal/domain/services"
	"raspyx/internal/dto"
	"strconv"
	"strings"
)

type UserUseCase struct {
	repo interfaces.UserRepository
	svc  services.UserService
}

func NewUserUseCase(repo interfaces.UserRepository, svc services.UserService) *UserUseCase {
	return &UserUseCase{repo: repo, svc: svc}
}
func (uc *UserUseCase) Create(ctx context.Context, userDTO *dto.RegisterUserRequest) (*dto.RegisterUserResponse, error) {
	const op = "usecase.user.Create"

	// Generating new uuid
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, ErrGeneratingUUID)
	}

	// DTO to model
	user := &models.User{UUID: newUUID, Username: userDTO.Username, AccessLevel: 0}

	// User validation
	valid := uc.svc.Validate(user)
	if !valid {
		return nil, fmt.Errorf("%s: %w", op, ErrInvalidUser)
	}

	// Getting password hash
	passwordHash, err := uc.svc.GeneratePasswordHash(userDTO.Password)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	user.PasswordHash = passwordHash

	// Adding user to db
	err = uc.repo.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &dto.RegisterUserResponse{UUID: user.UUID}, nil
}

func (uc *UserUseCase) Login(ctx context.Context, jwt config.JWT, userDTO *dto.LoginUserRequest) (*dto.LoginUserResponse, error) {
	const op = "usecase.user.Login"

	// Getting user from db with given username
	user, err := uc.repo.GetByUsername(ctx, userDTO.Username)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, fmt.Errorf("%s: %w", op, ErrInvalidCreds)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if !uc.svc.CheckPassword(userDTO.Password, user.PasswordHash) {
		return nil, fmt.Errorf("%s: %w", op, ErrInvalidCreds)
	}

	token, err := uc.svc.CreateJWT(user.Username, user.AccessLevel, jwt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &dto.LoginUserResponse{Token: token, TokenType: "bearer"}, nil
}

func (uc *UserUseCase) Get(ctx context.Context) ([]*dto.UserDTO, error) {
	const op = "usecase.user.Get"

	// Getting all users from db
	users, err := uc.repo.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Model to DTO
	var usersDTO []*dto.UserDTO
	for _, user := range users {
		usersDTO = append(usersDTO, &dto.UserDTO{
			UUID:        user.UUID,
			Username:    user.Username,
			AccessLevel: user.AccessLevel,
		})
	}

	return usersDTO, nil
}

func (uc *UserUseCase) GetByUUID(ctx context.Context, UUID string) (*dto.UserDTO, error) {
	const op = "usecase.user.GetByUUID"

	// Parsing user uuid
	userUUID, err := uuid.Parse(UUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, ErrInvalidUUID)
	}

	// Getting user from db with given uuid
	user, err := uc.repo.GetByUUID(ctx, userUUID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Model to DTO
	userDTO := &dto.UserDTO{
		UUID:        user.UUID,
		Username:    user.Username,
		AccessLevel: user.AccessLevel,
	}

	return userDTO, nil
}

func (uc *UserUseCase) GetByUsername(ctx context.Context, username string) (*dto.UserDTO, error) {
	const op = "usecase.user.GetByUsername"

	// Getting user from db with given username
	user, err := uc.repo.GetByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Model to DTO
	userDTO := &dto.UserDTO{
		UUID:        user.UUID,
		Username:    user.Username,
		AccessLevel: user.AccessLevel,
	}

	return userDTO, nil
}

func (uc *UserUseCase) GetByAccessLevel(ctx context.Context, accessLevel string) ([]*dto.UserDTO, error) {
	const op = "usecase.user.GetByAccessLevel"

	// Converting access level to int
	al, err := strconv.Atoi(accessLevel)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Getting users from db with given AccessLevel
	users, err := uc.repo.GetByAccessLevel(ctx, al)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Model to DTO
	var usersDTO []*dto.UserDTO
	for _, user := range users {
		usersDTO = append(usersDTO, &dto.UserDTO{
			UUID:        user.UUID,
			Username:    user.Username,
			AccessLevel: user.AccessLevel,
		})
	}

	return usersDTO, nil
}

func (uc *UserUseCase) Update(ctx context.Context, UUID string, userDTO *dto.UpdateUserRequest) error {
	const op = "usecase.user.Update"

	// Parsing user uuid
	userUUID, err := uuid.Parse(UUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, ErrInvalidUUID)
	}

	// DTO to model
	user := &models.User{AccessLevel: userDTO.AccessLevel}

	// User validation
	valid := uc.svc.Validate(user)
	if !valid {
		return fmt.Errorf("%s: %w", op, ErrInvalidUser)
	}

	// Updating user in db with given user
	err = uc.repo.Update(ctx, &models.User{UUID: userUUID, AccessLevel: userDTO.AccessLevel})
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (uc *UserUseCase) Delete(ctx context.Context, UUID string) error {
	const op = "usecase.user.Delete"

	// Parsing user uuid
	userUUID, err := uuid.Parse(UUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, ErrInvalidUUID)
	}

	// Deleting user from db with given uuid
	err = uc.repo.Delete(ctx, userUUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
