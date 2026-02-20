package model

import "time"

type Company struct {
	ID          int64      `json:"id"`
	User        *User      `json:"user"`
	CompanyName string     `json:"companyName"`
	Website     string     `json:"website"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt,omitempty"`
}