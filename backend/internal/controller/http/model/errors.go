package model

import (
	"errors"
)

var (
	// validation
	ErrValidationFailed = errors.New("validation failed")
	ErrInvalidBody      = errors.New("invalid body")

	// auth
	ErrInvalidJWTSecret = errors.New("invalid jwt secret")
	ErrInvalidJWTToken  = errors.New("invalid jwt token")
	ErrInvalidUserID    = errors.New("invalid user id")

	// calculate
	ErrDivisionByZero = errors.New("division by zero")
)
