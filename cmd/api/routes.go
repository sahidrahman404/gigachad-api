package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.NotFound(app.notFound)
	mux.MethodNotAllowed(app.methodNotAllowed)

	mux.Use(app.recoverPanic)

	mux.Get("/status", app.status)

	mux.Route("/v1", func(r chi.Router) {
		// users resources
		r.Route("/users", func(r chi.Router) {
			// user registration
			r.Post("/", app.registerUserHandler)
			// user activation
			r.Put("/activated", app.activateUserHandler)
			// password reset
			r.Put("/password", app.updateUserPasswordHandler)
		})

		// tokens resources
		r.Route("/tokens", func(r chi.Router) {
			// user login
			r.Post("/authentication", app.createAuthenticationTokenHandler)

			// user activation
			r.Post("/activation", app.createActivationTokenHandler)

			// password reset
			r.Post(
				"/password-reset",
				app.createPasswordResetTokenHandler,
			)
		})

		// equipment resources
		r.Route("/equipments", func(r chi.Router) {
			r.Post("/", app.createEquipmentHandler)
			r.Get("/", app.listEquipmentHandler)
		})
	})

	return mux
}
