package http

import (
	"errors"

	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/internal/controller/http/model"
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/internal/controller/http/model/history"
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/internal/core"
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func (r *Router) getAllHistory(ctx *fiber.Ctx) error {
	var getAllParams history.GetAllParams

	fiberError, parseOrValidationError := parseQueryAndValidate(ctx, r.formValidator, &getAllParams)
	if fiberError != nil || parseOrValidationError != nil {
		return fiberError
	}

	userID, err := getUserIDFromContext(ctx)
	if err != nil {
		logger.Log().Error(ctx.UserContext(), err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse(err.Error()))
	}

	histories, err := r.historyService.GetAllUserHistoryByID(ctx.UserContext(), *userID, getAllParams.ToCoreGetAllParams())
	if err != nil {
		logger.Log().Error(ctx.UserContext(), err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse(err.Error()))
	}

	pagination := paginate(len(histories), getAllParams.Limit, getAllParams.Offset)
	response := history.ToResponse(pagination, histories)

	return ctx.Status(fiber.StatusOK).JSON(model.OKResponse(response))
}

func (r *Router) saveHistory(ctx *fiber.Ctx) error {
	var historyRequest history.History

	fiberError, parseOrValidationError := parseBodyAndValidate(ctx, r.formValidator, &historyRequest)
	if fiberError != nil || parseOrValidationError != nil {
		return fiberError
	}

	userID, err := getUserIDFromContext(ctx)
	if err != nil {
		logger.Log().Error(ctx.UserContext(), err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse(err.Error()))
	}

	savedHistory, err := r.historyService.SaveHistory(ctx.UserContext(), historyRequest.ToCoreHistory(*userID))
	if err != nil {
		logger.Log().Error(ctx.UserContext(), err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse(err.Error()))
	}

	historyResponse := history.ToHistoryResponse(savedHistory)

	return ctx.Status(fiber.StatusOK).JSON(model.OKResponse(historyResponse))
}

func (r *Router) deleteHistory(ctx *fiber.Ctx) error {
	userID, err := getUserIDFromContext(ctx)
	if err != nil {
		logger.Log().Error(ctx.UserContext(), err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse(err.Error()))
	}

	err = r.historyService.DeleteAllHistory(ctx.UserContext(), *userID)
	if err != nil {
		logger.Log().Error(ctx.UserContext(), err.Error())
		if errors.Is(err, core.ErrHistoryNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(model.ErrorResponse(err.Error()))
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse(err.Error()))
	}

	return ctx.SendStatus(fiber.StatusNoContent)

}

func (r *Router) deleteHistoryByID(ctx *fiber.Ctx) error {
	userID, err := getUserIDFromContext(ctx)
	if err != nil {
		logger.Log().Error(ctx.UserContext(), err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse(err.Error()))
	}

	historyID, err := ctx.ParamsInt("record_id")
	if err != nil {
		logger.Log().Error(ctx.UserContext(), err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse(err.Error()))
	}

	err = r.historyService.DeleteHistoryByID(ctx.UserContext(), historyID, *userID)
	if err != nil {
		logger.Log().Error(ctx.UserContext(), err.Error())
		if errors.Is(err, core.ErrHistoryNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(model.ErrorResponse(err.Error()))
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse(err.Error()))
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
