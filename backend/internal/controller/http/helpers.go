package http

import (
	"fmt"
	"errors"
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/internal/controller/http/model"
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/internal/controller/http/model/validator"
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/internal/controller/http/model/pagination"
)

func parseQueryAndValidate(ctx *fiber.Ctx, formValidator validator.FormValidatorService, data interface{}) (fiberError, parseOrValidationError error) {
	if err := ctx.QueryParser(data); err != nil {
		logger.Log().Error(ctx.UserContext(), err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse(err.Error())), err
	}

	return validate(ctx, formValidator, data)
}

func parseBodyAndValidate(ctx *fiber.Ctx, formValidator validator.FormValidatorService, data interface{}) (fiberError, parseOrValidationError error) {
	if err := ctx.BodyParser(data); err != nil {
		if errors.Is(err, fiber.ErrUnprocessableEntity) {
			logger.Log().Debug(ctx.UserContext(), err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse(model.ErrInvalidBody.Error())), err
		}

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

func getJWTSecretFromContext(ctx *fiber.Ctx) (*string, error) {
	secret, ok := ctx.UserContext().Value(jwtSecretContextKey).(string)
	if !ok {
		return nil, model.ErrInvalidJWTSecret
	}

	return &secret, nil
}

func getUserIDFromContext(ctx *fiber.Ctx) (*int, error) {
	userID, ok := ctx.UserContext().Value(userIDContextKey).(int)
	if !ok {
		return nil, model.ErrInvalidUserID
	}

	return &userID, nil
}

func paginate(total, limit, offset int) pagination.Pagination {
	var (
		currentPage = (offset / limit) + 1
		perPage     = limit
		totalPages  = (total + perPage - 1) / perPage
	)

	return pagination.Pagination{
		Total:       total,
		TotalPages:  totalPages,
		CurrentPage: currentPage,
		PerPage:     perPage,
	}
}
