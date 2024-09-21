package http

import (
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/internal/controller/http/model"
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func (r *Router) signup(ctx *fiber.Ctx) error {
	jwtSecret, err := getJWTSecretFromContext(ctx)
	if err != nil {
		logger.Log().Error(ctx.UserContext(), err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse(err))
	}

	token, err := r.authService.Signup(ctx.UserContext(), *jwtSecret)
	if err != nil {
		logger.Log().Error(ctx.UserContext(), err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse(err))
	}

	return ctx.Status(fiber.StatusOK).JSON(model.OKResponse(token))
}
