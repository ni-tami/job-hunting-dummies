package model

import (
	"time"

	gql "github.com/ni-tami/job-hunting-dummies-service/internal/graphql/model"
	pb "github.com/ni-tami/job-hunting-dummies-service/pb/out/go/job_hunting_dummies"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type CompanyCreate struct {
	User        UserCreate `json:"user"`
	CompanyName string     `json:"company_name"`
	Website     string     `json:"website"`
	Description string     `json:"description"`
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

func (c *Company) ToCreateCompanyGRPC() *pb.CreateCompanyResponse {
	return &pb.CreateCompanyResponse{
		Id:          c.ID,
		UserId:      c.UserID,
		CompanyName: c.CompanyName,
		Description: c.Description,
		Website:     c.Website,
		CreatedAt:   timestamppb.New(c.CreatedAt),
	}
}

func NewCompanyCreateFromGQL(gqlInput gql.NewCompany) CompanyCreate {
	return CompanyCreate{
		User: UserCreate{
			Name:     gqlInput.User.Name,
			Username: gqlInput.User.Username,
		},
		CompanyName: gqlInput.CompanyName,
		Website:     gqlInput.Website,
		Description: gqlInput.Description,
	}
}

func NewCompanyCreateFromGRPC(grpcInput *pb.CreateCompanyRequest) CompanyCreate {
	return CompanyCreate{
		User: UserCreate{
			Name:     grpcInput.User.Name,
			Username: grpcInput.User.Username,
		},
		CompanyName: grpcInput.CompanyName,
		Website:     grpcInput.Website,
		Description: grpcInput.Description,
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
