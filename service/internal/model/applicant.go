package model

import "time"

type Applicant struct {
	ID        int64      `json:"id"`
	User      *User      `json:"user"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}
