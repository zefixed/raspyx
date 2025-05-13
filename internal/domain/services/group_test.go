package services

import (
	"raspyx/internal/domain/models"
	"testing"
)

func TestGroupService_Validate(t *testing.T) {
	tests := []struct {
		name      string
		group     *models.Group
		wantValid bool
	}{
		{
			name:      "valid format with 3 letters",
			group:     &models.Group{Number: "12A-345 XYZ"},
			wantValid: true,
		},
		{
			name:      "valid format with letter",
			group:     &models.Group{Number: "12A-345"},
			wantValid: true,
		},
		{
			name:      "valid format with no letters",
			group:     &models.Group{Number: "123-345"},
			wantValid: true,
		},
		{
			name:      "invalid format with extra number",
			group:     &models.Group{Number: "123-34531"},
			wantValid: false,
		},
		{
			name:      "invalid format with 1 letter",
			group:     &models.Group{Number: "12A-345 X"},
			wantValid: false,
		},
		{
			name:      "invalid format with 2 letters",
			group:     &models.Group{Number: "12A-345 XY"},
			wantValid: false,
		},
		{
			name:      "invalid format with missing digit",
			group:     &models.Group{Number: "1A-345"},
			wantValid: false,
		},
		{
			name:      "valid format with Cyrillic letters",
			group:     &models.Group{Number: "12А-345 ПИШ"},
			wantValid: true,
		},
		{
			name:      "invalid format with Cyrillic letters in incorrect position",
			group:     &models.Group{Number: "1А2-345 ПИШ"},
			wantValid: false,
		},
		{
			name:      "empty string",
			group:     &models.Group{Number: ""},
			wantValid: false,
		},
	}

	groupService := NewGroupService()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid := groupService.Validate(tt.group)
			if valid != tt.wantValid {
				t.Errorf("GroupService.Validate() valid = %v, wantValid %v", valid, tt.wantValid)
			}
		})
	}
}
