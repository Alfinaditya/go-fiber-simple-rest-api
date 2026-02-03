package database

import (
	"fmt"
	"github.com/Alfinaditya/go-fiber-simple-rest-api/app/models"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func newPostgresConnection(dsn string) (*gorm.DB, error) {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed connecteding to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed initiate database SQL: %w", err)
	}

	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err = sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed ping database PostgreSQL: %w", err)
	}

	migrate(db)

	return db, nil
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.Author{},
		&models.Book{},
		&models.User{},
	)
}
