package model

import (
	"time"

	gql "github.com/ni-tami/job-hunting-dummies-service/internal/graphql/model"
)

type User struct {
	ID        int64      `json:"id"`
	Username  string     `json:"username"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type UserUpdate struct {
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) TableName() string {
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

func NewUserUpdateFromGQL(gqlInput *gql.UpdateUser) UserUpdate {
	return UserUpdate{
		Name:      gqlInput.Name,
		UpdatedAt: *gqlInput.UpdatedAt,
	}
}
