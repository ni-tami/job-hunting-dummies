package model

import "time"

type ApplicationStatus int
const (
	Applied ApplicationStatus = iota
	Reviewed
	PendingInterview
	InterviewScheduled
	PendingResult
	Accepted
	Rejected
)

var applicationStatus = map[ApplicationStatus]string{
	Applied: "applied",
	Reviewed: "reviewed",
	PendingInterview: "pending_interview",
	InterviewScheduled: "interview_scheduled",
	PendingResult: "pending_result",
	Accepted: "accepted",
	Rejected: "rejected",
}

func (s ApplicationStatus) String() string {
    return applicationStatus[s]
}

type (
    Users struct {
    id int
    username string
    name string
    created_at time.Time
    updated_at time.Time
    deleted_at time.Time
  }

  Applicants struct {
    id int
    user_id int
    created_at time.Time
    updated_at time.Time
    deleted_at time.Time
  }

  Companies struct {
    id int
    user_id int
    company_name string
    website string
    description string
    created_at time.Time
    updated_at time.Time
    deleted_at time.Time
  }

  Jobs struct {
    id int
    company_id int
    title int
    description string
    requirements []string
    created_at time.Time
    updated_at time.Time
    deleted_at time.Time
  }

  Applications struct {
    id int
    applicant_id int
    job_id int
    status ApplicationStatus
    created_at time.Time
    updated_at time.Time
    deleted_at time.Time
  }
)