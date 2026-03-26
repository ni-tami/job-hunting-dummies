package activities

import (
	"fmt"
	"time"

	"github.com/ni-tami/job-hunting-dummies-workflows/app/populate_company/models"
	"github.com/ni-tami/job-hunting-dummies-workflows/app/shared/config"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// ExecuteExternalGenerateCompanyChildWorkflow as handler for python activity
func ExecuteExternalGenerateCompanyChildWorkflow(ctx workflow.Context) (models.CreateCompany, error) {
	logger := workflow.GetLogger(ctx)

	var randCompanyResp models.CreateCompany
	aCtx := workflow.WithActivityOptions(
		ctx, workflow.ActivityOptions{
			TaskQueue:              config.PyPopulateCompanyTaskQueueName,
			ScheduleToCloseTimeout: 5 * time.Second,
			RetryPolicy:            &temporal.RetryPolicy{MaximumAttempts: 1},
		},
	)

	generateCompanyActivity := workflow.ExecuteActivity(aCtx, config.PyPopulateActivityName)
	if generateCompanyActivity == nil {
		// TODO: refactor
		workflowNotFoundErr := fmt.Errorf("cannot execute child workflow %s", config.PyPopulateActivityName)
		logger.Error(workflowNotFoundErr.Error())
		return randCompanyResp, workflowNotFoundErr
	}
	err := generateCompanyActivity.Get(aCtx, &randCompanyResp)
	if err != nil {
		logger.Error(fmt.Sprintf("Execute activity %s: %v.", config.PyPopulateActivityName, randCompanyResp))
		return randCompanyResp, err
	}
	return randCompanyResp, nil
}
