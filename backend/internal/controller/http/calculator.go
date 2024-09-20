package http

import (
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/internal/controller/http/model/calculator"
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/Pramod-Devireddy/go-exprtk"
)

func (r *Router) calculate(ctx *fiber.Ctx) error {
	var calculator calculator.Calculator 
	
	if err := ctx.QueryParser(&calculator); err != nil {
		logger.Log().Error(ctx.UserContext(), err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	exprtkObj := exprtk.NewExprtk()
	defer exprtkObj.Delete()

	exprtkObj.SetExpression(calculator.Expression);
	
	err := exprtkObj.CompileExpression()
	if err != nil {
		logger.Log().Error(ctx.UserContext(), err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(exprtkObj.GetEvaluatedValue())
}