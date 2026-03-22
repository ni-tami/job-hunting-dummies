package main

import (
	"log"

	pb "github.com/ni-tami/job-hunting-dummies-workflows/app/pb/out/go/job_hunting_dummies"
	"github.com/ni-tami/job-hunting-dummies-workflows/app/populate_user/activities"
	grpcClient "github.com/ni-tami/job-hunting-dummies-workflows/app/populate_user/client"
	"github.com/ni-tami/job-hunting-dummies-workflows/app/populate_user/config"
	"github.com/ni-tami/job-hunting-dummies-workflows/app/populate_user/workflows"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	config.Load()
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
