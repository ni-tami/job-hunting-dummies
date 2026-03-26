package workflows

import (
	"time"

	cActivities "github.com/ni-tami/job-hunting-dummies-workflows/app/populate_company/activities"
	"github.com/ni-tami/job-hunting-dummies-workflows/app/populate_company/models"
	sharedactivities "github.com/ni-tami/job-hunting-dummies-workflows/app/shared/activities"
	"go.temporal.io/sdk/converter"
	"go.temporal.io/sdk/workflow"
)

// PopulateCompanyWorkflow workflow definition
func PopulateCompanyWorkflow(ctx workflow.Context, parallelism int) (results []models.CreateCompany, err error) {
	// Extract schedule metadata
	info := workflow.GetInfo(ctx)

	//lint:ignore SA1019 - this is a sample
	scheduledByIDPayload := info.SearchAttributes.IndexedFields["TemporalScheduledById"]
	var scheduledByID string
	err = converter.GetDefaultDataConverter().FromPayload(scheduledByIDPayload, &scheduledByID)
	if err != nil {
		return []models.CreateCompany{}, err
	}
	//lint:ignore SA1019 - this is a sample
	startTimePayload := info.SearchAttributes.IndexedFields["TemporalScheduledStartTime"]
	var startTime time.Time
	err = converter.GetDefaultDataConverter().FromPayload(startTimePayload, &startTime)
	if err != nil {
		return []models.CreateCompany{}, err
	}

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx1 := workflow.WithActivityOptions(ctx, ao)

	for i := 0; i < parallelism; i++ {
		// Start a goroutine in a workflow safe way
		workflow.Go(ctx1, func(gCtx workflow.Context) {
			var (
				randCompany     models.CreateCompany
				JobHuntActivity *sharedactivities.JobHuntServiceActivity
			)
			err = workflow.ExecuteChildWorkflow(gCtx, cActivities.ExecuteExternalGenerateCompanyChildWorkflow).Get(gCtx, &randCompany)
			if err != nil {
				// Very naive error handling. Only the last error will be returned by the workflow
				return
			}
			err = workflow.ExecuteActivity(gCtx, JobHuntActivity.CreateCompanyRPCActivity, randCompany).Get(gCtx, &err)
			if err != nil {
				return
			}
			results = append(results, randCompany)
		})
	}

	_ = workflow.Await(ctx, func() bool {
		return err != nil || len(results) == parallelism
	})
	return
}
