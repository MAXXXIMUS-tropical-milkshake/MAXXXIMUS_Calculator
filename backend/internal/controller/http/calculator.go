package http

import (
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/internal/controller/http/model"
	calculatorModel "github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/internal/controller/http/model/calculator"
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/pkg/logger"
	"github.com/Knetic/govaluate"
	"github.com/gofiber/fiber/v2"
)

func (r *Router) calculate(ctx *fiber.Ctx) error {
	var calculator calculatorModel.Calculator 

	fiberError, parseOrValidationError := parseQueryAndValidate(ctx, r.formValidator, &calculator)
	if fiberError != nil || parseOrValidationError != nil {
		return fiberError
	}

	expression, err := govaluate.NewEvaluableExpression(calculator.Expression)

	if err != nil {
		logger.Log().Error(ctx.UserContext(), err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse(err.Error()))
	}
	
	result, err := expression.Evaluate(nil)

	if err != nil {
		logger.Log().Error(ctx.UserContext(), err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(model.OKResponse(result))
}