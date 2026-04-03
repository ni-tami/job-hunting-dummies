package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/ni-tami/job-hunting-dummies-service/internal/config"
	gqlModel "github.com/ni-tami/job-hunting-dummies-service/internal/graphql/model"
	"github.com/ni-tami/job-hunting-dummies-service/internal/model"
	"github.com/ni-tami/job-hunting-dummies-service/internal/repository"
)

type (
	JobPortalUsecase interface {
		// TODO: might need to separate into multiple services if this grows too big
		CreateJob(ctx context.Context, newJob gqlModel.NewJob) (*model.Job, error)
		CreateCompany(ctx context.Context, newCompany model.CompanyCreate) (*model.Company, error)
		CreateCompanies(ctx context.Context, newCompanies []model.CompanyCreate) ([]*model.Company, error)
		CreateApplicant(ctx context.Context, newApplicant gqlModel.NewApplicant) (*model.Applicant, error)
		CreateApplication(ctx context.Context, newApplication gqlModel.NewApplication) (*model.Application, error)
		CreateUser(ctx context.Context, newUser model.UserCreate) (*model.User, error)
		CreateUsers(ctx context.Context, newUsers []model.UserCreate) ([]*model.User, error)

		GetJobByID(ctx context.Context, id int64) (*model.Job, error)
		GetCompanyByID(ctx context.Context, id int64) (*model.Company, error)
		GetApplicantByID(ctx context.Context, id int64) (*model.Applicant, error)
		GetApplicationByID(ctx context.Context, id int64) (*model.Application, error)
		GetUserByID(ctx context.Context, id int64) (*model.User, error)

		GetRandomSourceCompanies(ctx context.Context, size int32) ([]model.SourceCompany, error)

		GetJobsByIds(ctx context.Context, ids []int64) ([]model.Job, []error)
		GetCompaniesByIds(ctx context.Context, ids []int64) ([]model.Company, []error)
		GetApplicantsByIds(ctx context.Context, ids []int64) ([]model.Applicant, []error)
		GetApplicationsByIds(ctx context.Context, ids []int64) ([]model.Application, []error)
		GetUsersByIds(ctx context.Context, ids []int64) ([]model.User, []error)

		GetJobs(ctx context.Context) ([]*model.Job, error)
		GetCompanies(ctx context.Context) ([]*model.Company, error)
		GetApplicants(ctx context.Context) ([]*model.Applicant, error)
		GetApplications(ctx context.Context) ([]*model.Application, error)
		GetUsers(ctx context.Context) ([]*model.User, error)

		UpdateJobByID(ctx context.Context, id int64, updatedJob *gqlModel.UpdateJob) mutationResponse
		UpdateCompanyByID(ctx context.Context, id int64, updatedCompany *gqlModel.UpdateCompany) mutationResponse
		UpdateApplicationByID(ctx context.Context, id int64, updatedApplication *gqlModel.UpdateApplication) mutationResponse
		UpdateUserByID(ctx context.Context, id int64, updatedUser *gqlModel.UpdateUser) mutationResponse

		DeleteJobByID(ctx context.Context, id int64) mutationResponse
		DeleteCompanyByID(ctx context.Context, id int64) mutationResponse
		DeleteApplicantByID(ctx context.Context, id int64) mutationResponse
		DeleteApplicationByID(ctx context.Context, id int64) mutationResponse
		DeleteUserByID(ctx context.Context, id int64) mutationResponse

		GetApplicationsByApplicantID(ctx context.Context, applicantID int64) ([]*model.Application, error)
		GetApplicationsByJobID(ctx context.Context, jobID int64) ([]*model.Application, error)
		GetJobsByCompanyID(ctx context.Context, companyID int64) ([]*model.Job, error)
	}

	jobPortalUsecase struct {
		repo repository.JobPortalRepository
	}

	mutationResponse struct {
		Err       error
		RowCount  *int
		UpdatedAt *time.Time
	}
)

func NewJobPortalUsecase(repo repository.JobPortalRepository) JobPortalUsecase {
	return jobPortalUsecase{
		repo: repo,
	}
}

