package http

import (
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/internal/controller/http/model"
	calculatorModel "github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/internal/controller/http/model/calculator"
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/pkg/logger"
	"github.com/Pramod-Devireddy/go-exprtk"
	"github.com/gofiber/fiber/v2"
)

func (r *Router) calculate(ctx *fiber.Ctx) error {
	var calculator calculatorModel.Calculator 
	fiberError, parseOrValidationError := parseQueryAndValidate(ctx, r.formValidator, &calculator)
	if fiberError != nil || parseOrValidationError != nil {
		return fiberError
	}

	exprtkObj := exprtk.NewExprtk()
	defer exprtkObj.Delete()

	exprtkObj.SetExpression(calculator.Expression);
	
	err := exprtkObj.CompileExpression()
	if err != nil {
		logger.Log().Error(ctx.UserContext(), err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse(err.Error()))
	}

	res := exprtkObj.GetEvaluatedValue()

	return ctx.Status(fiber.StatusOK).JSON(model.OKResponse(res))
}