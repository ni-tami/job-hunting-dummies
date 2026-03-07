package model

import (
	"time"

	gql "github.com/ni-tami/job-hunting-dummies-service/internal/graphql/model"
)

type Applicant struct {
	ID        int64      `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at"`
	User      *User      `json:"user"                 gorm:"-;foreignKey:user_id;constraint:OnUpdate:CASCADE,OnDELETE:CASCADE"`
	UserID    int64      `json:"user_id"`
}

func (a *Applicant) TableName() string {
	return "applicants"
}

func (a *Applicant) ToGQL() gql.Applicant {
	return gql.Applicant{
		ID: a.ID,
		User: &gql.User{
			ID:       a.User.ID,
			Name:     a.User.Name,
			Username: a.User.Username,
		},
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
		DeletedAt: a.DeletedAt,
	}
}
