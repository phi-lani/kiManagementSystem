package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/lib/pq"
	"github.com/phi-lani/kimanagementsystem/config"
	"github.com/phi-lani/kimanagementsystem/models"
	"github.com/phi-lani/kimanagementsystem/utils"
)

type RegistrationRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Request struct for registering a key individual
type KeyIndividualRegistrationRequest struct {
	Username        string   `json:"username"`
	Email           string   `json:"email"`
	Password        string   `json:"password"`
	FullName        string   `json:"full_name"`
	Qualifications  []string `json:"qualifications"`
	Experience      []string `json:"experience"`
	ContactDetails  string   `json:"contact_details"`
	Area            string   `json:"area"`
	AssetTypes      []string `json:"asset_types"`
	ClassOfBusiness []string `json:"class_of_business"`
	REExams         []string `json:"re_exams"`
	CPDPoints       int      `json:"cpd_points"`
}

// Request struct for registering a startup
type StartupRegistrationRequest struct {
	Username           string `json:"username"`
	Email              string `json:"email"`
	Password           string `json:"password"`
	Name               string `json:"name"`
	Industry           string `json:"industry"`
	Website            string `json:"website"`
	ContactInformation string `json:"contact_information"`
	Area               string `json:"area"`
}

type SearchKeyIndividualsRequest struct {
	Qualifications  []string `json:"qualifications"`
	Experience      []string `json:"experience"`
	Area            string   `json:"area"`
	ClassOfBusiness []string `json:"class_of_business"` // Updated field
}

