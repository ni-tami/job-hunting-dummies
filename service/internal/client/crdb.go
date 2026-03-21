package client

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/ni-tami/job-hunting-dummies-service/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type crdbConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// https://gist.githubusercontent.com/david-yappeter/2be557fc9af25b3fc000962b6f45fa6c/raw/b211147f3e367cbcc03753fc340653ab452e289c/database-1.go
func NewCrdbConn() *gorm.DB {
	config := crdbConfig{
		Host:     config.Koanf.String("crdb.host"),
		Port:     config.Koanf.Int("crdb.port"),
		User:     config.Koanf.String("crdb.user"),
		Password: config.Koanf.String("crdb.password"),
		DBName:   config.Koanf.String("crdb.dbname"),
		SSLMode:  config.Koanf.String("crdb.sslmode"),
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	db, err := gorm.Open(
		postgres.Open(
			dsn,
		),
		initConfig(),
	)
	if err != nil {
		log.Fatal("Unable to connect to DB")
	}
	fmt.Println("Connected to CRDB.")
	return db
}

// initConfig Initialize Config
func initConfig() *gorm.Config {
	return &gorm.Config{
		Logger:         initLog(),
		NamingStrategy: initNamingStrategy(),
	}
}

// initLog Connection Log Configuration
func initLog() logger.Interface {
	f, _ := os.Create("gorm.log")
	newLogger := logger.New(log.New(io.MultiWriter(f, os.Stdout), "\r\n", log.LstdFlags), logger.Config{
		Colorful:      true,
		LogLevel:      logger.Info,
		SlowThreshold: time.Second,
	})
	return newLogger
}

// initNamingStrategy Init NamingStrategy
func initNamingStrategy() *schema.NamingStrategy {
	return &schema.NamingStrategy{
		SingularTable: true,
		TablePrefix:   "",
	}
}
