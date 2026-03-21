package dataloader

import (
	"context"
	"net/http"
	"time"

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
	UserByIdLoader        *dataloadgen.Loader[int64, model.User]
	ApplicationByIdLoader *dataloadgen.Loader[int64, model.Application]
	JobByIdLoader         *dataloadgen.Loader[int64, model.Job]
	CompanyByIdLoader     *dataloadgen.Loader[int64, model.Company]
	ApplicantByIdLoader   *dataloadgen.Loader[int64, model.Applicant]
}

func NewLoaders(conn *gorm.DB) *Loaders {
	r := repository.NewJobPortalRepository(conn)
	return &Loaders{
		UserByIdLoader:        dataloadgen.NewLoader(r.GetUsersByIds, dataloadgen.WithWait(time.Millisecond), dataloadgen.WithBatchCapacity(3)),
		ApplicationByIdLoader: dataloadgen.NewLoader(r.GetApplicationsByIds, dataloadgen.WithWait(time.Millisecond), dataloadgen.WithBatchCapacity(3)),
		JobByIdLoader:         dataloadgen.NewLoader(r.GetJobsByIds, dataloadgen.WithWait(time.Millisecond), dataloadgen.WithBatchCapacity(3)),
		CompanyByIdLoader:     dataloadgen.NewLoader(r.GetCompaniesByIds, dataloadgen.WithWait(time.Millisecond), dataloadgen.WithBatchCapacity(3)),
		ApplicantByIdLoader:   dataloadgen.NewLoader(r.GetApplicantsByIds, dataloadgen.WithWait(time.Millisecond), dataloadgen.WithBatchCapacity(3)),
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

// GetUserByID returns single user by id efficiently
func GetUserByID(ctx context.Context, userId int64) (model.User, error) {
	loaders := For(ctx)
	return loaders.UserByIdLoader.Load(ctx, userId)
}

// GetJobByID returns Job by ID efficiently
func GetJobByID(ctx context.Context, jobId int64) (model.Job, error) {
	loaders := For(ctx)
	return loaders.JobByIdLoader.Load(ctx, jobId)
}

// GetCompanyByID returns Company by ID efficiently
func GetCompanyByID(ctx context.Context, companyId int64) (model.Company, error) {
	loaders := For(ctx)
	return loaders.CompanyByIdLoader.Load(ctx, companyId)
}

// GetApplicantByID returns Applicant by ID efficiently
func GetApplicantByID(ctx context.Context, applicantId int64) (model.Applicant, error) {
	loaders := For(ctx)
	return loaders.ApplicantByIdLoader.Load(ctx, applicantId)
}

// GetApplicationByID returns Application by ID efficiently
func GetApplicationByID(ctx context.Context, applicationId int64) (model.Application, error) {
	loaders := For(ctx)
	return loaders.ApplicationByIdLoader.Load(ctx, applicationId)
}

// GetUsersByIds returns single user by id efficiently
func GetUsersByIds(ctx context.Context, userIds []int64) ([]model.User, error) {
	loaders := For(ctx)
	return loaders.UserByIdLoader.LoadAll(ctx, userIds)
}

// GetJobsByIds returns Job by ID efficiently
func GetJobsByIds(ctx context.Context, jobIds []int64) ([]model.Job, error) {
	loaders := For(ctx)
	return loaders.JobByIdLoader.LoadAll(ctx, jobIds)
}

// GetCompaniesByIds returns Company by ID efficiently
func GetCompaniesByIds(ctx context.Context, companyIds []int64) ([]model.Company, error) {
	loaders := For(ctx)
	return loaders.CompanyByIdLoader.LoadAll(ctx, companyIds)
}

// GetApplicantsByIds returns Applicant by ID efficiently
func GetApplicantsByIds(ctx context.Context, applicantIds []int64) ([]model.Applicant, error) {
	loaders := For(ctx)
	return loaders.ApplicantByIdLoader.LoadAll(ctx, applicantIds)
}

// GetApplicationsByIds returns Application by ID efficiently
func GetApplicationsByIds(ctx context.Context, applicationIds []int64) ([]model.Application, error) {
	loaders := For(ctx)
	return loaders.ApplicationByIdLoader.LoadAll(ctx, applicationIds)
}
