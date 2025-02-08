package db

import (
	"fmt"
	"log"

	"github.com/Malayt04/BookTicket/backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(envConfig *config.EnvConfig, DBMigrator func(db *gorm.DB) error) *gorm.DB {
	uri := fmt.Sprintf(
		`host=%s dbname=%s password=%s sslmode=%s port=5432`,
		envConfig.DB_HOST,
		envConfig.DB_NAME,
		envConfig.DB_PASSWORD,
		envConfig.DB_SSLMODE,
	)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	log.Println("Connected to database")

	return db
}