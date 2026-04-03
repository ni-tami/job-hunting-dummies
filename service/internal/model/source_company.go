package model

import (
	pb "github.com/ni-tami/job-hunting-dummies-service/pb/out/go/job_hunting_dummies"
)

type SourceCompany struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	Website         string `json:"website"`
	LongDescription string `json:"long_description"`
}

func (c *SourceCompany) TableName() string {
	return "source_companies"
}

func (c *SourceCompany) ToSourceCompanyGRPC() *pb.SourceCompanyReponse {
	return &pb.SourceCompanyReponse{
		SourceId:        c.ID,
		Name:            c.Name,
		LongDescription: c.LongDescription,
		Website:         c.Website,
	}
}
