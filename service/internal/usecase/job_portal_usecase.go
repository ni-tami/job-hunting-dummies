package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/ni-tami/job-hunting-dummies-service/internal/graphql/model"
	"github.com/ni-tami/job-hunting-dummies-service/internal/repository"
)

type JobPortalUsecase interface {
	CreateJob(ctx context.Context, newJob model.NewJob) (*model.Job, error)
	CreateCompany(ctx context.Context, newCompany model.NewCompany) (*model.Company, error)
	CreateApplicant(ctx context.Context, newApplicant model.NewApplicant) (*model.Applicant, error)
	CreateApplication(ctx context.Context, newApplication model.NewApplication) (*model.Application, error)
	CreateUser(ctx context.Context, newUser model.NewUser) (*model.User, error)

	GetJobByID(ctx context.Context, id int) (*model.Job, error)
	GetCompanyByID(ctx context.Context, id int) (*model.Company, error)
	GetApplicantByID(ctx context.Context, id int) (*model.Applicant, error)
	GetApplicationByID(ctx context.Context, id int) (*model.Application, error)
	GetUserByID(ctx context.Context, id int) (*model.User, error)

	GetJobs(ctx context.Context) ([]*model.Job, error)
	GetCompanies(ctx context.Context) ([]*model.Company, error)
	GetApplicants(ctx context.Context) ([]*model.Applicant, error)
	GetApplications(ctx context.Context) ([]*model.Application, error)
	GetUsers(ctx context.Context) ([]*model.User, error)

	UpdateJobByID(ctx context.Context, id int, updatedJob model.Job) error
	UpdateCompanyByID(ctx context.Context, id int, updatedCompany model.Company) error
	UpdateApplicantByID(ctx context.Context, id int, updatedApplicant model.Applicant) error
	UpdateApplicationByID(ctx context.Context, id int, updatedApplication model.Application) error
	UpdateUserByID(ctx context.Context, id int, updatedUser model.User) error

	DeleteJobByID(ctx context.Context, id int) error
	DeleteCompanyByID(ctx context.Context, id int) error
	DeleteApplicantByID(ctx context.Context, id int) error
	DeleteApplicationByID(ctx context.Context, id int) error
	DeleteUserByID(ctx context.Context, id int) error
}

type jobPortalUsecase struct {
	repo repository.JobPortalRepository
}

func NewJobPortalUsecase(repo repository.JobPortalRepository) JobPortalUsecase {
	return jobPortalUsecase{
		repo: repo,
	}
}

func (u jobPortalUsecase) CreateJob(ctx context.Context, newJob model.NewJob) (*model.Job, error) {
	company, err := u.GetCompanyByID(ctx, int(newJob.CompanyID))
	if err != nil {
		return nil, fmt.Errorf("failed to get company for job: %w", err)
	}
	if company == nil {
		return nil, fmt.Errorf("company not found for job")
	}
	id := time.Now().UnixMicro()
	job := &model.Job{
		ID:           id,
		Company:      company,
		Title:        newJob.Title,
		Description:  newJob.Description,
		Requirements: newJob.Requirements,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DeletedAt:    nil,
	}
	err = u.repo.CreateJob(ctx, job)
	if err != nil {
		return nil, fmt.Errorf("failed to create job: %w", err)
	}
	return job, nil
}

func (u jobPortalUsecase) CreateCompany(ctx context.Context, newCompany model.NewCompany) (*model.Company, error) {
	newUser := model.NewUser{
		Username: newCompany.User.Username,
		Name:     newCompany.User.Name,
	}
	user, err := u.CreateUser(ctx, newUser)
	if err != nil {
		fmt.Println("Fatal: Failed to create user for company")
	}
	id := time.Now().UnixMicro()
	company := &model.Company{
		ID:          id,
		User:        user,
		CompanyName: newCompany.CompanyName,
		Website:     newCompany.Website,
		Description: newCompany.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
	}
	err = u.repo.CreateCompany(ctx, company)
	if err != nil {
		return nil, fmt.Errorf("failed to create company: %w", err)
	}
	return company, nil
}

func (u jobPortalUsecase) CreateApplicant(ctx context.Context, newApplicant model.NewApplicant) (*model.Applicant, error) {
	newUser := model.NewUser{
		Username: newApplicant.User.Username,
		Name:     newApplicant.User.Name,
	}
	user, err := u.CreateUser(ctx, newUser)
	if err != nil {
		fmt.Println("Fatal: Failed to create user for applicant")
	}
	id := time.Now().UnixMicro()
	applicant := &model.Applicant{
		ID:        id,
		User:      user,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}
	err = u.repo.CreateApplicant(ctx, applicant)
	if err != nil {
		return nil, fmt.Errorf("failed to create applicant: %w", err)
	}
	return applicant, nil
}

