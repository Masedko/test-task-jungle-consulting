package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"test-task-jungle-consulting/configuration"
	"test-task-jungle-consulting/internal/core/models"
)

func ConnectDB(config *configuration.EnvConfigModel) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.DBHost,
		config.DBUserName,
		config.DBUserPassword,
		config.DBName,
		config.DBPort,
		config.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database!\n", err.Error())
	}
	//	Extension for postgresql uuid support
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	err = db.AutoMigrate(
		models.Image{},
		models.User{},
	)
	if err != nil {
		log.Fatal("Migration Failed:\n", err.Error())
	}

	log.Println("Successfully connected to the database")

	return db
}
