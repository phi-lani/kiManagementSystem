package main

import (
	"github.com/phi-lani/kimanagementsystem/config"
	"github.com/phi-lani/kimanagementsystem/handlers"
	"github.com/phi-lani/kimanagementsystem/middleware"
	"github.com/phi-lani/kimanagementsystem/models"
	"log"
	"net/http"
)

func main() {
	config.InitDB()

	// Migrate the schema in main to avoid circular dependency
	config.DB.AutoMigrate(&models.User{}, &models.KeyIndividualProfile{}, &models.StartupProfile{}, &models.UserDocument{})

	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)
	http.Handle("/uploadDocument", middleware.TokenValidationMiddleware(http.HandlerFunc(handlers.UploadDocument)))
	http.Handle("/viewProfile", middleware.TokenValidationMiddleware(http.HandlerFunc(handlers.ViewProfile)))
	http.Handle("/updateProfile", middleware.TokenValidationMiddleware(http.HandlerFunc(handlers.UpdateProfile)))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
