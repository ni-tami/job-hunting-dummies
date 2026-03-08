package migration

import (
	"fmt"

	"github.com/ni-tami/job-hunting-dummies-service/internal/client"
	"github.com/ni-tami/job-hunting-dummies-service/internal/model"
)

func MigrateTables() {
	db := client.NewCrdbConn()
	fmt.Println("Migrating tables...")
	err := db.AutoMigrate(model.User{}, model.Applicant{}, model.Company{}, model.Application{}, model.Job{})
	if err != nil {
		fmt.Printf("Failed to migrate tables: %v\n", err)
		return
	}
	fmt.Println("Tables migrated successfully.")
}
