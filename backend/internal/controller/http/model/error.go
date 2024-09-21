package model

import (
	"errors"
)

var (
	ErrValidationFailed = errors.New("validation failed")

	ErrInvalidJWTSecret = errors.New("invalid jwt secret")
	ErrInvalidJWTToken  = errors.New("invalid jwt token")
	ErrInvalidUserID    = errors.New("invalid user id")
)
