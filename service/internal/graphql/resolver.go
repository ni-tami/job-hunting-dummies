// nolint:gci
package graphql

import (
	"github.com/ni-tami/job-hunting-dummies-service/internal/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

type Resolver struct {
	jobPortalUsecase usecase.JobPortalUsecase
}

func NewResolver(jobPortalUsecase usecase.JobPortalUsecase) *Resolver {
	return &Resolver{
		jobPortalUsecase: jobPortalUsecase,
	}
}
