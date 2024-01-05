package main

import (
	"log"

	"github.com/sahidrahman404/gigachad-api/internal/database"
	"github.com/sahidrahman404/gigachad-api/internal/env"
	"github.com/sahidrahman404/gigachad-api/internal/shared"
	"github.com/sahidrahman404/gigachad-api/internal/smtp"
	"github.com/sahidrahman404/gigachad-api/internal/workoutreminder"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

type config struct {
	db struct {
		dsn string
	}
	smtp struct {
		host     string
		username string
		password string
		from     string
		port     int
	}
}

func main() {
	var cfg config
	cfg.db.dsn = env.GetString("DB_DSN", "db.sqlite")
	cfg.smtp.host = env.GetString("SMTP_HOST", "example.smtp.host")
	cfg.smtp.port = env.GetInt("SMTP_PORT", 25)
	cfg.smtp.username = env.GetString("SMTP_USERNAME", "example_username")
	cfg.smtp.password = env.GetString("SMTP_PASSWORD", "pa55word")
	cfg.smtp.from = env.GetString("SMTP_FROM", "Example Name <no_reply@example.org>")
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client.", err)
	}
	defer c.Close()

	w := worker.New(c, shared.DailyWorkoutReminderTaskQueueName, worker.Options{})
	w.RegisterWorkflow(smtp.SendDailyWorkoutReminderWorkflow)
	w.RegisterWorkflow(workoutreminder.CreateDailyWorkoutScheduleWorkflow)
	w.RegisterWorkflow(workoutreminder.UpdateDailyWorkoutScheduleWorkflow)

	mailer := smtp.NewMailer(
		cfg.smtp.host,
		cfg.smtp.port,
		cfg.smtp.username,
		cfg.smtp.password,
		cfg.smtp.from,
	)

	w.RegisterActivity(mailer)

	db, err := database.New(cfg.db.dsn)
	if err != nil {
		log.Fatalln("unable to create DB connection", err)
	}
	defer db.Db.Close()

	workoutReminder := workoutreminder.NewWorkoutReminder(&c, db.Ent)

	w.RegisterActivity(workoutReminder)

	// Start listening to the Task Queue.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
