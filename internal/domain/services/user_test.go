package services

import (
	"errors"
	"raspyx/internal/domain/models"
	"testing"
)

func TestUserService_GeneratePasswordHash(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  error
	}{
		{
			"valid password is less 72 bytes long",
			"qwerty",
			nil,
		},
		{
			"valid password 71 bytes long",
			"12345678901234567890123456789012345678901234567890123456789012345678901",
			nil,
		},
		{
			"invalid password 72 bytes long",
			"123456789012345678901234567890123456789012345678901234567890123456789012",
			nil,
		},
		{
			"invalid password 73 bytes long",
			"1234567890123456789012345678901234567890123456789012345678901234567890123",
			InvalidPassword,
		},
		{
			"invalid password 74 bytes long",
			"12345678901234567890123456789012345678901234567890123456789012345678901234",
			InvalidPassword,
		},
	}

	userService := NewUserService()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := userService.GeneratePasswordHash(tt.password)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("UserService.GeneratePasswordHash() err = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserService_CheckPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		hash     string
		want     bool
	}{
		{
			"valid password",
			"qwerty",
			"$2a$10$OAenL4SlOcLa1sBQDcoCeOKwJerm04KEQvgh4EWfVfk8QAe0xN8cS",
			true,
		},
		{
			"invalid password",
			"qwerty1",
			"$2a$10$OAenL4SlOcLa1sBQDcoCeOKwJerm04KEQvgh4EWfVfk8QAe0xN8cS",
			false,
		},
		{
			"invalid password",
			"qwerty",
			"$2a$10$OAenL4SlOcLa1sBQDcoCeOKwJerm04KEQvgh4EWfVfk8QAe0xN8cs",
			false,
		},
		{
			"empty hash",
			"qwerty",
			"",
			false,
		},
		{
			"empty password",
			"",
			"$2a$10$OAenL4SlOcLa1sBQDcoCeOKwJerm04KEQvgh4EWfVfk8QAe0xN8cS",
			false,
		},
	}

	userService := NewUserService()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := userService.CheckPassword(tt.password, tt.hash)
			if got != tt.want {
				t.Errorf("UserService.CheckPassword() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_Validate(t *testing.T) {
	tests := []struct {
		name string
		user *models.User
		want bool
	}{
		{
			"valid user role admin",
			&models.User{Role: "admin"},
			true,
		},
		{
			"valid user role moderator",
			&models.User{Role: "moderator"},
			true,
		},
		{
			"valid user role user",
			&models.User{Role: "user"},
			true,
		},
		{
			"invalid user role",
			&models.User{Role: ""},
			false,
		},
	}

	userService := NewUserService()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := userService.Validate(tt.user)
			if got != tt.want {
				t.Errorf("UserService.Validate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
