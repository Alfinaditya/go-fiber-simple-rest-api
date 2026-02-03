package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"
)

var dbInstance *gorm.DB

// ConnectDB initializes the database connection (call once at startup)
func ConnectDB() {
	dbPgsqlHost := os.Getenv("DB_POSTGRES_HOST")
	dbPgsqlPort := os.Getenv("DB_POSTGRES_PORT")
	dbPgsqlUser := os.Getenv("DB_POSTGRES_USER")
	dbPgsqlPassword := os.Getenv("DB_POSTGRES_PASSWORD")
	dbPgsqlName := os.Getenv("DB_POSTGRES_NAME")
	dbPgsqlSchema := os.Getenv("DB_POSTGRES_SCHEMA")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s search_path=%s sslmode=disable TimeZone=UTC",
		dbPgsqlHost, dbPgsqlPort, dbPgsqlUser, dbPgsqlPassword, dbPgsqlName, dbPgsqlSchema,
	)

	db, err := newPostgresConnection(dsn)
	if err != nil {
		log.Fatal(err)
		return
	}

	dbInstance = db
}

func GetDB() *gorm.DB {
	return dbInstance
}
