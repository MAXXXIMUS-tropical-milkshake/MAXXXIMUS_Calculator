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

func (s *store) GetAllUserHistoryByID(ctx context.Context, userID int, params core.GetAllHistoryParams) (data []core.History, err error) {
	var histories []core.History
	err = s.DB.WithContext(ctx).Where("user_id = ?", userID).Find(&histories).Error
	if err != nil {
		logger.Log().Error(ctx, err.Error())
		return nil, err
	}
	return histories, nil
}
