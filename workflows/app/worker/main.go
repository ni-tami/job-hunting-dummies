package main

import (
	"log"

	"github.com/ni-tami/job-hunting-dummies-workflows/app/internal/activities"
	grpcClient "github.com/ni-tami/job-hunting-dummies-workflows/app/internal/client"
	"github.com/ni-tami/job-hunting-dummies-workflows/app/internal/config"
	"github.com/ni-tami/job-hunting-dummies-workflows/app/internal/workflows"
	pb "github.com/ni-tami/job-hunting-dummies-workflows/app/pb/job_hunting_dummies"
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

	conn := grpcClient.NewGrpcClient()
	defer func() {
		_ = conn.Close()
	}()

	grpcClient := pb.NewJobHuntServiceClient(conn)

	srvActivities := activities.NewJobHuntServiceActivity(grpcClient)
	w := worker.New(c, config.PopulateUserTaskQueueName, worker.Options{})

	w.RegisterWorkflow(workflows.PopulateUserWorkflow)
	w.RegisterActivity(srvActivities)
	w.RegisterActivity(activities.FetchRandomUserActivity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
