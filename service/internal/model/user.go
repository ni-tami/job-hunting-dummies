package model

import (
	"time"

	gql "github.com/ni-tami/job-hunting-dummies-service/internal/graphql/model"
	pb "github.com/ni-tami/job-hunting-dummies-service/pb/out/go/job_hunting_dummies"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type User struct {
	ID        int64      `json:"id"`
	Username  string     `json:"username"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type UserCreate struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}

type UserUpdate struct {
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *UserUpdate) TableName() string {
	return "users"
}

func (u *User) ToGQL() *gql.User {
	return &gql.User{
		ID:        u.ID,
		Username:  u.Username,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt,
	}
}

func (u *User) ToCreateUserGRPC() *pb.CreateUserResponse {
	return &pb.CreateUserResponse{
		Id:        int32(u.ID),
		Username:  u.Username,
		Name:      u.Name,
		CreatedAt: timestamppb.New(u.CreatedAt),
	}
}

func NewUserCreateFromGQL(gqlInput gql.NewUser) UserCreate {
	return UserCreate{
		Name:     gqlInput.Name,
		Username: gqlInput.Username,
	}
}

func NewUserCreateFromGRPC(grpcInput *pb.CreateUserRequest) UserCreate {
	return UserCreate{
		Name:     grpcInput.Name,
		Username: grpcInput.Username,
	}
}

func NewUserUpdateFromGQL(gqlInput *gql.UpdateUser) UserUpdate {
	return UserUpdate{
		Name:      gqlInput.Name,
		UpdatedAt: *gqlInput.UpdatedAt,
	}
}
