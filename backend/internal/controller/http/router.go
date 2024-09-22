package http

import (
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/internal/controller/http/model/validator"
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/internal/core"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Router struct {
	app            *fiber.App
	formValidator  validator.FormValidatorService
	authService    core.AuthService
	historyService core.HistoryService
}

func NewRouter(
	app 		   *fiber.App,
	formValidator  validator.FormValidatorService,
	authService    core.AuthService,
	historyService core.HistoryService,
) {
	router := &Router{
		app:            app,
		formValidator:  formValidator,
		authService:    authService,
		historyService: historyService,
	}

	router.initRequestMiddlewares()

	router.initRoutes()

	router.initResponseMiddlewares()
}

func (r *Router) initRoutes() {
	r.app.Get("/message", r.message)

	v1 := r.app.Group("/api/v1")

	// calculator
	v1.Get("/calculate", r.calculate)

	// history
	v1.Get("/history", r.protectedMiddleware(), r.getAllHistory)
	v1.Post("/history", r.protectedMiddleware(), r.saveHistory)
	v1.Delete("/history", r.protectedMiddleware(), r.deleteHistory)
	v1.Delete("/history/:record_id", r.protectedMiddleware(), r.deleteHistoryByID)

	// auth
	v1.Post("/auth/signup", r.signup)

	// e.g. protected
	v1.Get("/protected", r.protectedMiddleware(), r.protected)
}

// initRequestMiddlewares initializes all middlewares for http requests
func (r *Router) initRequestMiddlewares() {
	r.app.Use(logger.New())
}

// initResponseMiddlewares initializes all middlewares for http response
func (r *Router) initResponseMiddlewares() {}
