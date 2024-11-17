package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/phi-lani/kimanagementsystem/config"
	"github.com/phi-lani/kimanagementsystem/handlers"
	"github.com/phi-lani/kimanagementsystem/middleware"
	"github.com/phi-lani/kimanagementsystem/models"
)

// Deploy the smart contract
// npx hardhat ignition deploy ./ignition/modules/KIManagement.ts --network localhost

// Run the ethereum node
// npx hardhat node

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
		// Log a critical error if migration fails
		log.Fatalf("Critical error during migration: %v", err)
	}

	//====================================================================================
	//Smart contract initialization
	const contractAddress = " 0x5FbDB2315678afecb367f032d93F642f64180aa3" // Replace with your deployed contract address
	config.InitBlockchain(contractAddress)
	// Create a new router
	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/register/admin", handlers.RegisterAdmin).Methods("POST")
	router.HandleFunc("/register/startup", handlers.RegisterStartup).Methods("POST")             // Route for Startup registration
	router.HandleFunc("/register/keyindividual", handlers.RegisterKeyIndividual).Methods("POST") // Route for Key Individual registration
	router.HandleFunc("/verify-otp", handlers.VerifyOTP).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")
	//router.HandleFunc("/send-otp", handlers.SendOTP).Methods("GET") // Route for sending OTP

	// Protected routes using middleware for token validation
	router.Handle("/uploadDocument", middleware.TokenValidationMiddleware(http.HandlerFunc(handlers.UploadDocument))).Methods("POST")
	router.Handle("/viewProfile", middleware.TokenValidationMiddleware(http.HandlerFunc(handlers.ViewProfile))).Methods("GET")
	router.Handle("/updateProfile", middleware.TokenValidationMiddleware(http.HandlerFunc(handlers.UpdateProfile))).Methods("PUT")
	router.Handle("/downloadDocument", middleware.TokenValidationMiddleware(http.HandlerFunc(handlers.DownloadDocument))).Methods("GET")

	// Startup-specific routes
	startupRouter := router.PathPrefix("/startup").Subrouter()
	startupRouter.Use(middleware.TokenValidationMiddleware)
	startupRouter.Use(middleware.StartupOnly)
	startupRouter.HandleFunc("/searchKeyIndividuals", handlers.SearchKeyIndividuals).Methods("GET")
	startupRouter.HandleFunc("/sendMessage", handlers.SendMessage).Methods("POST") // Route to send messages

	// Admin-only routes (secured with token validation and role check)
	adminRouter := router.PathPrefix("/admin").Subrouter()
	adminRouter.Use(middleware.TokenValidationMiddleware) // Ensure the user is authenticated
	adminRouter.Use(middleware.AdminOnly)                 // Ensure the user has an admin role
	adminRouter.HandleFunc("/viewUnverifiedDocuments", handlers.ViewUnverifiedDocuments).Methods("GET")
	adminRouter.HandleFunc("/verifyDocument", handlers.VerifyDocument).Methods("POST")
	//adminRouter.HandleFunc("/downloadLogs", handlers.DownloadLogs).Methods("GET")

	// Start the server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
