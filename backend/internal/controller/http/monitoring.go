package http

import "github.com/gofiber/fiber/v2"

func (r *Router) messageFromMaxim(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).SendString("Счастливых Вам Голодных игр и пусть удача всегда будет с вами")
}
