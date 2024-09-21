package core

import (
	"context"
	"time"
)

type (
	History struct {
		ID         int       `gorm:"column:id"`
		UserID     int       `gorm:"column:user_id"`
		Expression string    `gorm:"column:expression"`
		CreatedAt  time.Time `gorm:"column:created_at"`
	}

	HistoryService interface {
		GetAllUserHistoryByID(ctx context.Context, userID int, params GetAllHistoryParams) (data []History, err error)
	}

	HistoryStore interface {
		GetAllUserHistoryByID(ctx context.Context, userID int, params GetAllHistoryParams) (data []History, err error)
	}

	GetAllHistoryParams struct {
		Limit  int
		Offset int
	}
)

func (History) TableName() string {
	return "history"
}
