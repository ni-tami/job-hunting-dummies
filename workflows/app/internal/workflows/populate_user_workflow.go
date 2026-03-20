package workflows

import (
	"time"

	"github.com/ni-tami/job-hunting-dummies-workflows/app/internal/activities"
	"github.com/ni-tami/job-hunting-dummies-workflows/app/internal/models"
	"go.temporal.io/sdk/converter"
	"go.temporal.io/sdk/workflow"
)

/**
* selfnote: Modified from go sample goroutine + schedules
* This sample workflow demonstrates how to use multiple Temporal goroutines (instead of native goroutine) to process a
* a sequence of activities in parallel.
* In Temporal workflow, you should create goroutines using workflow.Go method.
 */

// PopulateUserWorkflow workflow definition
func PopulateUserWorkflow(ctx workflow.Context, parallelism int) (results []models.CreateUser, err error) {
	// Extract schedule metadata
	info := workflow.GetInfo(ctx)

	// Workflow Executions started by a Schedule have the following additional properties appended to their search attributes
	//lint:ignore SA1019 - this is a sample
	scheduledByIDPayload := info.SearchAttributes.IndexedFields["TemporalScheduledById"]
	var scheduledByID string
	err = converter.GetDefaultDataConverter().FromPayload(scheduledByIDPayload, &scheduledByID)
	if err != nil {
		return []models.CreateUser{}, err
	}
	//lint:ignore SA1019 - this is a sample
	startTimePayload := info.SearchAttributes.IndexedFields["TemporalScheduledStartTime"]
	var startTime time.Time
	err = converter.GetDefaultDataConverter().FromPayload(startTimePayload, &startTime)
	if err != nil {
		return []models.CreateUser{}, err
	}

	// Set activity options on ctx1 (like schedule example)
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx1 := workflow.WithActivityOptions(ctx, ao)

	for i := 0; i < parallelism; i++ {
		// Start a goroutine in a workflow safe way
		workflow.Go(ctx1, func(gCtx workflow.Context) {
			var randUser models.CreateUser
			err = workflow.ExecuteActivity(gCtx, activities.FetchRandomUserActivity).Get(gCtx, &randUser)
			if err != nil {
				// Very naive error handling. Only the last error will be returned by the workflow
				return
			}
			var JobHuntActivity *activities.JobHuntServiceActivity
			err = workflow.ExecuteActivity(gCtx, JobHuntActivity.CreateUserRPCActivity, randUser).Get(gCtx, &err)
			if err != nil {
				return
			}
			results = append(results, randUser)
		})
	}

	_ = workflow.Await(ctx, func() bool {
		return err != nil || len(results) == parallelism
	})
	return
}
