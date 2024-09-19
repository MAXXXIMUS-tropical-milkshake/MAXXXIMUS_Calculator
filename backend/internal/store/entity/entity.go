package entity

import (
	"context"
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/internal/core"
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/pkg/postgres"
)

type store struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) core.EntityStore {
	return &store{pg}
}

func (s *store) GetAll(_ context.Context, _ core.GetAllParams) (data []core.Entity, err error) {
	panic("implement me")
}

func (s *store) GetByID(_ context.Context, _ int) (data core.Entity, err error) {
	panic("implement me")
}
