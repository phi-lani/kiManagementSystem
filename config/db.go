package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DB")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	DB = db
	log.Println("Database connected")

	// Run migrations
	//DB.AutoMigrate(&models.User{}, &models.KeyIndividualProfile{}, &models.StartupProfile{}, &models.UserDocument{})
}
