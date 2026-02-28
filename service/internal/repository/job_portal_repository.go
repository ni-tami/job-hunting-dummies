package repository

import (
	"context"
	"github.com/ni-tami/job-hunting-dummies-service/internal/model"
	"gorm.io/gorm"
	"time"
)

type JobPortalRepository interface {
	CreateApplicant(ctx context.Context, newApplicant model.Applicant) (model.Applicant, error)
	CreateApplication(ctx context.Context, newApplication model.Application) (model.Application, error)
	CreateCompany(ctx context.Context, newCompany model.Company) (model.Company, error)
	CreateJob(ctx context.Context, newJob model.Job) (model.Job, error)
	CreateUser(ctx context.Context, newUser model.User) (model.User, error)

	GetApplicantByID(ctx context.Context, id int) (model.Applicant, error)
	GetApplicationByID(ctx context.Context, id int) (model.Application, error)
	GetCompanyByID(ctx context.Context, id int) (model.Company, error)
	GetJobByID(ctx context.Context, id int) (model.Job, error)
	GetUserByID(ctx context.Context, id int) (model.User, error)

	GetApplicationsByApplicantID(ctx context.Context, applicantId int64) ([]model.Application, error)
	GetApplicationsByJobID(ctx context.Context, jobId int64) ([]model.Application, error)
	GetJobsByCompanyID(ctx context.Context, companyId int64) ([]model.Job, error)

	GetJobs(ctx context.Context) ([]model.Job, error)
	GetCompanies(ctx context.Context) ([]model.Company, error)
	GetApplicants(ctx context.Context) ([]model.Applicant, error)
	GetApplications(ctx context.Context) ([]model.Application, error)
	GetUsers(ctx context.Context) ([]model.User, error)

	UpdateApplicantByID(ctx context.Context, id int, updatedApplicant model.Applicant) (int64, error)
	UpdateApplicationByID(ctx context.Context, id int, updatedApplication model.Application) (int64, error)
	UpdateCompanyByID(ctx context.Context, id int, updatedCompany model.Company) (int64, error)
	UpdateJobByID(ctx context.Context, id int, updatedJob model.Job) (int64, error)
	UpdateUserByID(ctx context.Context, id int, updatedUser model.User) (int64, error)

	DeleteApplicantByID(ctx context.Context, id int) error
	DeleteApplicationByID(ctx context.Context, id int) error
	DeleteCompanyByID(ctx context.Context, id int) error
	DeleteJobByID(ctx context.Context, id int) error
	DeleteUserByID(ctx context.Context, id int) error
}

type jobPortalRepository struct {
	db *gorm.DB
}

func NewJobPortalRepository(db *gorm.DB) JobPortalRepository {
	return jobPortalRepository{
		db: db,
	}
}

func (r jobPortalRepository) CreateApplicant(ctx context.Context, newApplicant model.Applicant) (model.Applicant, error) {
	err := r.db.WithContext(ctx).Save(&newApplicant).Error
	return newApplicant, err
}

func (r jobPortalRepository) CreateApplication(ctx context.Context, newApplication model.Application) (model.Application, error) {
	err := r.db.WithContext(ctx).Save(&newApplication).Error
	return newApplication, err
}

func (r jobPortalRepository) CreateCompany(ctx context.Context, newCompany model.Company) (model.Company, error) {
	err := r.db.WithContext(ctx).Save(&newCompany).Error
	return newCompany, err
}

func (r jobPortalRepository) CreateJob(ctx context.Context, newJob model.Job) (model.Job, error) {
	err := r.db.WithContext(ctx).Save(&newJob).Error
	return newJob, err
}

func (r jobPortalRepository) CreateUser(ctx context.Context, newUser model.User) (model.User, error) {
	err := r.db.WithContext(ctx).Save(&newUser).Error
	return newUser, err
}

func (r jobPortalRepository) GetApplicantByID(ctx context.Context, id int) (model.Applicant, error) {
	var applicant model.Applicant
	err := r.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&applicant).Error
	return applicant, err
}

func (r jobPortalRepository) GetApplicationByID(ctx context.Context, id int) (model.Application, error) {
	var application model.Application
	err := r.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&application).Error
	return application, err
}

func (r jobPortalRepository) GetCompanyByID(ctx context.Context, id int) (model.Company, error) {
	var company model.Company
	err := r.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&company).Error
	return company, err
}

func (r jobPortalRepository) GetJobByID(ctx context.Context, id int) (model.Job, error) {
	var job model.Job
	err := r.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&job).Error
	return job, err
}

func (r jobPortalRepository) GetUserByID(ctx context.Context, id int) (model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&user).Error
	return user, err
}

