package models

import (
	pb "github.com/ni-tami/job-hunting-dummies-workflows/app/pb/out/go/job_hunting_dummies"
	userModel "github.com/ni-tami/job-hunting-dummies-workflows/app/populate_user/models"
)

type CreateCompany struct {
	CompanyName string               `json:"company_name"`
	User        userModel.CreateUser `json:"user"`
	Description string               `json:"description"`
	Website     string               `json:"website"`
}

func NewCreateCompanyFromRandomSourceCompanyPB(companyPB *pb.SourceCompanyReponse, user userModel.CreateUser) CreateCompany {
	return CreateCompany{
		CompanyName: companyPB.Name,
		Description: companyPB.LongDescription,
		Website: companyPB.Website,
		User: userModel.CreateUser{
			Name: user.Name,
			Username: user.Username,
		},
	}
}
