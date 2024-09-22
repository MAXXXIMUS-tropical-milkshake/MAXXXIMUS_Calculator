package history

import (
	"time"
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/internal/controller/http/model/pagination"
)

type (
	History struct {
		ID         int       `json:"id"`
		UserID     int       `json:"user_id"`
		Expression string    `json:"expression"`
		CreatedAt  time.Time `json:"created_at"`
	}

	HistoryResponse struct {
		UserID     int    `json:"user_id"`
		Expression string `json:"expression"`
	}

	Meta struct {
		Total       int `json:"total"`
		TotalPages  int `json:"total_pages"`
		CurrentPage int `json:"current_page"`
		PerPage     int `json:"per_page"`
	}

	Response struct {
		Meta    pagination.Pagination `json:"meta"`
		History []HistoryResponse     `json:"history"`
	}

	GetAllParams struct {
		Limit  int `query:"limit" validate:"gt=0"`
		Offset int `query:"offset" validate:"gte=0"`
	}
)