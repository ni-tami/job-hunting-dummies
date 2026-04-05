package sharedactivities

import (
	"context"
	"fmt"
	"log"

	pb "github.com/ni-tami/job-hunting-dummies-workflows/app/pb/out/go/job_hunting_dummies"
	cModel "github.com/ni-tami/job-hunting-dummies-workflows/app/populate_company/models"
	uModel "github.com/ni-tami/job-hunting-dummies-workflows/app/populate_user/models"
)

type JobHuntServiceActivity struct {
	grpcClient pb.JobHuntServiceClient
}

func NewJobHuntServiceActivity(grpcClient pb.JobHuntServiceClient) *JobHuntServiceActivity {
	return &JobHuntServiceActivity{
		grpcClient: grpcClient,
	}
}

func (a *JobHuntServiceActivity) CreateUserRPCActivity(ctx context.Context, randUser uModel.CreateUser) error {
	// Contact the server and print out its response.
	u, err := a.grpcClient.CreateUser(ctx, &pb.CreateUserRequest{Name: randUser.Name, Username: randUser.Username})
	if err != nil {
		log.Fatalf("could not create user: %v", err)
		return err
	}
	log.Printf("New User: %+v", u)
	return nil
}

// FetchRandomSourceCompanyWithNewAdminUserRPCActivity fetch source company created by new predefined "admin" user
func (a *JobHuntServiceActivity) FetchRandomSourceCompanyWithNewAdminUserRPCActivity(ctx context.Context, size int32) ([]cModel.CreateCompany, error) {
	companiesPB, err := a.grpcClient.GetRandomSourceCompanies(ctx,
		&pb.GetRandomSourceCompaniesRequest{
			Size: size,
		},
	)
	if err != nil {
		log.Fatalf("could not fetch source companies: %v", err)
		return nil, err
	}
	log.Printf("Fetch %d companies:", companiesPB.GetSize())
	var companies []cModel.CreateCompany
	for _, companyPB := range companiesPB.Companies {
		associatedAdminUser := uModel.CreateUser{
			Name:     fmt.Sprintf("%s Admin", companyPB.Name),
			Username: fmt.Sprintf("company-%d-admin", companyPB.SourceId),
		}
		company := cModel.NewCreateCompanyFromRandomSourceCompanyPB(companyPB, associatedAdminUser)
		companies = append(companies, company)
	}
	return companies, nil
}

func (a *JobHuntServiceActivity) CreateCompaniesRPCActivity(ctx context.Context, randCompanies []cModel.CreateCompany) error {
	var createCompaniesRequestPB []*pb.CreateCompanyRequest
	for _, randCompany := range randCompanies {
		createCompaniesRequestPB = append(createCompaniesRequestPB, &pb.CreateCompanyRequest{
			CompanyName: randCompany.CompanyName,
			User: &pb.CreateUserRequest{
				Name:     randCompany.User.Name,
				Username: randCompany.User.Username,
			},
			Description: randCompany.Description,
			Website:     randCompany.Website,
		})
	}
	createCompaniesResponsePB, err := a.grpcClient.CreateCompanies(ctx,
		&pb.CreateCompaniesRequest{
			Companies: createCompaniesRequestPB,
		},
	)
	if err != nil {
		log.Fatalf("could not create companies: %v", err)
		return err
	}
	log.Printf("New Companies Count: %+v", len(createCompaniesResponsePB.Companies))
	return nil
}
