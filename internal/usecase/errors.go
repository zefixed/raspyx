package usecase

import "errors"

var (
	ErrInvalidUUID    = errors.New("invalid uuid")
	ErrGeneratingUUID = errors.New("failed to generate uuid")
	ErrInvalidUser    = errors.New("invalid user")
	ErrInvalidCreds   = errors.New("invalid creds")
	ErrInvalidGroup   = errors.New("group is invalid")
)
