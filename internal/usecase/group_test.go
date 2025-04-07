package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"raspyx/internal/domain/interfaces/mocks"
	"raspyx/internal/domain/models"
	"raspyx/internal/domain/services"
	"raspyx/internal/dto"
	"testing"
)

func TestGroupUseCase_Create(t *testing.T) {
	tests := []struct {
		name           string
		groupDTO       *dto.CreateGroupRequest
		mockRepoError  error
		expectedError  error
		expectRepoCall bool
	}{
		{
			name: "Successful creation",
			groupDTO: &dto.CreateGroupRequest{
				Group: "221-352",
			},
			mockRepoError:  nil,
			expectedError:  nil,
			expectRepoCall: true,
		},
		{
			name: "Successful creation",
			groupDTO: &dto.CreateGroupRequest{
				Group: "22A-352",
			},
			mockRepoError:  nil,
			expectedError:  nil,
			expectRepoCall: true,
		},
		{
			name: "Invalid Group",
			groupDTO: &dto.CreateGroupRequest{
				Group: "1234-1234",
			},
			mockRepoError:  assert.AnError,
			expectedError:  ErrInvalidGroup,
			expectRepoCall: false,
		},
		{
			name: "Empty group",
			groupDTO: &dto.CreateGroupRequest{
				Group: "",
			},
			mockRepoError:  assert.AnError,
			expectedError:  ErrInvalidGroup,
			expectRepoCall: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mocks for repo and service
			mockRepo := new(mocks.GroupRepository)
			mockService := new(services.GroupService)

			if tt.expectRepoCall {
				mockRepo.On("Create", mock.Anything, mock.Anything).Return(tt.mockRepoError)
			}

			// Creating UseCase
			uc := NewGroupUseCase(mockRepo, *mockService)

			// Execute testing function
			_, err := uc.Create(context.Background(), tt.groupDTO)

			// Error checking
			if tt.expectedError != nil {
				assert.True(t, errors.Is(err, tt.expectedError))
			} else {
				assert.NoError(t, err)
			}

			// Checking calling mocks
			if tt.expectRepoCall {
				mockRepo.AssertExpectations(t)
			} else {
				mockRepo.AssertNotCalled(t, "Create")
			}
		})
	}
}

func TestGroupUseCase_Get(t *testing.T) {
	tests := []struct {
		name           string
		mockReturn     []*models.Group
		mockError      error
		expectedResult []*models.Group
		expectedError  error
	}{
		{
			name: "Successful fetch",
			mockReturn: []*models.Group{
				{UUID: uuid.New(), Number: "221-352"},
				{UUID: uuid.New(), Number: "321-123"},
			},
			mockError: nil,
			expectedResult: []*models.Group{
				{UUID: uuid.Nil, Number: "221-352"},
				{UUID: uuid.Nil, Number: "321-123"},
			},
			expectedError: nil,
		},
		{
			name:           "Error fetching groups",
			mockReturn:     nil,
			mockError:      assert.AnError,
			expectedResult: nil,
			expectedError:  assert.AnError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Мок
			mockRepo := new(mocks.GroupRepository)
			mockService := new(services.GroupService)

			mockRepo.On("Get", mock.Anything).Return(tt.mockReturn, tt.mockError)

			uc := NewGroupUseCase(mockRepo, *mockService)

			result, err := uc.Get(context.Background())

			if tt.expectedError != nil {
				assert.ErrorIs(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)

				assert.Equal(t, len(tt.expectedResult), len(result))
				for i := range tt.expectedResult {
					assert.Equal(t, tt.mockReturn[i].Number, result[i].Number)
				}
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestGroupUseCase_GetByUUID(t *testing.T) {
	tests := []struct {
		name          string
		inputUUID     string
		mockReturn    *models.Group
		mockError     error
		expectedError error
	}{
		{
			name:          "Invalid UUID format",
			inputUUID:     "invalid-uuid",
			expectedError: ErrInvalidUUID,
		},
		{
			name:          "Repository returns error",
			inputUUID:     uuid.New().String(),
			mockReturn:    nil,
			mockError:     assert.AnError,
			expectedError: assert.AnError,
		},
		{
			name:      "Successful fetch",
			inputUUID: uuid.New().String(),
			mockReturn: &models.Group{
				UUID:   uuid.New(),
				Number: "123-456",
			},
			mockError:     nil,
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(mocks.GroupRepository)
			mockService := new(services.GroupService)

			uc := NewGroupUseCase(mockRepo, *mockService)

			parsedUUID, err := uuid.Parse(tt.inputUUID)
			if err == nil {
				mockRepo.On("GetByUUID", mock.Anything, parsedUUID).Return(tt.mockReturn, tt.mockError)
			}

			group, err := uc.GetByUUID(context.Background(), tt.inputUUID)

			if tt.expectedError != nil {
				assert.ErrorIs(t, err, tt.expectedError)
				assert.Nil(t, group)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.mockReturn, group)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
