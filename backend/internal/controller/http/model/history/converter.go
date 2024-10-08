package history

import (
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/internal/core"
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/internal/controller/http/model/pagination"
)

func (p *GetAllParams) ToCoreGetAllParams() core.GetAllParams {
	return core.GetAllParams{
		Limit:  &p.Limit,
		Offset: &p.Offset,
	}
}

func ToResponse(meta pagination.Pagination, histories []core.History) Response {
	res := make([]HistoryResponse, len(histories))

	for i, history := range histories {
		res[i] = ToHistoryResponse(history)
	}

	return Response{
		Meta: 	 meta,
		History: res,
	}
}

func ToHistoryResponse(history core.History) HistoryResponse {
	return HistoryResponse{
		UserID:     history.UserID,
		Expression: history.Expression,
	}
}

func (h *History) ToCoreHistory(userID int) core.History {
	return core.History{
		UserID:     userID,
		Expression: h.Expression,
	}
}
