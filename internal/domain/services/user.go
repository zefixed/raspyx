package services

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"raspyx/config"
	"raspyx/internal/domain/models"
	"time"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

var (
	InvalidPassword = errors.New("password is longer than 72 bytes")
)

func (s *UserService) GeneratePasswordHash(password string) (string, error) {
	if len([]byte(password)) > 72 {
		return "", InvalidPassword
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (s *UserService) CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *UserService) Validate(user *models.User) bool {
	return user.AccessLevel < 100
}

func (s *UserService) CreateJWT(username string, accessLevel int, JWT config.JWT) (string, error) {
	claims := jwt.MapClaims{
		"sub":          username,
		"access_level": accessLevel,
		"exp":          time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JWT.JWTSecret))
}
