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
		UserByIdLoader:          dataloadgen.NewLoader(r.GetUserByID, dataloadgen.WithWait(time.Millisecond), dataloadgen.WithBatchCapacity(3)),
		ApplicationByIdLoader:   dataloadgen.NewLoader(r.GetApplicationByID, dataloadgen.WithWait(time.Millisecond), dataloadgen.WithBatchCapacity(3)),
		JobByIdLoader:           dataloadgen.NewLoader(r.GetJobByID, dataloadgen.WithWait(time.Millisecond), dataloadgen.WithBatchCapacity(3)),
		CompanyByIdLoader:       dataloadgen.NewLoader(r.GetCompanyByID, dataloadgen.WithWait(time.Millisecond), dataloadgen.WithBatchCapacity(3)),
		ApplicantByIdLoader:     dataloadgen.NewLoader(r.GetApplicantByID, dataloadgen.WithWait(time.Millisecond), dataloadgen.WithBatchCapacity(3)),
		UsersByIdsLoader:        dataloadgen.NewLoader(r.GetUsersByIds, dataloadgen.WithWait(time.Millisecond), dataloadgen.WithBatchCapacity(3)),
		ApplicationsByIdsLoader: dataloadgen.NewLoader(r.GetApplicationsByIds, dataloadgen.WithWait(time.Millisecond), dataloadgen.WithBatchCapacity(3)),
		JobsByIdsLoader:         dataloadgen.NewLoader(r.GetJobsByIds, dataloadgen.WithWait(time.Millisecond), dataloadgen.WithBatchCapacity(3)),
		CompaniesByIdsLoader:    dataloadgen.NewLoader(r.GetCompaniesByIds, dataloadgen.WithWait(time.Millisecond), dataloadgen.WithBatchCapacity(3)),
		ApplicantsByIdsLoader:   dataloadgen.NewLoader(r.GetApplicantsByIds, dataloadgen.WithWait(time.Millisecond), dataloadgen.WithBatchCapacity(3)),
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

// // GetApplicationsByApplicantID returns many applications by applicant id efficiently
// func GetApplicationsByApplicantID(ctx context.Context, applicantId int64) ([]*model.Application, error) {
// 	loaders := For(ctx)
// 	return loaders.ApplicationLoader.Load(ctx, applicantId)
// }

// // GetJobsByCompanyID returns many jobs by company id efficiently
// func GetJobsByCompanyID(ctx context.Context, companyId int64) ([]*model.Job, error) {
// 	loaders := For(ctx)
// 	return loaders.JobLoader.Load(ctx, companyId)
// }

// // GetApplicationsByJobID returns many applications by job id efficiently
// func GetApplicationsByJobID(ctx context.Context, jobId int64) ([]*model.Application, error) {
// 	loaders := For(ctx)
// 	return loaders.ApplicationByIdLoader.Load(ctx, jobId)
// }

// GetUserByID returns single user by id efficiently
func GetUserByID(ctx context.Context, userId string) (*model.User, error) {
	loaders := For(ctx)
	return loaders.UserByIdLoader.Load(ctx, userId)
}

// GetJobByID returns Job by ID efficiently
func GetJobByID(ctx context.Context, jobId int64) ([]*model.Application, error) {
	loaders := For(ctx)
	return loaders.JobByIdLoader.Load(ctx, jobId)
}

// GetCompanyByID returns Company by ID efficiently
func GetCompanyByID(ctx context.Context, companyId int64) ([]*model.Company, error) {
	loaders := For(ctx)
	return loaders.CompanyByIdLoader.Load(ctx, companyId)
}

// GetApplicantByID returns Applicant by ID efficiently
func GetApplicantByID(ctx context.Context, applicantId int64) ([]*model.Applicant, error) {
	loaders := For(ctx)
	return loaders.ApplicantByIdLoader.Load(ctx, applicantId)
}

// GetApplicationByID returns Application by ID efficiently
func GetApplicationByID(ctx context.Context, applicationId int64) ([]*model.Application, error) {
	loaders := For(ctx)
	return loaders.ApplicationByIdLoader.Load(ctx, applicationId)
}

// GetUsersByIds returns single user by id efficiently
func GetUsersByIds(ctx context.Context, userIds []string) (*model.User, error) {
	loaders := For(ctx)
	return loaders.UsersByIdsLoader.LoadAll(ctx, userIds)
}

// GetJobsByIds returns Job by ID efficiently
func GetJobsByIds(ctx context.Context, jobIds []int64) ([]*model.Application, error) {
	loaders := For(ctx)
	return loaders.JobsByIdsLoader.LoadAll(ctx, jobIds)
}

// GetCompaniesByIds returns Company by ID efficiently
func GetCompaniesByIds(ctx context.Context, companyIds []int64) ([]*model.Company, error) {
	loaders := For(ctx)
	return loaders.CompaniesByIdsLoader.LoadAll(ctx, companyIds)
}

// GetApplicantsByIds returns Applicant by ID efficiently
func GetApplicantsByIds(ctx context.Context, applicantIds []int64) ([]*model.Applicant, error) {
	loaders := For(ctx)
	return loaders.ApplicantsByIdsLoader.LoadAll(ctx, applicantIds)
}

// GetApplicationsByIds returns Application by ID efficiently
func GetApplicationsByIds(ctx context.Context, applicationIds []int64) ([]*model.Application, error) {
	loaders := For(ctx)
	return loaders.ApplicationsByIdsLoader.LoadAll(ctx, applicationIds)
}