// Register handles user registration and sends an OTP for email-based MFA
func Register(w http.ResponseWriter, r *http.Request) {
	var req RegistrationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if strings.EqualFold(req.Username, req.Email) {
		http.Error(w, "Username and email cannot be the same", http.StatusBadRequest)
		return
	}

	// Check if the username or email already exists
	var existingUser models.User
	if err := config.DB.Where("username = ? OR email = ?", req.Username, req.Email).First(&existingUser).Error; err == nil {
		http.Error(w, "Username or email already exists", http.StatusConflict)
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	otpCode := utils.GenerateOTP()
	otpExpiry := time.Now().Add(60 * time.Minute)

	if err := utils.SendOTPViaEmail(req.Email, otpCode); err != nil {
		http.Error(w, "Failed to send OTP email", http.StatusInternalServerError)
		return
	}

	// Create the user in the database
	user := models.User{
		Username:   req.Username,
		Email:      req.Email,
		Password:   hashedPassword,
		Role:       "user",
		MFAEnabled: true,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// Store the OTP in the OTP table
	otpRecord := models.OTP{
		Email:     req.Email,
		Code:      otpCode,
		ExpiresAt: otpExpiry,
	}

	log.Printf("Register stored: email=%s, otp=%s", otpRecord.Email, otpRecord.Code)

	if err := config.DB.Create(&otpRecord).Error; err != nil {
		http.Error(w, "Failed to store OTP", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("User registered successfully. Check your email for the OTP code.")
}

// Login handles user login
func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	fmt.Println("Loggin input username request:" + loginRequest.Username)
	fmt.Println("Loggin input password request:" + loginRequest.Password)

	// Find the user by username
	user, err := models.GetUserByUsername(loginRequest.Username)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	fmt.Println("Stored username:" + user.Username)
	fmt.Println("Stored password:" + user.Password)

	if !utils.CheckPasswordHash(loginRequest.Password, user.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate a JWT token
	token, err := utils.GenerateJWT(user.ID, user.Username, user.Role)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Set the token as a cookie in the response
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour), // Token expiration (can be adjusted)
		HttpOnly: true,                      // Prevents JavaScript access to the cookie
		Secure:   false,                     // Set to true in production for HTTPS
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login successful",
		"token":   token})
}

// RegisterKeyIndividual handles the registration of a key individual
func RegisterKeyIndividual(w http.ResponseWriter, r *http.Request) {
	var req KeyIndividualRegistrationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if strings.EqualFold(req.Username, req.Email) {
		http.Error(w, "Username and email cannot be the same", http.StatusBadRequest)
		return
	}

	// Check if the username or email already exists
	var existingUser models.User
	if err := config.DB.Where("username = ? OR email = ?", req.Username, req.Email).First(&existingUser).Error; err == nil {
		http.Error(w, "Username or email already exists", http.StatusConflict)
		return
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	// Create the user
	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     "key_individual",
	}

	if err := config.DB.Create(&user).Error; err != nil {
		log.Printf("Error creating user: %v", err) // Log the error
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// Generate the OTP and set the expiry time
	otpCode := utils.GenerateOTP()
	otpExpiry := time.Now().Add(10 * time.Minute) // Set OTP expiry to 10 minutes

	// Save the OTP to the database
	otpRecord := models.OTP{
		Email:     user.Email,
		Code:      otpCode,
		ExpiresAt: otpExpiry,
	}

	if err := config.DB.Create(&otpRecord).Error; err != nil {
		log.Printf("Error saving OTP: %v", err) // Log the error
		http.Error(w, "Failed to generate OTP", http.StatusInternalServerError)
		return
	}

	// Send OTP via email
	if err := utils.SendOTPViaEmail(user.Email, otpCode); err != nil {
		log.Printf("Error sending OTP email: %v", err) // Log the error
		http.Error(w, "Failed to send OTP email", http.StatusInternalServerError)
		return
	}

	// Create the Key Individual profile
	keyIndividualProfile := models.KeyIndividualProfile{
		UserID:          user.ID,
		FullName:        req.FullName,
		Qualifications:  req.Qualifications,
		Experience:      req.Experience,
		ContactDetails:  req.ContactDetails,
		Area:            req.Area,
		AssetTypes:      req.AssetTypes,
		ClassOfBusiness: req.ClassOfBusiness,
		REExams:         req.REExams,
		CPDPoints:       req.CPDPoints,
	}

	if err := config.DB.Create(&keyIndividualProfile).Error; err != nil {
		log.Printf("Error creating key individual profile: %v", err) // Log the error
		http.Error(w, "Error creating key individual profile", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Key Individual registered successfully. Check your email for the OTP code.")
}

// RegisterStartup handles the registration of startups and sends an OTP for email verification
func RegisterStartup(w http.ResponseWriter, r *http.Request) {
	var req StartupRegistrationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if strings.EqualFold(req.Username, req.Email) {
		http.Error(w, "Username and email cannot be the same", http.StatusBadRequest)
		return
	}

	// Check if the username or email already exists
	var existingUser models.User
	if err := config.DB.Where("username = ? OR email = ?", req.Username, req.Email).First(&existingUser).Error; err == nil {
		http.Error(w, "Username or email already exists", http.StatusConflict)
		return
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	// Create the user
	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     "startup",
	}

	if err := config.DB.Create(&user).Error; err != nil {
		log.Printf("Error creating user: %v", err)
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// Generate the OTP and set the expiry time
	otpCode := utils.GenerateOTP()
	otpExpiry := time.Now().Add(10 * time.Minute) // Set OTP expiry to 10 minutes

	// Save the OTP to the database
	otpRecord := models.OTP{
		Email:     user.Email,
		Code:      otpCode,
		ExpiresAt: otpExpiry,
	}

	if err := config.DB.Create(&otpRecord).Error; err != nil {
		log.Printf("Error saving OTP: %v", err) // Log the error
		http.Error(w, "Failed to generate OTP", http.StatusInternalServerError)
		return
	}

	// Send OTP via email
	if err := utils.SendOTPViaEmail(user.Email, otpCode); err != nil {
		log.Printf("Error sending OTP email: %v", err) // Log the error
		http.Error(w, "Failed to send OTP email", http.StatusInternalServerError)
		return
	}

	// Create the Startup profile
	startupProfile := models.StartupProfile{
		UserID:             user.ID,
		Name:               req.Name,
		Industry:           req.Industry,
		Website:            req.Website,
		ContactInformation: req.ContactInformation,
		Area:               req.Area,
	}

	if err := config.DB.Create(&startupProfile).Error; err != nil {
		log.Printf("Error creating startup profile: %v", err)
		http.Error(w, "Error creating startup profile", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Startup registered successfully. Check your email for the OTP code.")
}

func SearchKeyIndividuals(w http.ResponseWriter, r *http.Request) {
	var req SearchKeyIndividualsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Start building the query
	query := config.DB.Model(&models.KeyIndividualProfile{})

	// Filter by qualifications if provided
	if len(req.Qualifications) > 0 {
		query = query.Where("qualifications && ?", pq.Array(req.Qualifications)) // Using PostgreSQL array overlap operator
	}

	// Filter by experience if provided
	if len(req.Experience) > 0 {
		query = query.Where("experience && ?", pq.Array(req.Experience)) // Using PostgreSQL array overlap operator
	}

	// Filter by area if provided
	if req.Area != "" {
		query = query.Where("area = ?", req.Area)
	}

	// Filter by class of business if provided
	if len(req.ClassOfBusiness) > 0 {
		query = query.Where("class_of_business && ?", pq.Array(req.ClassOfBusiness)) // Using PostgreSQL array overlap operator
	}

	// Execute the query and fetch results
	var keyIndividuals []models.KeyIndividualProfile
	if err := query.Find(&keyIndividuals).Error; err != nil {
		http.Error(w, "Error fetching key individuals", http.StatusInternalServerError)
		return
	}

	// Respond with the search results
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(keyIndividuals)
}

// LoginAdmin handles admin login with MFA
func LoginAdmin(w http.ResponseWriter, r *http.Request) {
	var req RegistrationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Retrieve the user from the database
	var user models.User
	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Check if the user is an admin
	if user.Role != "admin" {
		http.Error(w, "Unauthorized: Admin access only", http.StatusForbidden)
		return
	}

	// Verify the password
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// Check if MFA is enabled for the admin
	if user.MFAEnabled {
		// Generate an OTP and send it to the admin's email
		otpCode := utils.GenerateOTP()
		user.OTP = otpCode
		user.OTPExpiry = time.Now().Add(5 * time.Minute) // OTP expires in 5 minutes
		config.DB.Save(&user)

		// Send the OTP via email
		utils.SendOTPViaEmail(user.Email, otpCode)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("OTP sent to your email for MFA verification")
		return
	}

	// Generate a JWT token if MFA is not required
	token, err := utils.GenerateJWT(user.ID, user.Username, user.Role)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// RegisterAdmin registers a new admin user and sends an OTP for MFA
func RegisterAdmin(w http.ResponseWriter, r *http.Request) {
	// Check for the API key in the request headers
	apiKey := r.Header.Get("X-API-Key")
	if apiKey != os.Getenv("ADMIN_API_KEY") {
		http.Error(w, "Unauthorized access", http.StatusUnauthorized)
		return
	}

	var req RegistrationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Check if the username or email already exists
	var existingUser models.User
	if err := config.DB.Where("username = ? OR email = ?", req.Username, req.Email).First(&existingUser).Error; err == nil {
		http.Error(w, "Username or email already exists", http.StatusConflict)
		return
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	// Generate an OTP
	otpCode := utils.GenerateOTP()
	otpExpiry := time.Now().Add(5 * time.Minute) // OTP expires in 5 minutes

	// Create the admin user
	adminUser := models.User{
		Username:   req.Username,
		Email:      req.Email,
		Password:   hashedPassword,
		Role:       "admin", // Set the role to "admin"
		MFAEnabled: true,    // Enable MFA for added security
		OTP:        otpCode,
		OTPExpiry:  otpExpiry,
	}

	// Save the admin user to the database
	if err := config.DB.Create(&adminUser).Error; err != nil {
		http.Error(w, "Error creating admin user", http.StatusInternalServerError)
		return
	}

	// Send the OTP via email
	if err := utils.SendOTPViaEmail(adminUser.Email, otpCode); err != nil {
		http.Error(w, "Failed to send OTP email", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Admin registered successfully. Check your email for the OTP code.")
}
