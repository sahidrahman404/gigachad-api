package main

import (
	"log"

	"github.com/sahidrahman404/gigachad-api/internal/activity"
	"github.com/sahidrahman404/gigachad-api/internal/workflow"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// The client and worker are heavyweight objects that should be created once per process.
	c, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, activity.SendEmailTaskQueueName, worker.Options{})
	w.RegisterWorkflow(workflow.NewUser)
	a := activity.NewMailer("mailcrab.fly.dev", 1025, "a48f9c515e7157", "4e9e00cd4aec90", "Gigachad <no-reply@giagachad.io>")
	w.RegisterActivity(a)
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
