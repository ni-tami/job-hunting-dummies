package model

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
