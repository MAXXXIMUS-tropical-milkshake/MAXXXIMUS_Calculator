package history

import (
	"context"

	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/internal/core"
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/pkg/logger"
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/pkg/postgres"
)

type store struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) core.HistoryStore {
	return &store{pg}
}

func (s *store) GetAllUserHistoryByID(ctx context.Context, userID int, params core.GetAllParams) ([]core.History, error) {
	var histories []core.History

	query := s.DB.WithContext(ctx).Where("user_id = ?", userID).Select("id", "user_id", "expression", "created_at")

	if params.Limit != nil {
		query = query.Limit(*params.Limit)
	}
	if params.Offset != nil {
		query = query.Offset(*params.Offset)
	}

	if err := query.Find(&histories).Error; err != nil {
		logger.Log().Error(ctx, err.Error())
		return nil, err
	}

	return histories, nil
}

func (s *store) SaveHistory(ctx context.Context, history core.History) (data core.History, err error) {
	if err := s.DB.WithContext(ctx).Create(&history).Error; err != nil {
		logger.Log().Error(ctx, err.Error())
		return core.History{}, err
	}
	return history, nil
}

func (s *store) DeleteAllHistory(ctx context.Context, userID int) (err error) {
	if err := s.DB.WithContext(ctx).Where("user_id = ?", userID).Delete(&core.History{}).Error; err != nil {
		logger.Log().Error(ctx, err.Error())
		return err
	}
	return nil
}

func (s *store) DeleteHistoryByID(ctx context.Context, historyID, userID int) (err error) {
	if err := s.DB.WithContext(ctx).Where("id = ? AND user_id = ?", historyID, userID).Delete(&core.History{}).Error; err != nil {
		logger.Log().Error(ctx, err.Error())
		return err
	}

	return nil
}