func (u jobPortalUsecase) CreateApplication(ctx context.Context, newApplication model.NewApplication) (*model.Application, error) {
	applicant, err := u.GetApplicantByID(ctx, int(newApplication.ApplicantID))
	if err != nil {
		return nil, fmt.Errorf("failed to get applicant for application")
	}
	if applicant == nil {
		return nil, fmt.Errorf("applicant not found for application")
	}
	job, err := u.GetJobByID(ctx, int(newApplication.JobID))
	if err != nil {
		return nil, fmt.Errorf("failed to get job for application")
	}
	if job == nil {
		return nil, fmt.Errorf("job not found for application")
	}
	id := time.Now().UnixMicro()
	application := &model.Application{
		ID:        id,
		Applicant: applicant,
		Job:       job,
		Status:    "APPLIED",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}
	err = u.repo.CreateApplication(ctx, application)
	if err != nil {
		return nil, fmt.Errorf("failed to create application: %w", err)
	}
	return application, nil
}

func (u jobPortalUsecase) CreateUser(ctx context.Context, newUser model.NewUser) (*model.User, error) {
	id := time.Now().UnixMicro()
	user := &model.User{
		ID:        id,
		Username:  newUser.Username,
		Name:      newUser.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}
	err := u.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return user, nil
}

func (u jobPortalUsecase) GetJobByID(ctx context.Context, id int) (*model.Job, error) {
	return u.repo.GetJobByID(ctx, id)
}

func (u jobPortalUsecase) GetCompanyByID(ctx context.Context, id int) (*model.Company, error) {
	return u.repo.GetCompanyByID(ctx, id)
}

func (u jobPortalUsecase) GetApplicantByID(ctx context.Context, id int) (*model.Applicant, error) {
	return u.repo.GetApplicantByID(ctx, id)
}

func (u jobPortalUsecase) GetApplicationByID(ctx context.Context, id int) (*model.Application, error) {
	return u.repo.GetApplicationByID(ctx, id)
}

func (u jobPortalUsecase) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	return u.repo.GetUserByID(ctx, id)
}

func (u jobPortalUsecase) GetJobs(ctx context.Context) ([]*model.Job, error) {
	return u.repo.GetJobs(ctx)
}

func (u jobPortalUsecase) GetCompanies(ctx context.Context) ([]*model.Company, error) {
	return u.repo.GetCompanies(ctx)
}

func (u jobPortalUsecase) GetApplicants(ctx context.Context) ([]*model.Applicant, error) {
	return u.repo.GetApplicants(ctx)
}

func (u jobPortalUsecase) GetApplications(ctx context.Context) ([]*model.Application, error) {
	return u.repo.GetApplications(ctx)
}

func (u jobPortalUsecase) GetUsers(ctx context.Context) ([]*model.User, error) {
	return u.repo.GetUsers(ctx)
}

func (u jobPortalUsecase) UpdateJobByID(ctx context.Context, id int, updatedJob model.Job) error {
	rowCount, err := u.repo.UpdateJobByID(ctx, id, updatedJob)
	if err == nil {
		fmt.Printf("Updated row count: %d", rowCount)
	}
	return err
}

func (u jobPortalUsecase) UpdateCompanyByID(ctx context.Context, id int, updatedCompany model.Company) error {
	rowCount, err := u.repo.UpdateCompanyByID(ctx, id, updatedCompany)
	if err == nil {
		fmt.Printf("Updated row count: %d", rowCount)
	}
	return err
}

func (u jobPortalUsecase) UpdateApplicantByID(ctx context.Context, id int, updatedApplicant model.Applicant) error {
	rowCount, err := u.repo.UpdateApplicantByID(ctx, id, updatedApplicant)
	if err == nil {
		fmt.Printf("Updated row count: %d", rowCount)
	}
	return err
}

func (u jobPortalUsecase) UpdateApplicationByID(ctx context.Context, id int, updatedApplication model.Application) error {
	rowCount, err := u.repo.UpdateApplicationByID(ctx, id, updatedApplication)
	if err == nil {
		fmt.Printf("Updated row count: %d", rowCount)
	}
	return err
}

func (u jobPortalUsecase) UpdateUserByID(ctx context.Context, id int, updatedUser model.User) error {
	rowCount, err := u.repo.UpdateUserByID(ctx, id, updatedUser)
	if err == nil {
		fmt.Printf("Updated row count: %d", rowCount)
	}
	return err
}

func (u jobPortalUsecase) DeleteJobByID(ctx context.Context, id int) error {
	return u.repo.DeleteJobByID(ctx, id)
}

func (u jobPortalUsecase) DeleteCompanyByID(ctx context.Context, id int) error {
	return u.repo.DeleteCompanyByID(ctx, id)
}

func (u jobPortalUsecase) DeleteApplicantByID(ctx context.Context, id int) error {
	return u.repo.DeleteApplicantByID(ctx, id)
}

func (u jobPortalUsecase) DeleteApplicationByID(ctx context.Context, id int) error {
	return u.repo.DeleteApplicationByID(ctx, id)
}

func (u jobPortalUsecase) DeleteUserByID(ctx context.Context, id int) error {
	return u.repo.DeleteUserByID(ctx, id)
}
