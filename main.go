package main

import (
	"log"
	"net/http"

	"github.com/phi-lani/kimanagementsystem/config"
	"github.com/phi-lani/kimanagementsystem/handlers"

	//"github.com/phi-lani/kimanagementsystem/handlers"
	"github.com/phi-lani/kimanagementsystem/middleware"
	"github.com/phi-lani/kimanagementsystem/models"
)

func main() {
	// Initialize database connection
	config.InitDB()

	// Migrate the schema in main to avoid circular dependency
	// config.DB.AutoMigrate(&models.User{}, &models.KeyIndividualProfile{}, &models.StartupProfile{}, &models.UserDocument{}, &models.Message{})

	err := config.DB.AutoMigrate(
		&models.User{},
		&models.KeyIndividualProfile{},
		&models.StartupProfile{},
		&models.UserDocument{},
		&models.Message{},
	)
	if err != nil {
		// Send alert to monitoring service (e.g., Sentry, Datadog)
		log.Fatalf("Critical error during migration: %v", err)
	}

	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)
	//http.Handle("/uploadDocument", middleware.TokenValidationMiddleware(http.HandlerFunc(handlers.UploadDocument)))
	http.Handle("/viewProfile", middleware.TokenValidationMiddleware(http.HandlerFunc(handlers.ViewProfile)))
	http.Handle("/updateProfile", middleware.TokenValidationMiddleware(http.HandlerFunc(handlers.UpdateProfile)))
	// http.Handle("/sendMessage", middleware.TokenValidationMiddleware(http.HandlerFunc(handlers.SendMessage)))
	// http.Handle("/viewMessages", middleware.TokenValidationMiddleware(http.HandlerFunc(handlers.ViewMessages)))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
