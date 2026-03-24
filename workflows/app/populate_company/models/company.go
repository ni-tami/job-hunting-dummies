package models

import (
	userModel "github.com/ni-tami/job-hunting-dummies-workflows/app/populate_user/models"
)

type CreateCompany struct {
	CompanyName string               `json:"company_name"`
	User        userModel.CreateUser `json:"user"`
	Description string               `json:"description"`
	Website     string               `json:"website"`
}
