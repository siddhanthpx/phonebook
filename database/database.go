package database

import (
	"fmt"
	"os"

	"github.com/siddhanthpx/phonebook/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	// Get necessary environment variables needed to connect to Postgres
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// Opening Connction to DB
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, name, port)
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	DB = connection

	connection.AutoMigrate(&models.User{})

	if err != nil {
		return err
	}

	return nil
}
