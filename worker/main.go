package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	ridis "github.com/cschleiden/go-workflows/backend/redis"
	"github.com/cschleiden/go-workflows/diag"
	"github.com/cschleiden/go-workflows/worker"
	"github.com/redis/go-redis/v9"
	"github.com/sahidrahman404/gigachad-api/internal/activity"
	"github.com/sahidrahman404/gigachad-api/internal/env"
	"github.com/sahidrahman404/gigachad-api/internal/workflow"
)

type config struct {
	address  string
	password string
	smtp     struct {
		host     string
		username string
		password string
		from     string
		port     int
	}
}

func main() {
	var cfg config
	cfg.address = env.GetString("REDIS_ADDRESS", "localhost:6379")
	cfg.password = env.GetString("REDIS_PASSWORD", "")
	cfg.smtp.host = env.GetString("SMTP_HOST", "example.smtp.host")
	cfg.smtp.port = env.GetInt("SMTP_PORT", 25)
	cfg.smtp.username = env.GetString("SMTP_USERNAME", "example_username")
	cfg.smtp.password = env.GetString("SMTP_PASSWORD", "pa55word")

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.address,
		Password: cfg.password,
		DB:       0,
	})
	b, err := ridis.NewRedisBackend(rdb, ridis.WithAutoExpiration(time.Hour*48))
	if err != nil {
		panic("could not start redis")
	}
	w := worker.New(b, nil)
	w.RegisterWorkflow(workflow.SendEmail)
	a := activity.NewMailer(cfg.smtp.host, cfg.smtp.port, cfg.smtp.username, cfg.smtp.password, cfg.smtp.from)
	w.RegisterActivity(a)
	if err := w.Start(context.Background()); err != nil {
		panic("could not start worker")
	}
	m := http.NewServeMux()
	m.Handle("/diag/", http.StripPrefix("/diag", diag.NewServeMux(b)))
	go http.ListenAndServe(":3500", m)

	c2 := make(chan os.Signal, 1)
	signal.Notify(c2, os.Interrupt)
	<-c2
}