func (r jobPortalRepository) GetApplicationsByApplicantID(ctx context.Context, applicantId int64) ([]model.Application, error) {
	var applications []model.Application
	err := r.db.WithContext(ctx).Where("deleted_at IS NULL AND applicant_id = ?", applicantId).Find(&applications).Error
	return applications, err
}

func (r jobPortalRepository) GetApplicationsByJobID(ctx context.Context, jobId int64) ([]model.Application, error) {
	var applications []model.Application
	err := r.db.WithContext(ctx).Where("deleted_at IS NULL AND job_id = ?", jobId).Find(&applications).Error
	return applications, err
}

func (r jobPortalRepository) GetJobsByCompanyID(ctx context.Context, companyId int64) ([]model.Job, error) {
	var jobs []model.Job
	err := r.db.WithContext(ctx).Where("deleted_at IS NULL AND company_id = ?", companyId).Find(&jobs).Error
	return jobs, err
}


func (r jobPortalRepository) GetJobs(ctx context.Context) ([]model.Job, error) {
	var jobs []model.Job
	err := r.db.WithContext(ctx).Where("deleted_at IS NULL").Find(&jobs).Error
	return jobs, err
}

func (r jobPortalRepository) GetCompanies(ctx context.Context) ([]model.Company, error) {
	var companies []model.Company
	err := r.db.WithContext(ctx).Where("deleted_at IS NULL").Find(&companies).Error
	return companies, err
}

func (r jobPortalRepository) GetApplicants(ctx context.Context) ([]model.Applicant, error) {
	var applicants []model.Applicant
	err := r.db.WithContext(ctx).Where("deleted_at IS NULL").Find(&applicants).Error
	return applicants, err
}

func (r jobPortalRepository) GetApplications(ctx context.Context) ([]model.Application, error) {
	var applications []model.Application
	err := r.db.WithContext(ctx).Where("deleted_at IS NULL").Find(&applications).Error
	return applications, err
}

func (r jobPortalRepository) GetUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	err := r.db.WithContext(ctx).Where("deleted_at IS NULL").Find(&users).Error
	return users, err
}

func (r jobPortalRepository) UpdateApplicantByID(ctx context.Context, id int, updatedApplicant model.Applicant) (int64, error) {
	results := r.db.WithContext(ctx).Model(&model.Applicant{}).Where("id = ? AND deleted_at IS NULL", id).Updates(&updatedApplicant)
	return results.RowsAffected, results.Error
}

func (r jobPortalRepository) UpdateApplicationByID(ctx context.Context, id int, updatedApplication model.Application) (int64, error) {
	results := r.db.WithContext(ctx).Model(&model.Application{}).Where("id = ? AND deleted_at IS NULL", id).Updates(&updatedApplication)
	return results.RowsAffected, results.Error
}

func (r jobPortalRepository) UpdateCompanyByID(ctx context.Context, id int, updatedCompany model.Company) (int64, error) {
	results := r.db.WithContext(ctx).Model(&model.Company{}).Where("id = ? AND deleted_at IS NULL", id).Updates(&updatedCompany)
	return results.RowsAffected, results.Error
}

func (r jobPortalRepository) UpdateJobByID(ctx context.Context, id int, updatedJob model.Job) (int64, error) {
	results := r.db.WithContext(ctx).Model(&model.Job{}).Where("id = ? AND deleted_at IS NULL", id).Updates(&updatedJob)
	return results.RowsAffected, results.Error
}

func (r jobPortalRepository) UpdateUserByID(ctx context.Context, id int, updatedUser model.User) (int64, error) {
	results := r.db.WithContext(ctx).Model(&model.User{}).Where("id = ? AND deleted_at IS NULL", id).Updates(&updatedUser)
	return results.RowsAffected, results.Error
}

func (r jobPortalRepository) DeleteApplicantByID(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Model(&model.Applicant{}).Where("id = ? AND deleted_at IS NULL", id).Update("deleted_at", time.Now()).Error
}

func (r jobPortalRepository) DeleteApplicationByID(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Model(&model.Application{}).Where("id = ? AND deleted_at IS NULL", id).Update("deleted_at", time.Now()).Error
}

func (r jobPortalRepository) DeleteCompanyByID(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Model(&model.Company{}).Where("id = ? AND deleted_at IS NULL", id).Update("deleted_at", time.Now()).Error
}

func (r jobPortalRepository) DeleteJobByID(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Model(&model.Job{}).Where("id = ? AND deleted_at IS NULL", id).Update("deleted_at", time.Now()).Error
}

func (r jobPortalRepository) DeleteUserByID(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Model(&model.User{}).Where("id = ? AND deleted_at IS NULL", id).Update("deleted_at", time.Now()).Error
}
