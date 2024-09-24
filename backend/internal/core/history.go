package core

import (
	"time"
	"context"
)

type (
	History struct {
		ID        int       `gorm:"column:id"`
		UserID    int       `gorm:"column:user_id"`
		Expression string   `gorm:"column:expression"`
		CreatedAt time.Time `jgorm:"column:created_at"`
	}

	HistoryService interface {
		GetAllUserHistoryByID(ctx context.Context, userID int, params GetAllParams) ([]History, error)
		SaveHistory(ctx context.Context, history History) (History, error) // не факт, что передаются такие параметры
		DeleteAllHistory(ctx context.Context, userID int) error
		DeleteHistoryByID(ctx context.Context, historyID int, userID int) error
	}

	HistoryStore interface {
		GetAllUserHistoryByID(ctx context.Context, userID int, params GetAllParams) ([]History, error)
		SaveHistory(ctx context.Context, history History) (History, error) // не факт, что передаются такие параметры
		DeleteAllHistory(ctx context.Context, userID int) error
		DeleteHistoryByID(ctx context.Context, historyID int, userID int) error
	}

	GetAllParams struct {
		Limit      *int    
		Offset     *int 
	}
)