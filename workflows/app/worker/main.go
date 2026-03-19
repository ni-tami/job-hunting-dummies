package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/ni-tami/job-hunting-dummies-workflows/app"
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

	w := worker.New(c, "schedule-goroutine", worker.Options{})

	w.RegisterWorkflow(app.SampleScheduledGoroutineWorkflow)
	w.RegisterActivity(app.Step1)
	w.RegisterActivity(app.Step2)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
