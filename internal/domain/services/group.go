package services

import (
	"raspyx/internal/domain/models"
	"regexp"
)

type GroupService struct{}

func NewGroupService() *GroupService {
	return &GroupService{}
}

func (s *GroupService) Validate(group *models.Group) bool {
	re := `^\d{2}[0-9a-zA-Zа-яА-Я]-\d{3,4}(\s[a-zA-Zа-яА-Я]{3})?$`
	match, err := regexp.MatchString(re, group.Number)
	if err != nil {
		return false
	}

	return match
}
