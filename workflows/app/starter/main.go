package main

import (
	"context"
	"log"
	"time"

	"github.com/ni-tami/job-hunting-dummies-workflows/app/internal/config"
	"github.com/ni-tami/job-hunting-dummies-workflows/app/internal/workflows"
	"github.com/pborman/uuid"
	"go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/client"
)

func main() {
	config.Load()

	ctx := context.Background()
	// The client is a heavyweight object that should be created once per process.
	c, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// This schedule ID can be user business logic identifier as well.
	scheduleID := "populate_user_" + uuid.New()
	workflowID := "populate_user_workflow_" + uuid.New()
	// Create the schedule, start with no spec so the schedule will not run.
	scheduleHandle, err := c.ScheduleClient().Create(ctx, client.ScheduleOptions{
		ID:   scheduleID,
		Spec: client.ScheduleSpec{},
		Action: &client.ScheduleWorkflowAction{
			ID:        workflowID,
			Workflow:  workflows.PopulateUserWorkflow,
			TaskQueue: config.PopulateUserTaskQueueName,
			Args:      []interface{}{1},
		},
	})
	if err != nil {
		log.Fatalln("Unable to create schedule", err)
	}

	// Manually trigger the schedule once
	log.Println("Manually triggering schedule", "ScheduleID", scheduleHandle.GetID())

	err = scheduleHandle.Trigger(ctx, client.ScheduleTriggerOptions{
		Overlap: enums.SCHEDULE_OVERLAP_POLICY_ALLOW_ALL,
	})
	if err != nil {
		log.Fatalln("Unable to trigger schedule", err)
	}

	// Update the schedule with a spec so it will run periodically,
	log.Println("Updating schedule", "ScheduleID", scheduleHandle.GetID())
	err = scheduleHandle.Update(ctx, client.ScheduleUpdateOptions{
		DoUpdate: func(schedule client.ScheduleUpdateInput) (*client.ScheduleUpdate, error) {
			schedule.Description.Schedule.Spec = &client.ScheduleSpec{
				// Run the schedule every 5s
				Intervals: []client.ScheduleIntervalSpec{
					{
						Every: config.PopulateUserScheduleIntervalHour,
					},
				},
			}

			return &client.ScheduleUpdate{
				Schedule: &schedule.Description.Schedule,
			}, nil
		},
	})
	if err != nil {
		log.Fatalln("Unable to update schedule", err)
	}

	log.Println("Waiting for schedule to complete actions", "ScheduleID", scheduleHandle.GetID())

	for {
		description, err := scheduleHandle.Describe(ctx)
		if err != nil {
			log.Fatalln("Unable to describe schedule", err)
		}
		if description.Schedule.State.RemainingActions != 0 {
			log.Println("Schedule has remaining actions", "ScheduleID", scheduleHandle.GetID(), "RemainingActions", description.Schedule.State.RemainingActions)
			time.Sleep(30 * time.Second)
		} else {
			break
		}
	}
}
