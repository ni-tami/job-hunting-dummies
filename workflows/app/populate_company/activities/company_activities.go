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
func ExecuteExternalGenerateCompanyChildWorkflow(ctx workflow.Context, companies []models.CreateCompany) ([]models.CreateCompany, error) {
	logger := workflow.GetLogger(ctx)

	var randCompaniesResp []models.CreateCompany
	aCtx := workflow.WithActivityOptions(
		ctx, workflow.ActivityOptions{
			TaskQueue:              config.PyPopulateCompanyTaskQueueName,
			ScheduleToCloseTimeout: 5 * time.Second,
			RetryPolicy:            &temporal.RetryPolicy{MaximumAttempts: 1},
		},
	)

	generateCompanyActivity := workflow.ExecuteActivity(aCtx, config.PyPopulateActivityName, companies)
	if generateCompanyActivity == nil {
		// TODO: refactor
		workflowNotFoundErr := fmt.Errorf("cannot execute child workflow %s", config.PyPopulateActivityName)
		logger.Error(workflowNotFoundErr.Error())
		return randCompaniesResp, workflowNotFoundErr
	}
	err := generateCompanyActivity.Get(aCtx, &randCompaniesResp)
	if err != nil {
		logger.Error(fmt.Sprintf("Execute activity %s: %v.", config.PyPopulateActivityName, randCompaniesResp))
		return randCompaniesResp, err
	}
	return randCompaniesResp, nil
}
