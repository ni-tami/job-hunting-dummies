package model

import (
	"fmt"
	"time"

	gql "github.com/ni-tami/job-hunting-dummies-service/internal/graphql/model"
)

type Job struct {
	ID           int64      `json:"id"`
	Company      *Company   `json:"company"              gorm:"-;foreignKey:company_id;constraint:OnUpdate:CASCADE,OnDELETE:CASCADE"`
	CompanyID    int64      `json:"company_id"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	Requirements []string   `json:"requirements"         gorm:"serializer:json"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}

type JobUpdate struct {
	Title        *string   `json:"title"`
	Description  *string   `json:"description"`
	Requirements []string  `json:"requirements" gorm:"serializer:json"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (j *Job) TableName() string {
	return "jobs"
}

func (j *JobUpdate) TableName() string {
	return "jobs"
}

func (j *Job) ToGQL() *gql.Job {
	return &gql.Job{
		ID: j.ID,
		Company: &gql.Company{
			ID: j.CompanyID,
		},
		Title:        j.Title,
		Description:  j.Description,
		Requirements: j.Requirements,
		CreatedAt:    j.CreatedAt,
		UpdatedAt:    j.UpdatedAt,
		DeletedAt:    j.DeletedAt,
	}
}

func NewJobUpdateFromGQL(gqlInput *gql.UpdateJob) JobUpdate {
	fmt.Println("in model job: gqlInput:", gqlInput)
	return JobUpdate{
		Title:        gqlInput.Title,
		Description:  gqlInput.Description,
		Requirements: gqlInput.Requirements,
		UpdatedAt:    *gqlInput.UpdatedAt,
	}
}
