package activities

import (
	"context"
	"fmt"

	"github.com/ni-tami/job-hunting-dummies-workflows/app/populate_company/models"
	"github.com/ni-tami/job-hunting-dummies-workflows/app/shared/config"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
)

const (
	pyTaskQueue  = "generate-company-py-queue"
	activityName = "generate-company-activity"
)

type GenerateCompanyActivity struct {
	tc client.Client
}

func NewGenerateCompanyActivity(tc client.Client) *GenerateCompanyActivity {
	return &GenerateCompanyActivity{
		tc: tc,
	}
}

// // GenerateCompanyWorkflow simply returns the result of the generate-company activity.
// func GenerateCompanyWorkflow(ctx workflow.Context) (models.CreateCompany, error) {
// 	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
// 		// Give it only 5 seconds to schedule and run with no retries
// 		ScheduleToCloseTimeout: 5 * time.Second,
// 		RetryPolicy:            &temporal.RetryPolicy{MaximumAttempts: 1},
// 	})
// 	var response models.CreateCompany
// 	err := workflow.ExecuteActivity(ctx, activityName).Get(ctx, &response)
// 	return response, err
// }

// ExecuteExternalGenerateCompanyActivity execute python workflow and returns result
func (a *GenerateCompanyActivity) ExecuteExternalGenerateCompanyActivity(ctx context.Context) (models.CreateCompany, error) {
	logger := activity.GetLogger(ctx)

	var randCompanyResp models.CreateCompany
	workflowRun, err := a.tc.ExecuteActivity(ctx, client.StartActivityOptions{
		TaskQueue: config.PopulateCompanyTaskQueueName,
	}, activityName,
	)
	if err != nil {
		logger.Info("Unable to query external python workflow %v", err.Error())
	}
	err = workflowRun.Get(ctx, &randCompanyResp)
	logger.Info(fmt.Sprintf("Execute activity %s: %v.", activityName, randCompanyResp))
	return randCompanyResp, err
}
