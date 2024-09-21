package core

import "context"

type (
	AuthService interface {
		Signup(ctx context.Context, jwtSecret string) (token *string, err error)
	}
)
