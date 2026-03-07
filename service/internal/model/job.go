package model

import (
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

func (j *Job) TableName() string {
	return "jobs"
}

func (j *Job) ToGQL() *gql.Job {
	return &gql.Job{
		ID: j.ID,
		Company: &gql.Company{
			ID: j.CompanyID,
			// 	User: &gql.User{
			// 		ID:       j.Company.User.ID,
			// 		Name:     j.Company.User.Name,
			// 		Username: j.Company.User.Username,
			// 	},
			// 	CompanyName: j.Company.CompanyName,
			// 	Website:     j.Company.Website,
			// 	Description: j.Company.Description,
		},
		Title:        j.Title,
		Description:  j.Description,
		Requirements: j.Requirements,
		CreatedAt:    j.CreatedAt,
		UpdatedAt:    j.UpdatedAt,
		DeletedAt:    j.DeletedAt,
	}
}
