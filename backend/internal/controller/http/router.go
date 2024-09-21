package http

import (
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/internal/controller/http/model/validator"
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/internal/core"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Router struct {
	app           *fiber.App
	formValidator validator.FormValidatorService
	entityService core.EntityService
	authService   core.AuthService
}

func NewRouter(
	app *fiber.App,
	entityService core.EntityService,
	formValidator validator.FormValidatorService,
	authService core.AuthService,
) {
	router := &Router{
		app:           app,
		formValidator: formValidator,
		entityService: entityService,
		authService:   authService,
	}

	router.initRequestMiddlewares()

	router.initRoutes()

	router.initResponseMiddlewares()
}

func (r *Router) initRoutes() {
	r.app.Get("/message", r.messageFromMaxim)

	v1 := r.app.Group("/api/v1")

	// entities
	v1.Get("/entities", r.getEntities)
	v1.Get("/entities/:id", r.getEntityByID)

	// calculator
	v1.Get("/calculate", r.calculate)

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
