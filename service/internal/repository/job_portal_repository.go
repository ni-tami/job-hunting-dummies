package repository

import (
	"context"
	"time"

	"github.com/ni-tami/job-hunting-dummies-service/internal/model"
	"gorm.io/gorm"
)

type (
	JobPortalRepository interface {
		CreateApplicant(ctx context.Context, newApplicant *model.Applicant) error
		CreateApplication(ctx context.Context, newApplication *model.Application) error
		CreateCompany(ctx context.Context, newCompany *model.Company) error
		CreateJob(ctx context.Context, newJob *model.Job) error
		CreateUser(ctx context.Context, newUser *model.User) error

		GetApplicantByID(ctx context.Context, id int64) (*model.Applicant, error)
		GetApplicationByID(ctx context.Context, id int64) (*model.Application, error)
		GetCompanyByID(ctx context.Context, id int64) (*model.Company, error)
		GetJobByID(ctx context.Context, id int64) (*model.Job, error)
		GetUserByID(ctx context.Context, id int64) (*model.User, error)

		GetApplicationsByApplicantID(ctx context.Context, applicantId int64) ([]*model.Application, error)
		GetApplicationsByJobID(ctx context.Context, jobId int64) ([]*model.Application, error)
		GetJobsByCompanyID(ctx context.Context, companyId int64) ([]*model.Job, error)

		GetJobs(ctx context.Context) ([]*model.Job, error)
		GetCompanies(ctx context.Context) ([]*model.Company, error)
		GetApplicants(ctx context.Context) ([]*model.Applicant, error)
		GetApplications(ctx context.Context) ([]*model.Application, error)
		GetUsers(ctx context.Context) ([]*model.User, error)

		UpdateApplicationByID(ctx context.Context, id int64, updatedApplication model.ApplicationUpdate) (int, error)
		UpdateCompanyByID(ctx context.Context, id int64, updatedCompany model.CompanyUpdate) (int, error)
		UpdateJobByID(ctx context.Context, id int64, updatedJob model.JobUpdate) (int, error)
		UpdateUserByID(ctx context.Context, id int64, updatedUser model.UserUpdate) (int, error)

		DeleteApplicantByID(ctx context.Context, id int64, deletionTime time.Time) (int, error)
		DeleteApplicationByID(ctx context.Context, id int64, deletionTime time.Time) (int, error)
		DeleteCompanyByID(ctx context.Context, id int64, deletionTime time.Time) (int, error)
		DeleteJobByID(ctx context.Context, id int64, deletionTime time.Time) (int, error)
		DeleteUserByID(ctx context.Context, id int64, deletionTime time.Time) (int, error)
	}

	jobPortalRepository struct {
		db *gorm.DB
	}
)

func NewJobPortalRepository(db *gorm.DB) JobPortalRepository {
	return jobPortalRepository{
		db: db,
	}
}

func (r jobPortalRepository) CreateApplicant(ctx context.Context, newApplicant *model.Applicant) error {
	err := r.db.WithContext(ctx).Save(&newApplicant).Error
	return err
}

func (r jobPortalRepository) CreateApplication(ctx context.Context, newApplication *model.Application) error {
	err := r.db.WithContext(ctx).Save(&newApplication).Error
	return err
}

func (r jobPortalRepository) CreateCompany(ctx context.Context, newCompany *model.Company) error {
	err := r.db.WithContext(ctx).Save(&newCompany).Error
	return err
}

func (r jobPortalRepository) CreateJob(ctx context.Context, newJob *model.Job) error {
	err := r.db.WithContext(ctx).Save(&newJob).Error
	return err
}

func (r jobPortalRepository) CreateUser(ctx context.Context, newUser *model.User) error {
	err := r.db.WithContext(ctx).Save(&newUser).Error
	return err
}

func (r jobPortalRepository) GetApplicantByID(ctx context.Context, id int64) (*model.Applicant, error) {
	var applicant *model.Applicant
	err := r.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&applicant).Error
	return applicant, err
}

func (r jobPortalRepository) GetApplicationByID(ctx context.Context, id int64) (*model.Application, error) {
	var application *model.Application
	err := r.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&application).Error
	return application, err
}

func (r jobPortalRepository) GetCompanyByID(ctx context.Context, id int64) (*model.Company, error) {
	var company *model.Company
	err := r.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&company).Error
	return company, err
}

func (r jobPortalRepository) GetJobByID(ctx context.Context, id int64) (*model.Job, error) {
	var job *model.Job
	err := r.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&job).Error
	return job, err
}

func (r jobPortalRepository) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	var user *model.User
	err := r.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&user).Error
	return user, err
}

func (r jobPortalRepository) GetApplicationsByApplicantID(ctx context.Context, applicantId int64) ([]*model.Application, error) {
	var applications []*model.Application
	err := r.db.WithContext(ctx).Where("deleted_at IS NULL AND applicant_id = ?", applicantId).Find(&applications).Error
	return applications, err
}

