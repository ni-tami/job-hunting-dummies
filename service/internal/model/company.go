package model

import (
	"time"

	gql "github.com/ni-tami/job-hunting-dummies-service/internal/graphql/model"
)

type Company struct {
	ID          int64      `json:"id"`
	UserID      int64      `json:"user_id"`
	User        *User      `json:"user"                 gorm:"-;foreignKey:user_id;constraint:OnUpdate:CASCADE,OnDELETE:CASCADE"`
	CompanyName string     `json:"company_name"`
	Website     string     `json:"website"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

type CompanyUpdate struct {
	CompanyName *string   `json:"company_name"`
	Website     *string   `json:"website"`
	Description *string   `json:"description"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (c *Company) TableName() string {
	return "companies"
}

func (c *CompanyUpdate) TableName() string {
	return "companies"
}

func (c *Company) ToGQL() *gql.Company {
	return &gql.Company{
		ID: c.ID,
		User: &gql.User{
			ID: c.UserID,
		},
		CompanyName: c.CompanyName,
		Website:     c.Website,
		Description: c.Description,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
		DeletedAt:   c.DeletedAt,
	}
}

func NewCompanyUpdateFromGQL(gqlInput *gql.UpdateCompany) CompanyUpdate {
	return CompanyUpdate{
		CompanyName: gqlInput.CompanyName,
		Website:     gqlInput.Website,
		Description: gqlInput.Description,
		UpdatedAt:   *gqlInput.UpdatedAt,
	}
}
