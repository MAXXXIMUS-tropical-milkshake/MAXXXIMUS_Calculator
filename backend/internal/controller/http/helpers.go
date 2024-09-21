package http

import (
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/internal/controller/http/model/validator"
	"fmt"
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/internal/controller/http/model"
)

func parseQueryAndValidate(ctx *fiber.Ctx, formValidator validator.FormValidatorService, data interface{}) (fiberError, parseOrValidationError error) {
	if err := ctx.QueryParser(data); err != nil {
		logger.Log().Error(ctx.UserContext(), err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse(err.Error())), err
	}

	return validate(ctx, formValidator, data)
}

func validate(ctx *fiber.Ctx, formValidator validator.FormValidatorService, data interface{}) (fiberError, parseOrValidationError error) {
	errs := formValidator.Validate(data)
	if len(errs) > 0 {
		logger.Log().Info(ctx.UserContext(), fmt.Sprintf("%v", errs))
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(model.ErrorResponse(fiber.Map{
			"validation_errors": errs,
		})), model.ErrValidationFailed
	}
	return nil, nil
}