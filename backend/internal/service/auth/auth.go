package auth

import (
	"context"

	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/internal/core"
)

type service struct {
	userStore core.UserStore
}

func New(userStore core.UserStore) core.AuthService {
	return &service{
		userStore: userStore,
	}
}

func (s *service) Signup(ctx context.Context, jwtSecret string) (token *string, err error) {
	userID, err := s.userStore.AddUser(ctx, core.User{})
	if err != nil {
		return nil, err
	}

	token, err = getJWTToken(jwtSecret, userID)
	if err != nil {
		return nil, err
	}

	return token, nil
}