func (u jobPortalUsecase) CreateJob(ctx context.Context, newJob gqlModel.NewJob) (*model.Job, error) {
	company, err := u.GetCompanyByID(ctx, int64(newJob.CompanyID))
	if err != nil {
		return nil, fmt.Errorf("failed to get company for job: %w", err)
	}
	if company == nil {
		return nil, fmt.Errorf("company not found for job")
	}
	id := time.Now().UnixMicro()
	job := &model.Job{
		ID:           id,
		CompanyID:    company.ID,
		Title:        newJob.Title,
		Description:  newJob.Description,
		Requirements: newJob.Requirements,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	err = u.repo.CreateJob(ctx, job)
	if err != nil {
		return nil, fmt.Errorf("failed to create job: %w", err)
	}
	return job, nil
}

func (u jobPortalUsecase) CreateCompany(ctx context.Context, newCompany model.CompanyCreate) (*model.Company, error) {
	newUser := model.UserCreate{
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
		UserID:      user.ID,
		CompanyName: newCompany.CompanyName,
		Website:     newCompany.Website,
		Description: newCompany.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err = u.repo.CreateCompany(ctx, company)
	if err != nil {
		return nil, fmt.Errorf("failed to create company: %w", err)
	}
	return company, nil
}

func (u jobPortalUsecase) CreateCompanies(ctx context.Context, newCompanies []model.CompanyCreate) ([]*model.Company, error) {
	var (
		newUsers	             []model.UserCreate
		newUsernameToCompanyMap  map[string]*model.Company
		companies                []*model.Company
	)
	newUsers = make([]model.UserCreate, len(newCompanies))
	newUsernameToCompanyMap = make(map[string]*model.Company)
	companies = make([]*model.Company, len(newCompanies))
	for _, newCompany := range newCompanies {
		newUser := model.UserCreate{
			Username: newCompany.User.Username,
			Name:     newCompany.User.Name,
		}
		newUsers = append(newUsers, newUser)
		
		id := time.Now().UnixMicro()
		company := &model.Company{
			ID:          id,
			CompanyName: newCompany.CompanyName,
			Website:     newCompany.Website,
			Description: newCompany.Description,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		newUsernameToCompanyMap[newUser.Username] = company
	}
	users, err := u.CreateUsers(ctx, newUsers)
	for _, user := range users {
		newUsernameToCompanyMap[user.Username].UserID = user.ID
		companies = append(companies, newUsernameToCompanyMap[user.Username])
	}
	if err != nil {
		fmt.Println("Fatal: Failed to create user for company")
	}

	err = u.repo.CreateCompanies(ctx, companies)
	if err != nil {
		return nil, fmt.Errorf("failed to create company: %w", err)
	}
	return companies, nil
}

func (u jobPortalUsecase) CreateApplicant(ctx context.Context, newApplicant gqlModel.NewApplicant) (*model.Applicant, error) {
	newUser := model.UserCreate{
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
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = u.repo.CreateApplicant(ctx, applicant)
	if err != nil {
		return nil, fmt.Errorf("failed to create applicant: %w", err)
	}
	return applicant, nil
}

func (u jobPortalUsecase) CreateApplication(ctx context.Context, newApplication gqlModel.NewApplication) (*model.Application, error) {
	applicant, err := u.GetApplicantByID(ctx, int64(newApplication.ApplicantID))
	if err != nil {
		return nil, fmt.Errorf("failed to get applicant for application")
	}
	if applicant == nil {
		return nil, fmt.Errorf("applicant not found for application")
	}
	job, err := u.GetJobByID(ctx, int64(newApplication.JobID))
	if err != nil {
		return nil, fmt.Errorf("failed to get job for application")
	}
	if job == nil {
		return nil, fmt.Errorf("job not found for application")
	}
	id := time.Now().UnixMicro()
	application := &model.Application{
		ID:          id,
		ApplicantID: applicant.ID,
		JobID:       job.ID,
		Status:      "APPLIED",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err = u.repo.CreateApplication(ctx, application)
	if err != nil {
		return nil, fmt.Errorf("failed to create application: %w", err)
	}
	return application, nil
}

func (u jobPortalUsecase) CreateUser(ctx context.Context, newUser model.UserCreate) (*model.User, error) {
	id := time.Now().UnixMicro()
	user := &model.User{
		ID:        id,
		Username:  newUser.Username,
		Name:      newUser.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := u.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return user, nil
}

func (u jobPortalUsecase) CreateUsers(ctx context.Context, newUsers []model.UserCreate) ([]*model.User, error) {
	var users []*model.User
	for _, newUser := range newUsers {
		id := time.Now().UnixMicro()
		user := &model.User{
			ID:        id,
			Username:  newUser.Username,
			Name:      newUser.Name,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		users = append(users, user)
	}
	err := u.repo.CreateUsers(ctx, users)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return users, nil
}

func (u jobPortalUsecase) GetJobByID(ctx context.Context, id int64) (*model.Job, error) {
	return u.repo.GetJobByID(ctx, id)
}

func (u jobPortalUsecase) GetRandomSourceCompanies(ctx context.Context, size int32) ([]model.SourceCompany, error) {
	if size > config.MaxGetRandomCompaniesSize {
		return nil, fmt.Errorf("decrease number of random companies to get (max=%d)", size)
	}
	ids, err := u.repo.GetRandomSourceCompanyIds(ctx, size)
	if err != nil {
		return nil, fmt.Errorf("failed to get random source company ids: %v. Error: %+v", ids, err)
	}
	sourceCompanies, err := u.repo.GetSourceCompanyByIds(ctx, ids)
	return sourceCompanies, err
}

func (u jobPortalUsecase) GetCompanyByID(ctx context.Context, id int64) (*model.Company, error) {
	return u.repo.GetCompanyByID(ctx, id)
}

func (u jobPortalUsecase) GetApplicantByID(ctx context.Context, id int64) (*model.Applicant, error) {
	return u.repo.GetApplicantByID(ctx, id)
}

func (u jobPortalUsecase) GetApplicationByID(ctx context.Context, id int64) (*model.Application, error) {
	return u.repo.GetApplicationByID(ctx, id)
}

func (u jobPortalUsecase) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
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

func (u jobPortalUsecase) UpdateJobByID(ctx context.Context, id int64, updatedJobInput *gqlModel.UpdateJob) mutationResponse {
	now := time.Now()
	if updatedJobInput.UpdatedAt == nil {
		updatedJobInput.UpdatedAt = &now
	}
	updatedJob := model.NewJobUpdateFromGQL(updatedJobInput)
	rowCount, err := u.repo.UpdateJobByID(ctx, id, updatedJob)
	if err != nil {
		return mutationResponse{
			Err:       err,
			UpdatedAt: nil,
		}
	}
	return mutationResponse{
		Err:       nil,
		RowCount:  &rowCount,
		UpdatedAt: &now,
	}
}

func (u jobPortalUsecase) UpdateCompanyByID(ctx context.Context, id int64, updatedCompanyInput *gqlModel.UpdateCompany) mutationResponse {
	now := time.Now()
	if updatedCompanyInput.UpdatedAt == nil {
		updatedCompanyInput.UpdatedAt = &now
	}
	updatedCompany := model.NewCompanyUpdateFromGQL(updatedCompanyInput)
	rowCount, err := u.repo.UpdateCompanyByID(ctx, id, updatedCompany)
	if err != nil {
		return mutationResponse{
			Err:       err,
			UpdatedAt: nil,
		}
	}
	return mutationResponse{
		Err:       nil,
		RowCount:  &rowCount,
		UpdatedAt: &now,
	}
}

func (u jobPortalUsecase) UpdateApplicationByID(ctx context.Context, id int64, updatedApplicationInput *gqlModel.UpdateApplication) mutationResponse {
	now := time.Now()
	if updatedApplicationInput.UpdatedAt == nil {
		updatedApplicationInput.UpdatedAt = &now
	}
	updatedApplication := model.NewApplicationUpdateFromGQL(updatedApplicationInput)
	rowCount, err := u.repo.UpdateApplicationByID(ctx, id, updatedApplication)
	if err != nil {
		return mutationResponse{
			Err:       err,
			UpdatedAt: nil,
		}
	}
	return mutationResponse{
		Err:       nil,
		RowCount:  &rowCount,
		UpdatedAt: &now,
	}
}

func (u jobPortalUsecase) UpdateUserByID(ctx context.Context, id int64, updatedUserInput *gqlModel.UpdateUser) mutationResponse {
	now := time.Now()
	if updatedUserInput.UpdatedAt == nil {
		updatedUserInput.UpdatedAt = &now
	}
	updatedUser := model.NewUserUpdateFromGQL(updatedUserInput)
	rowCount, err := u.repo.UpdateUserByID(ctx, id, updatedUser)
	if err != nil {
		return mutationResponse{
			Err:       err,
			UpdatedAt: nil,
		}
	}
	return mutationResponse{
		Err:       nil,
		RowCount:  &rowCount,
		UpdatedAt: &now,
	}
}

func (u jobPortalUsecase) DeleteJobByID(ctx context.Context, id int64) mutationResponse {
	now := time.Now()
	rowCount, err := u.repo.DeleteJobByID(ctx, id, now)
	if err != nil {
		return mutationResponse{
			Err: err,
		}
	}
	return mutationResponse{
		Err:       nil,
		RowCount:  &rowCount,
		UpdatedAt: &now,
	}
}

func (u jobPortalUsecase) DeleteCompanyByID(ctx context.Context, id int64) mutationResponse {
	now := time.Now()
	rowCount, err := u.repo.DeleteCompanyByID(ctx, id, now)
	if err != nil {
		return mutationResponse{
			Err: err,
		}
	}
	return mutationResponse{
		Err:       nil,
		RowCount:  &rowCount,
		UpdatedAt: &now,
	}
}

func (u jobPortalUsecase) DeleteApplicantByID(ctx context.Context, id int64) mutationResponse {
	now := time.Now()
	rowCount, err := u.repo.DeleteApplicantByID(ctx, id, now)
	if err != nil {
		return mutationResponse{
			Err: err,
		}
	}
	return mutationResponse{
		Err:       nil,
		RowCount:  &rowCount,
		UpdatedAt: &now,
	}
}

func (u jobPortalUsecase) DeleteApplicationByID(ctx context.Context, id int64) mutationResponse {
	now := time.Now()
	rowCount, err := u.repo.DeleteApplicationByID(ctx, id, now)
	if err != nil {
		return mutationResponse{
			Err: err,
		}
	}
	return mutationResponse{
		Err:       nil,
		RowCount:  &rowCount,
		UpdatedAt: &now,
	}
}

func (u jobPortalUsecase) DeleteUserByID(ctx context.Context, id int64) mutationResponse {
	now := time.Now()
	rowCount, err := u.repo.DeleteUserByID(ctx, id, now)
	if err != nil {
		return mutationResponse{
			Err: err,
		}
	}
	return mutationResponse{
		Err:       nil,
		RowCount:  &rowCount,
		UpdatedAt: &now,
	}
}

func (u jobPortalUsecase) GetApplicationsByApplicantID(ctx context.Context, applicantID int64) ([]*model.Application, error) {
	return u.repo.GetApplicationsByApplicantID(ctx, applicantID)
}

func (u jobPortalUsecase) GetApplicationsByJobID(ctx context.Context, jobID int64) ([]*model.Application, error) {
	return u.repo.GetApplicationsByJobID(ctx, jobID)
}

func (u jobPortalUsecase) GetJobsByCompanyID(ctx context.Context, companyID int64) ([]*model.Job, error) {
	return u.repo.GetJobsByCompanyID(ctx, companyID)
}

func (u jobPortalUsecase) GetJobsByIds(ctx context.Context, ids []int64) ([]model.Job, []error) {
	return u.repo.GetJobsByIds(ctx, ids)
}

func (u jobPortalUsecase) GetCompaniesByIds(ctx context.Context, ids []int64) ([]model.Company, []error) {
	return u.repo.GetCompaniesByIds(ctx, ids)
}

func (u jobPortalUsecase) GetApplicantsByIds(ctx context.Context, ids []int64) ([]model.Applicant, []error) {
	return u.repo.GetApplicantsByIds(ctx, ids)
}

func (u jobPortalUsecase) GetApplicationsByIds(ctx context.Context, ids []int64) ([]model.Application, []error) {
	return u.repo.GetApplicationsByIds(ctx, ids)
}

func (u jobPortalUsecase) GetUsersByIds(ctx context.Context, ids []int64) ([]model.User, []error) {
	return u.repo.GetUsersByIds(ctx, ids)
}
