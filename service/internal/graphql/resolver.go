// nolint:gci
package graphql

import "github.com/ni-tami/job-hunting-dummies-service/internal/graphql/model"
import "github.com/ni-tami/job-hunting-dummies-service/internal/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

type Resolver struct {
	applications []*model.Application
	companies []*model.Company
	applicants []*model.Applicant
	jobs []*model.Job
	users []*model.User
	jobPortalUsecase usecase.JobPortalUsecase
}