func (r jobPortalRepository) GetApplicationsByJobID(ctx context.Context, jobId int64) ([]*model.Application, error) {
	var applications []*model.Application
	err := r.db.WithContext(ctx).Where("deleted_at IS NULL AND job_id = ?", jobId).Find(&applications).Error
	return applications, err
}

func (r jobPortalRepository) GetJobsByCompanyID(ctx context.Context, companyId int64) ([]*model.Job, error) {
	var jobs []*model.Job
	err := r.db.WithContext(ctx).Where("deleted_at IS NULL AND company_id = ?", companyId).Find(&jobs).Error
	return jobs, err
}

func (r jobPortalRepository) GetJobs(ctx context.Context) ([]*model.Job, error) {
	var jobs []*model.Job
	err := r.db.WithContext(ctx).Where("deleted_at IS NULL").Find(&jobs).Error
	return jobs, err
}

func (r jobPortalRepository) GetCompanies(ctx context.Context) ([]*model.Company, error) {
	var companies []*model.Company
	err := r.db.WithContext(ctx).Where("deleted_at IS NULL").Find(&companies).Error
	return companies, err
}

func (r jobPortalRepository) GetApplicants(ctx context.Context) ([]*model.Applicant, error) {
	var applicants []*model.Applicant
	err := r.db.WithContext(ctx).Where("deleted_at IS NULL").Find(&applicants).Error
	return applicants, err
}

func (r jobPortalRepository) GetApplications(ctx context.Context) ([]*model.Application, error) {
	var applications []*model.Application
	err := r.db.WithContext(ctx).Where("deleted_at IS NULL").Find(&applications).Error
	return applications, err
}

func (r jobPortalRepository) GetUsers(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	err := r.db.WithContext(ctx).Where("deleted_at IS NULL").Find(&users).Error
	return users, err
}

func (r jobPortalRepository) UpdateApplicationByID(ctx context.Context, id int64, updatedApplication model.ApplicationUpdate) (int, error) {
	results := r.db.WithContext(ctx).Model(model.ApplicationUpdate{}).Where("id = ? AND deleted_at IS NULL", id).Updates(updatedApplication)
	return int(results.RowsAffected), results.Error
}

func (r jobPortalRepository) UpdateCompanyByID(ctx context.Context, id int64, updatedCompany model.CompanyUpdate) (int, error) {
	results := r.db.WithContext(ctx).Model(model.CompanyUpdate{}).Where("id = ? AND deleted_at IS NULL", id).Updates(updatedCompany)
	return int(results.RowsAffected), results.Error
}

func (r jobPortalRepository) UpdateJobByID(ctx context.Context, id int64, updatedJob model.JobUpdate) (int, error) {
	results := r.db.WithContext(ctx).Model(model.JobUpdate{}).Where("id = ? AND deleted_at IS NULL", id).Updates(updatedJob)
	return int(results.RowsAffected), results.Error
}

func (r jobPortalRepository) UpdateUserByID(ctx context.Context, id int64, updatedUser model.UserUpdate) (int, error) {
	results := r.db.WithContext(ctx).Model(model.UserUpdate{}).Where("id = ? AND deleted_at IS NULL", id).Updates(updatedUser)
	return int(results.RowsAffected), results.Error
}

func (r jobPortalRepository) DeleteApplicantByID(ctx context.Context, id int64, deletionTime time.Time) (int, error) {
	result := r.db.WithContext(ctx).Model(&model.Applicant{}).Where("id = ? AND deleted_at IS NULL", id).Update("deleted_at", deletionTime)
	return int(result.RowsAffected), result.Error
}

func (r jobPortalRepository) DeleteApplicationByID(ctx context.Context, id int64, deletionTime time.Time) (int, error) {
	result := r.db.WithContext(ctx).Model(&model.Application{}).Where("id = ? AND deleted_at IS NULL", id).Update("deleted_at", deletionTime)
	return int(result.RowsAffected), result.Error
}

func (r jobPortalRepository) DeleteCompanyByID(ctx context.Context, id int64, deletionTime time.Time) (int, error) {
	result := r.db.WithContext(ctx).Model(&model.Company{}).Where("id = ? AND deleted_at IS NULL", id).Update("deleted_at", deletionTime)
	return int(result.RowsAffected), result.Error
}

func (r jobPortalRepository) DeleteJobByID(ctx context.Context, id int64, deletionTime time.Time) (int, error) {
	result := r.db.WithContext(ctx).Model(&model.Job{}).Where("id = ? AND deleted_at IS NULL", id).Update("deleted_at", deletionTime)
	return int(result.RowsAffected), result.Error
}

func (r jobPortalRepository) DeleteUserByID(ctx context.Context, id int64, deletionTime time.Time) (int, error) {
	result := r.db.WithContext(ctx).Model(&model.User{}).Where("id = ? AND deleted_at IS NULL", id).Update("deleted_at", deletionTime)
	return int(result.RowsAffected), result.Error
}
