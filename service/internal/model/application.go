package model

import "time"

type Application struct {
	ID        int64      `json:"id"`
	Applicant *Applicant `json:"applicant"`
	Job       *Job       `json:"job"`
	Status    string     `json:"status"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}
