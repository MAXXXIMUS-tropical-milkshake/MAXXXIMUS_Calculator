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

func (s *service) GetAllUserHistoryByID(ctx context.Context, userID int, params core.GetAllParams) (data []core.History, err error) {
	return s.historyStore.GetAllUserHistoryByID(ctx, userID, params)
}

func (s *service) SaveHistory(ctx context.Context, history core.History) (data core.History, err error) {
	return s.historyStore.SaveHistory(ctx, history)
}

func (s *service) DeleteAllHistory(ctx context.Context, userID int) (err error) {
	return s.historyStore.DeleteAllHistory(ctx, userID)
}

func (s *service) DeleteHistoryByID(ctx context.Context, historyID int) (err error) {
	return s.historyStore.DeleteHistoryByID(ctx, historyID)
}