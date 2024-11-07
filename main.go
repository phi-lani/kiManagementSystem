package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/phi-lani/kimanagementsystem/config"
	"github.com/phi-lani/kimanagementsystem/handlers"
	"github.com/phi-lani/kimanagementsystem/middleware"
	"github.com/phi-lani/kimanagementsystem/models"
)

func main() {
	// Initialize the database connection
	config.InitDB()

	// Migrate the schema for all required models
	err := config.DB.AutoMigrate(
		&models.User{},
		&models.KeyIndividualProfile{},
		&models.StartupProfile{},
		&models.UserDocument{},
		&models.Message{},
	)
	if err != nil {
		// Log a critical error if migration fails
		log.Fatalf("Critical error during migration: %v", err)
	}

	// Create a new router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/register", handlers.Register).Methods("POST")
	router.HandleFunc("/verify-otp", handlers.VerifyOTP).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")

	// Protected routes using middleware for token validation
	// router.Handle("/uploadDocument", middleware.TokenValidationMiddleware(http.HandlerFunc(handlers.UploadDocument))).Methods("POST")
	router.Handle("/viewProfile", middleware.TokenValidationMiddleware(http.HandlerFunc(handlers.ViewProfile))).Methods("GET")
	router.Handle("/updateProfile", middleware.TokenValidationMiddleware(http.HandlerFunc(handlers.UpdateProfile))).Methods("PUT")
	// router.Handle("/sendMessage", middleware.TokenValidationMiddleware(http.HandlerFunc(handlers.SendMessage))).Methods("POST")
	// router.Handle("/viewMessages", middleware.TokenValidationMiddleware(http.HandlerFunc(handlers.ViewMessages))).Methods("GET")

	// Start the server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
