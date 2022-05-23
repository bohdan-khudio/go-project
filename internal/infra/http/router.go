package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/test_server/internal/infra/http/controllers"
)

func Router(eventController *controllers.EventController, userController *controllers.UserController) http.Handler {
	router := chi.NewRouter()

	// Health
	router.Group(func(healthRouter chi.Router) {
		healthRouter.Use(middleware.RedirectSlashes)

		healthRouter.Route("/ping", func(healthRouter chi.Router) {
			healthRouter.Get("/", PingHandler())

			healthRouter.Handle("/*", NotFoundJSON())
		})
	})

	router.Group(func(apiRouter chi.Router) {
		apiRouter.Use(middleware.RedirectSlashes)

		apiRouter.Route("/v1", func(apiRouter chi.Router) {

			apiRouter.Group(func(apiRouter chi.Router) {
				AddEventRoutes(&apiRouter, eventController)
				AddUserRoutes(&apiRouter, userController)

				apiRouter.Handle("/*", NotFoundJSON())
			})
			apiRouter.Handle("/*", NotFoundJSON())
		})
	})

	return router
}

func AddEventRoutes(router *chi.Router, eventController *controllers.EventController) {
	(*router).Route("/events", func(apiRouter chi.Router) {
		apiRouter.Get(
			"/",
			eventController.FindAll(),
		)
		apiRouter.Get(
			"/{id}",
			eventController.FindOne(),
		)
		apiRouter.Post(
			"/",
			eventController.Create(),
		)
		apiRouter.Put(
			"/",
			eventController.Update(),
		)
		apiRouter.Delete(
			"/{id}",
			eventController.Delete())
	})
}

func AddUserRoutes(router *chi.Router, userController *controllers.UserController) {
	(*router).Route("/users", func(apiRouter chi.Router) {
		apiRouter.Get(
			"/",
			userController.FindAll(),
		)
		apiRouter.Get(
			"/{id}",
			userController.FindOne(),
		)
	})
}
