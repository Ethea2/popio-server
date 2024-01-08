package db

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectDatabase() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	sqlDB, err := sql.Open("pgx", os.Getenv("CONNECTION_STRING"))
	Database, err = gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	return nil
}
