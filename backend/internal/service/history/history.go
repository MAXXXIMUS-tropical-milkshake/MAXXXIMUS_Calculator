package history

import (
	"context"

	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/internal/core"
)

type service struct {
	historyStore core.HistoryStore
}

func New(store core.HistoryStore) core.HistoryService {
	return &service{
		historyStore: store,
	}
}

func (s *service) GetAllUserHistoryByID(ctx context.Context, userID int, params core.GetAllHistoryParams) (data []core.History, err error) {
	return s.GetAllUserHistoryByID(ctx, userID, params)
}
