package main

import (
	"net/http"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/sahidrahman404/gigachad-api/internal/gql"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.NotFound(app.notFound)
	mux.MethodNotAllowed(app.methodNotAllowed)

	mux.Use(app.recoverPanic)

	mux.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://gigachad.buzz", "http://localhost:3000"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	mux.Use(app.authenticate)

	mux.Get("/status", app.status)

	srv := handler.NewDefaultServer(
		gql.NewSchema(app.ent, app.mailer, app.storage, app.logger, &app.wg, &app.config.Imgproxy, &app.config.AWSConfig, app.purifier, app.temporalClient),
	)
	srv.Use(entgql.Transactioner{TxOpener: app.ent})
	mux.Handle("/gql", playground.Handler("Gigachad", "/query"))
	mux.Handle("/query", srv)

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

			// Add token to client cookie using http cookies - for spa
			r.Get("/set/{tokenPlainText}", app.setCookieHandler)
			r.Get("/get", app.getCookieHandler)
			r.Delete("/delete/{tokenPlainText}", app.deleteCookieHandler)
			r.Post("/img", app.getSignedUrl)
			r.Post("/imgs", app.getTransformedUrls)
			r.Post("/sign-s3", app.getUploadURL)
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
			r.Get("/{routineID}", app.requireActivatedUser(app.getRoutineHandler))
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

		// Workout resources
		r.Route("/workouts", func(r chi.Router) {
			r.Post("/", app.requireActivatedUser(app.createWorkoutHandler))
			r.Get("/", app.requireActivatedUser(app.listWorkoutHandler))
			r.Get("/{workoutID}", app.requireActivatedUser(app.getWorkoutHandler))
		})

		// WorkoutLog resources
		r.Route("/workoutLogs", func(r chi.Router) {
			r.Post("/", app.requireActivatedUser(app.createWorkoutLogHandler))
		})
	})

	return mux
}
