package grpc_service

import (
	"context"
	"log"

	"github.com/ni-tami/job-hunting-dummies-service/internal/model"
	"github.com/ni-tami/job-hunting-dummies-service/internal/usecase"
	pb "github.com/ni-tami/job-hunting-dummies-service/pb/out/go/job_hunting_dummies"
)

// grpcServer is used to implement JobHuntService.
type grpcServer struct {
	usecase usecase.JobPortalUsecase
	pb.UnimplementedJobHuntServiceServer
}

func NewJobHuntServerImpl(usecase usecase.JobPortalUsecase) pb.JobHuntServiceServer {
	return &grpcServer{
		usecase: usecase,
	}
}

// CreateUser implements JobHuntService
func (s *grpcServer) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	userCreatePayload := model.NewUserCreateFromGRPC(request)
	user, err := s.usecase.CreateUser(ctx, userCreatePayload)
	if err != nil {
		log.Fatalf("Failed to create user. Error %v", err)
	}
	return user.ToCreateUserGRPC(), nil
}

// CreateCompany implements JobHuntService
func (s *grpcServer) CreateCompany(ctx context.Context, request *pb.CreateCompanyRequest) (*pb.CreateCompanyResponse, error) {
	companyCreatePayload := model.NewCompanyCreateFromGRPC(request)
	company, err := s.usecase.CreateCompany(ctx, companyCreatePayload)
	if err != nil {
		log.Fatalf("Failed to create company. Error %v", err)
	}
	return company.ToCreateCompanyGRPC(), nil
}
