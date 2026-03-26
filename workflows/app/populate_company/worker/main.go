package main

import (
	"log"

	pb "github.com/ni-tami/job-hunting-dummies-workflows/app/pb/out/go/job_hunting_dummies"
	cActivities "github.com/ni-tami/job-hunting-dummies-workflows/app/populate_company/activities"
	cWorkflow "github.com/ni-tami/job-hunting-dummies-workflows/app/populate_company/workflows"
	sharedactivities "github.com/ni-tami/job-hunting-dummies-workflows/app/shared/activities"
	grpcClient "github.com/ni-tami/job-hunting-dummies-workflows/app/shared/client"
	"github.com/ni-tami/job-hunting-dummies-workflows/app/shared/config"
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

	// initiate activities
	// extActivities := cActivities.NewGenerateCompanyActivity(c)
	srvActivities := sharedactivities.NewJobHuntServiceActivity(grpcClient)
	// cActivities.NewGenerateCompanyActivity(c)

	w := worker.New(c, config.PopulateCompanyTaskQueueName, worker.Options{})

	w.RegisterWorkflow(cWorkflow.PopulateCompanyWorkflow)
	w.RegisterWorkflow(cActivities.ExecuteExternalGenerateCompanyChildWorkflow)
	w.RegisterActivity(srvActivities)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
