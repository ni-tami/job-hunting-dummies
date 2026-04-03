package sharedactivities

import (
	"context"
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


func (a *JobHuntServiceActivity) FetchCompanyRPCActivity(ctx context.Context) (*cModel.CreateCompany, error) {
	return nil, nil
}
	u, err := a.grpcClient.GetCompany(ctx,
		&pb.GetCompanyRequest{
			CompanyName: randCompany.CompanyName,
			User: &pb.GetUserRequest{
				Name:     randCompany.User.Name,
				Username: randCompany.User.Username,
			},
			Description: randCompany.Description,
			Website:     randCompany.Website,
		},
	)
	if err != nil {
		log.Fatalf("could not create user: %v", err)
		return err
	}
// 	log.Printf("New Company: %+v", u)
// 	return u, nil
// }

func (a *JobHuntServiceActivity) CreateCompanyRPCActivity(ctx context.Context, randCompany cModel.CreateCompany) error {
	u, err := a.grpcClient.CreateCompany(ctx,
		&pb.CreateCompanyRequest{
			CompanyName: randCompany.CompanyName,
			User: &pb.CreateUserRequest{
				Name:     randCompany.User.Name,
				Username: randCompany.User.Username,
			},
			Description: randCompany.Description,
			Website:     randCompany.Website,
		},
	)
	if err != nil {
		log.Fatalf("could not create user: %v", err)
		return err
	}
	log.Printf("New Company: %+v", u)
	return nil
}
