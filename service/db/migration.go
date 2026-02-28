package migration

import (
	"fmt"
	"github.com/ni-tami/job-hunting-dummies-service/internal/model"
	"github.com/ni-tami/job-hunting-dummies-service/internal/client"
)

func MigrateTables() {
	db := client.NewCrdbConn()
	fmt.Println("Migrating tables...")
	db.AutoMigrate(model.User{}, model.Applicant{}, model.Company{}, model.Application{}, model.Job{})
}
