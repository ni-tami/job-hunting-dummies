package config

import "time"

const (
	PopulateUserTaskQueueName        = "populate-user"
	PopulateUserScheduleIntervalHour = 24 * time.Hour

	PopulateCompanyTaskQueueName        = "populate-company"
	PopulateCompanyScheduleIntervalHour = 24 * time.Hour

	PyPopulateCompanyTaskQueueName = "generate-company-py-queue"
	PyPopulateActivityName         = "generate-company-py-activity"
)
