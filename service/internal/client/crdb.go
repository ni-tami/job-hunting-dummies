package client

import (
  "log"
  "fmt"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

func NewCrdbConn() *gorm.DB {
	db, err := gorm.Open(
		postgres.Open(
			"host=localhost port=26257 user=job_hunting_service dbname=job_hunt_service sslmode=disable",
		),
		initConfig(),
	)
	if err != nil {
		log.Fatal("Unable to connect to DB")
	}
	fmt.Println("Connected to CRDB.")
	return db
}

//initConfig Initialize Config
func initConfig() *gorm.Config {
	return &gorm.Config{
		Logger:         initLog(),
		NamingStrategy: initNamingStrategy(),
	}
}

//initLog Connection Log Configuration
func initLog() logger.Interface {
	f, _ := os.Create("gorm.log")
	newLogger := logger.New(log.New(io.MultiWriter(f, os.Stdout), "\r\n", log.LstdFlags), logger.Config{
		Colorful:      true,
		LogLevel:      logger.Info,
		SlowThreshold: time.Second,
	})
	return newLogger
}

//initNamingStrategy Init NamingStrategy
func initNamingStrategy() *schema.NamingStrategy {
	return &schema.NamingStrategy{
		SingularTable: true,
		TablePrefix:   "",
	}
}
