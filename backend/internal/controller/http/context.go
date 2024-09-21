package http

type contextKey string

const (
	jwtSecretContextKey = contextKey("jwt_secret")
	userIDContextKey    = contextKey("user_id")
)
