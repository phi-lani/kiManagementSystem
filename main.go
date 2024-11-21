package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/phi-lani/kimanagementsystem/config"
	appHandlers "github.com/phi-lani/kimanagementsystem/handlers"
	"github.com/phi-lani/kimanagementsystem/middleware"
	"github.com/phi-lani/kimanagementsystem/models"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize the database connection
	config.InitDB()

	// Migrate the schema for all required models
	err = config.DB.AutoMigrate(
		&models.User{},
		&models.KeyIndividualProfile{},
		&models.StartupProfile{},
		&models.UserDocument{},
		&models.Message{},
		&models.OTP{},
	)
	if err != nil {
		log.Fatalf("Critical error during migration: %v", err)
	}

	// Smart contract initialization
	const contractAddress = "0x5FbDB2315678afecb367f032d93F642f64180aa3" // Replace with your deployed contract address
	config.InitBlockchain(contractAddress)

	// Create a new router
	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/register/admin", appHandlers.RegisterAdmin).Methods("POST")
	router.HandleFunc("/register/startup", appHandlers.RegisterStartup).Methods("POST")
	router.HandleFunc("/register/keyindividual", appHandlers.RegisterKeyIndividual).Methods("POST")
	router.HandleFunc("/verify-otp", appHandlers.VerifyOTP).Methods("POST")
	router.HandleFunc("/login", appHandlers.Login).Methods("POST")

	// Protected routes using middleware for token validation
	router.Handle("/uploadDocument", middleware.TokenValidationMiddleware(http.HandlerFunc(appHandlers.UploadDocument))).Methods("POST")
	router.Handle("/viewProfile", middleware.TokenValidationMiddleware(http.HandlerFunc(appHandlers.ViewProfile))).Methods("GET")
	router.Handle("/updateProfile", middleware.TokenValidationMiddleware(http.HandlerFunc(appHandlers.UpdateKeyIndividualProfile))).Methods("PUT")
	router.Handle("/downloadDocument", middleware.TokenValidationMiddleware(http.HandlerFunc(appHandlers.DownloadDocument))).Methods("GET")

	// router.Handle("/uploadDocument", http.HandlerFunc(appHandlers.UploadDocument)).Methods("POST")
	// router.Handle("/viewProfile", http.HandlerFunc(appHandlers.ViewProfile)).Methods("GET")
	// router.Handle("/updateProfile", http.HandlerFunc(appHandlers.UpdateKeyIndividualProfile)).Methods("PUT")
	// router.Handle("/downloadDocument", http.HandlerFunc(appHandlers.DownloadDocument)).Methods("GET")

	// Startup-specific routes
	startupRouter := router.PathPrefix("/startup").Subrouter()
	startupRouter.Use(middleware.TokenValidationMiddleware)
	startupRouter.Use(middleware.StartupOnly)
	startupRouter.HandleFunc("/searchKeyIndividuals", appHandlers.SearchKeyIndividuals).Methods("GET")
	startupRouter.HandleFunc("/sendMessage", appHandlers.SendMessage).Methods("POST")

	// Admin-only routes (secured with token validation and role check)
	adminRouter := router.PathPrefix("/admin").Subrouter()
	adminRouter.Use(middleware.TokenValidationMiddleware)
	adminRouter.Use(middleware.AdminOnly)
	adminRouter.HandleFunc("/login", appHandlers.LoginAdmin).Methods("POST")
	startupRouter.HandleFunc("/updateProfile", appHandlers.UpdateStartupProfile).Methods("PUT")
	adminRouter.HandleFunc("/viewUnverifiedDocuments", appHandlers.ViewUnverifiedDocuments).Methods("GET")
	adminRouter.HandleFunc("/verifyDocument", appHandlers.VerifyDocument).Methods("POST")

	// CORS configuration
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000"}) // Replace "*" with specific domains for production
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	// Start the server with CORS enabled
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headersOk, originsOk, methodsOk)(router)))
}
