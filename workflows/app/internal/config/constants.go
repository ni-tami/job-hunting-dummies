package config

import "time"

const (
	PopulateUserTaskQueueName        = "populate-user"
	PopulateUserScheduleIntervalHour = 24 * time.Hour
)
