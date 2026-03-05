package model

import "time"

type Application struct {
	ID          int64      `json:"id"`
	ApplicantID int64      `json:"applicant_id"`
	Applicant   *Applicant `json:"applicant"            gorm:"-;foreignKey:applicant_id;constraint:OnUpdate:CASCADE,OnDELETE:CASCADE"`
	JobID       int64      `json:"job_id"`
	Job         *Job       `json:"job"                  gorm:"-;foreignKey:job_id;constraint:OnUpdate:CASCADE,OnDELETE:CASCADE"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}
