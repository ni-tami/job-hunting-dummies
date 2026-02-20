package usecase

import (
	"fmt"
	"context"
	"github.com/ni-tami/job-hunting-dummies-service/internal/model"
	"github.com/ni-tami/job-hunting-dummies-service/internal/repository"
)

type JobPortalUsecase interface {
	CreateJob(ctx context.Context, newJob model.Job) (model.Job, error)
	CreateCompany(ctx context.Context, newCompany model.Company) (model.Company, error)
	CreateApplicant(ctx context.Context, newApplicant model.Applicant) (model.Applicant, error)
	CreateApplication(ctx context.Context, newApplication model.Application) (model.Application, error)
	CreateUser(ctx context.Context, newUser model.User) (model.User, error)

	GetJobByID(ctx context.Context, id int) (model.Job, error)
	GetCompanyByID(ctx context.Context, id int) (model.Company, error)
	GetApplicantByID(ctx context.Context, id int) (model.Applicant, error)
	GetApplicationByID(ctx context.Context, id int) (model.Application, error)
	GetUserByID(ctx context.Context, id int) (model.User, error)

	GetJobs(ctx context.Context) ([]model.Job, error)
	GetCompanies(ctx context.Context) ([]model.Company, error)
	GetApplicants(ctx context.Context) ([]model.Applicant, error)
	GetApplications(ctx context.Context) ([]model.Application, error)
	GetUsers(ctx context.Context) ([]model.User, error)

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

func (u jobPortalUsecase) CreateJob(ctx context.Context, newJob model.Job) (model.Job, error) {
	return u.repo.CreateJob(ctx, newJob)
}

func (u jobPortalUsecase) CreateCompany(ctx context.Context, newCompany model.Company) (model.Company, error) {
	return u.repo.CreateCompany(ctx, newCompany)
}

func (u jobPortalUsecase) CreateApplicant(ctx context.Context, newApplicant model.Applicant) (model.Applicant, error) {
	return u.repo.CreateApplicant(ctx, newApplicant)
}

func (u jobPortalUsecase) CreateApplication(ctx context.Context, newApplication model.Application) (model.Application, error) {
	return u.repo.CreateApplication(ctx, newApplication)
}

func (u jobPortalUsecase) CreateUser(ctx context.Context, newUser model.User) (model.User, error) {
	return u.repo.CreateUser(ctx, newUser)
}

func (u jobPortalUsecase) GetJobByID(ctx context.Context, id int) (model.Job, error) {
	return u.repo.GetJobByID(ctx, id)
}

func (u jobPortalUsecase) GetCompanyByID(ctx context.Context, id int) (model.Company, error) {
	return u.repo.GetCompanyByID(ctx, id)
}

func (u jobPortalUsecase) GetApplicantByID(ctx context.Context, id int) (model.Applicant, error) {
	return u.repo.GetApplicantByID(ctx, id)
}

func (u jobPortalUsecase) GetApplicationByID(ctx context.Context, id int) (model.Application, error) {
	return u.repo.GetApplicationByID(ctx, id)
}

func (u jobPortalUsecase) GetUserByID(ctx context.Context, id int) (model.User, error) {
	return u.repo.GetUserByID(ctx, id)
}

func (u jobPortalUsecase) GetJobs(ctx context.Context) ([]model.Job, error) {
	return u.repo.GetJobs(ctx)
}

func (u jobPortalUsecase) GetCompanies(ctx context.Context) ([]model.Company, error) {
	return u.repo.GetCompanies(ctx)
}

func (u jobPortalUsecase) GetApplicants(ctx context.Context) ([]model.Applicant, error) {
	return u.repo.GetApplicants(ctx)
}

func (u jobPortalUsecase) GetApplications(ctx context.Context) ([]model.Application, error) {
	return u.repo.GetApplications(ctx)
}

func (u jobPortalUsecase) GetUsers(ctx context.Context) ([]model.User, error) {
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
