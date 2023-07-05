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
			r.Put("/{equipmentID}", app.updateEquipmentHandler)
		})

		// MusclesGroup resources
		r.Route("/musclesgroup", func(r chi.Router) {
			r.Post("/", app.createMusclesGroupHandler)
			r.Get("/", app.listMusclesGroupHandler)
			r.Put("/{musclesgroupID}", app.updateMusclesGroupHandler)
		})

		// ExerciseType resources
		r.Route("/exercisetype", func(r chi.Router) {
			r.Post("/", app.createExerciseTypeHandler)
			r.Get("/", app.listExerciseTypeHandler)
			r.Put("/{exercisetypeID}", app.updateExerciseTypeHandler)
		})

		// Exercises resources
		r.Route("/exercises", func(r chi.Router) {
			r.Post("/", app.createExerciseHandler)
			r.Get("/", app.listExerciseHandler)
			r.Put("/{exerciseID}", app.updateExerciseHandler)
		})

		// Routine resources
		r.Route("/routines", func(r chi.Router) {
			r.Post("/", app.requireActivatedUser(app.createRoutineHandler))
			r.Get("/", app.requireActivatedUser(app.listRoutineHandler))
			r.Put("/{routineID}", app.requireActivatedUser(app.updateRoutineHandler))
			r.Delete("/{routineID}", app.requireActivatedUser(app.deleteRoutineHandler))
		})

		// RoutineExercise resources
		r.Route("/routineExercises", func(r chi.Router) {
			r.Post("/", app.requireActivatedUser(app.createRoutineExerciseHandler))
			r.Get("/{routineID}", app.requireActivatedUser(app.getRoutineExerciseHandler))
			r.Put(
				"/{routineExerciseID}",
				app.requireActivatedUser(app.updateRoutineExerciseHandler),
			)
			r.Delete(
				"/{routineExerciseID}",
				app.requireActivatedUser(app.deleteRoutineExerciseHandler),
			)
		})
	})

	return mux
}
