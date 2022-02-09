package infrastructures

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GetDatabaseConnection -> Retrieve reference for database using package gorm
func GetDatabaseConnection() *gorm.DB {
	db, err := gorm.Open(postgres.Open(mountConnectionString()), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		panic("Unable to connect database")
	}

	return db
}

func mountConnectionString() string {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSslMode := os.Getenv("DB_SSLMODE")

	connectionString := "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s"

	return fmt.Sprintf(connectionString, dbHost, dbUser, dbPassword, dbName, dbPort, dbSslMode)
}
