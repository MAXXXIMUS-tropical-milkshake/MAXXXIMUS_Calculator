package http

import (
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/internal/controller/http/model/history"
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/internal/controller/http/model"
	"github.com/gofiber/fiber/v2"
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/pkg/logger"
)

func (r *Router) getAllHistory(ctx *fiber.Ctx) error {
	var getAllParams history.GetAllParams

	fiberError, parseOrValidationError := parseQueryAndValidate(ctx, r.formValidator, &getAllParams)
	if fiberError != nil || parseOrValidationError != nil {
		return fiberError
	}

	userID, err := getUserIDFromContext(ctx)
	if err != nil{
		logger.Log().Error(ctx.UserContext(), err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse(err.Error()))
	}

	histories, err := r.historyService.GetAllUserHistoryByID(ctx.UserContext(), *userID, getAllParams.ToCoreGetAllParams())
	if err != nil {
		logger.Log().Error(ctx.UserContext(), err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse(err.Error()))
	}

	total := len(histories)

	pagination := paginate(total, getAllParams.Limit, getAllParams.Offset)

	response := history.ToResponse(pagination, histories)

	return ctx.Status(fiber.StatusOK).JSON(model.OKResponse(response))
}

func (r *Router) saveHistory(ctx *fiber.Ctx) error {
	panic("implement me")
}

func (r *Router) deleteHistory(ctx *fiber.Ctx) error {
	panic("implement me")
}

func (r *Router) deleteHistoryByID(ctx *fiber.Ctx) error {
	panic("implement me")
}