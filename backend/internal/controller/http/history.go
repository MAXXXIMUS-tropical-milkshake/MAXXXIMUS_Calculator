package http

import (
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/internal/controller/http/model"
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/internal/controller/http/model/history"
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func getUserIDFromContext(ctx *fiber.Ctx) (*int, error) {
	panic("unimplemented")
}

func (r *Router) getHistory(ctx *fiber.Ctx) error {
	var allParams history.GetAllHistoryParams

	fiberError, parseOrValidationError := parseQueryAndValidate(ctx, r.formValidator, &allParams)
	if fiberError != nil || parseOrValidationError != nil {
		return fiberError
	}

	/*id, err := getUserIDFromContext(ctx)
	if err != nil{
		logger.Log().Error(ctx.UserContext(), err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse(err.Error()))
	}*/
	id := 1
	histories, err := r.historyService.GetAllUserHistoryByID(ctx.UserContext(), id, allParams.ToCoreGetAllHistoryParams())
	if err != nil {
		logger.Log().Error(ctx.UserContext(), err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse(err.Error()))
	}
	return ctx.Status(fiber.StatusOK).JSON(model.OKResponse(histories))

}
