package client

import (
	"context"
	"net/http"
	"time"

	// gqlModel "github.com/ni-tami/job-hunting-dummies-service/internal/graphql/model"
	model "github.com/ni-tami/job-hunting-dummies-service/internal/model"
	"github.com/ni-tami/job-hunting-dummies-service/internal/repository"
	"github.com/vikstrous/dataloadgen"
	"gorm.io/gorm"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

type Loaders struct {
	UserLoader        *dataloadgen.Loader[string, *model.User]
	ApplicationLoader *dataloadgen.Loader[string, *model.Application]
	JobLoader         *dataloadgen.Loader[string, *model.Job]
	CompanyLoader     *dataloadgen.Loader[string, *model.Company]
	ApplicantLoader   *dataloadgen.Loader[string, *model.Applicant]
}

func NewLoaders(conn *gorm.DB) *Loaders {
	r := repository.NewJobPortalRepository(conn)
	return &Loaders{
		UserLoader:        dataloadgen.NewLoader(r.GetUserByID, dataloadgen.WithWait(time.Millisecond), dataloadgen.WithBatchCapacity(3)),
		ApplicationLoader: dataloadgen.NewLoader(r.GetApplicationByID, dataloadgen.WithWait(time.Millisecond), dataloadgen.WithBatchCapacity(3)),
		JobLoader:         dataloadgen.NewLoader(r.GetJobByID, dataloadgen.WithWait(time.Millisecond), dataloadgen.WithBatchCapacity(3)),
		CompanyLoader:     dataloadgen.NewLoader(r.GetCompanyByID, dataloadgen.WithWait(time.Millisecond), dataloadgen.WithBatchCapacity(3)),
		ApplicantLoader:   dataloadgen.NewLoader(r.GetApplicantByID, dataloadgen.WithWait(time.Millisecond), dataloadgen.WithBatchCapacity(3)),
	}
}

func Middleware(conn *gorm.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loader := NewLoaders(conn)
		r = r.WithContext(context.WithValue(r.Context(), loadersKey, loader))
		next.ServeHTTP(w, r)
	})
}

// For returns the dataloader for a given context
func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

// GetApplicationsByApplicantID returns many applications by applicant id efficiently
func GetApplicationsByApplicantID(ctx context.Context, applicantId int64) ([]*model.Application, error) {
	loaders := For(ctx)
	return loaders.ApplicationLoader.Load(ctx, applicantId)
}

// GetJobsByCompanyID returns many jobs by company id efficiently
func GetJobsByCompanyID(ctx context.Context, companyId int64) ([]*model.Job, error) {
	loaders := For(ctx)
	return loaders.JobLoader.Load(ctx, companyId)
}

// GetApplicationsByJobID returns many applications by job id efficiently
func GetApplicationsByJobID(ctx context.Context, jobId int64) ([]*model.Application, error) {
	loaders := For(ctx)
	return loaders.ApplicationLoader.Load(ctx, jobId)
}

// GetUser returns single user by id efficiently
func GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	loaders := For(ctx)
	return loaders.UserLoader.Load(ctx, userID)
}

// GetJobByID returns Job by ID efficiently
func GetJobByID(ctx context.Context, jobId int64) ([]*model.Application, error) {
	loaders := For(ctx)
	return loaders.ApplicationLoader.Load(ctx, jobId)
}

// GetCompanyByID returns Company by ID efficiently
func GetCompanyByID(ctx context.Context, companyId int64) ([]*model.Company, error) {
	loaders := For(ctx)
	return loaders.CompanyLoader.Load(ctx, companyId)
}

// GetApplicantByID returns Applicant by ID efficiently
func GetApplicantByID(ctx context.Context, applicantId int64) ([]*model.Applicant, error) {
	loaders := For(ctx)
	return loaders.ApplicantLoader.Load(ctx, applicantId)
}

// GetApplicationByID returns Application by ID efficiently
func GetApplicationByID(ctx context.Context, applicationId int64) ([]*model.Application, error) {
	loaders := For(ctx)
	return loaders.ApplicationLoader.Load(ctx, applicationId)
}
