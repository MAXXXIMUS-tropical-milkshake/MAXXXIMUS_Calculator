package http

import (
	"context"
	"fmt"
	"strings"

	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/internal/controller/http/model"
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func InjectJWTSecretMiddleware(secret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.UserContext()
		ctx = context.WithValue(ctx, jwtSecretContextKey, secret)

		c.SetUserContext(ctx)

		return c.Next()
	}
}

func (r *Router) protectedMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		secret, err := getJWTSecretFromContext(c)
		if err != nil {
			logger.Log().Error(c.UserContext(), err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse(err))
		}

		tokenString := c.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer")
		tokenString = strings.TrimSpace(tokenString)

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}

			return []byte(*secret), nil
		})
		if err != nil {
			logger.Log().Error(c.UserContext(), err.Error())
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			id, ok := claims["user_id"].(float64)
			if !ok {
				logger.Log().Error(c.UserContext(), model.ErrInvalidJWTToken.Error())
				return c.SendStatus(fiber.StatusUnauthorized)
			}

			ctx := c.UserContext()
			ctx = context.WithValue(ctx, userIDContextKey, int(id))
			c.SetUserContext(ctx)
		} else {
			logger.Log().Error(c.UserContext(), model.ErrInvalidJWTToken.Error())
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		return c.Next()
	}
}
