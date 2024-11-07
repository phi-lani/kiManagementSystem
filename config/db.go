package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	//dsn := os.Getenv("DB")
	var err error
	db, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "pgx",
		DSN:        "host=localhost port=5432 user=postgres password=#Pn19970104! dbname=newblockchain_db sslmode=disable",
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	DB = db
	log.Println("Database connected")

	// Run migrations
	// DB.AutoMigrate(&models.User{}, &models.KeyIndividualProfile{}, &models.StartupProfile{}, &models.UserDocument{})
}
