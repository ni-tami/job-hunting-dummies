package model

import (
	"time"
)

type Applicant struct {
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	ID        int64      `json:"id"`
	UpdatedAt time.Time  `json:"updated_at"`
	User      *User      `json:"user"                 gorm:"-;foreignKey:user_id;constraint:OnUpdate:CASCADE,OnDELETE:CASCADE"`
	UserID    int64      `json:"user_id"`
}
