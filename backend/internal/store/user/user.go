package user

import (
	"context"
	"time"

	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/internal/core"
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/pkg/postgres"
)

type store struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) core.UserStore {
	return &store{pg}
}

func (s *store) AddUser(ctx context.Context, user core.User) (id int, err error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := s.DB.WithContext(ctx)

	if err := query.Create(&user).Error; err != nil {
		return 0, err
	}

	return user.ID, nil
}
