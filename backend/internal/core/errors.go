package core

import (
	"gorm.io/gorm"
	"errors"
)

var (
	ErrRecordNotFound  = gorm.ErrRecordNotFound

	ErrHistoryNotFound = errors.New("history not found")
)