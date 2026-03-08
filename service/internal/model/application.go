package model

import (
	"time"

	gql "github.com/ni-tami/job-hunting-dummies-service/internal/graphql/model"
)

type Application struct {
	ID          int64      `json:"id"`
	ApplicantID int64      `json:"applicant_id"`
	Applicant   *Applicant `json:"applicant"            gorm:"-;foreignKey:applicant_id;constraint:OnUpdate:CASCADE,OnDELETE:CASCADE"`
	JobID       int64      `json:"job_id"`
	Job         *Job       `json:"job"                  gorm:"-;foreignKey:job_id;constraint:OnUpdate:CASCADE,OnDELETE:CASCADE"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

type ApplicationUpdate struct {
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (a *Application) TableName() string {
	return "applications"
}

func (a *ApplicationUpdate) TableName() string {
	return "applications"
}

func (a *Application) ToGQL() *gql.Application {
	return &gql.Application{
		ID: a.ID,
		Applicant: &gql.Applicant{
			ID: a.ApplicantID,
			// 	User: &gql.User{
			// 		ID:       a.Applicant.User.ID,
			// 		Name:     a.Applicant.User.Name,
			// 		Username: a.Applicant.User.Username,
			// 	},
		},
		Job: &gql.Job{
			ID: a.JobID,
			// 	Title: a.Job.Title,
		},
		Status:    a.Status,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
		DeletedAt: a.DeletedAt,
	}
}

func NewApplicationUpdateFromGQL(gqlInput *gql.UpdateApplication) ApplicationUpdate {
	return ApplicationUpdate{
		Status:    gqlInput.Status,
		UpdatedAt: *gqlInput.UpdatedAt,
	}
}
