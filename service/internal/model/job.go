package model

import "time"

type Job struct {
	ID           int64      `json:"id"`
	Company      *Company   `json:"company"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	Requirements []string   `json:"requirements"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    *time.Time `json:"deletedAt,omitempty"`
}
