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

// // Smart contract variables
// var contractAddress = common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3") // Replace with your contract address
// var backend = config.GetBlockchainBackend()                                             // Implement this to get your Ethereum client
// var auth = config.GetAuth()

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

	// Use the smart contract instance and `auth` from config
	contract := config.GetContractInstance()
	auth := config.GetAuth()

	// Call the smart contract's RegisterKI method
	tx, err := contract.RegisterKI(auth, req.FullName, req.Qualifications[0], req.AssetTypes[0], req.Experience[0])
	if err != nil {
		http.Error(w, "Failed to register KI on the blockchain", http.StatusInternalServerError)
		return
	}

	log.Printf("Transaction sent: %s", tx.Hash().Hex())

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
	// Parse the query parameters
	qualifications := r.URL.Query()["qualifications"]
	experience := r.URL.Query()["experience"]
	area := r.URL.Query().Get("area")
	classOfBusiness := r.URL.Query()["class_of_business"]

	log.Printf("Search parameters: Qualifications=%v, Experience=%v, Area=%v, ClassOfBusiness=%v",
		qualifications, experience, area, classOfBusiness)

	// Build the query dynamically
	query := config.DB.Model(&models.KeyIndividualProfile{})

	if len(qualifications) > 0 {
		query = query.Where("qualifications && ?", pq.Array(qualifications))
	}

	if len(experience) > 0 {
		query = query.Where("experience && ?", pq.Array(experience))
	}

	if area != "" {
		query = query.Where("area = ?", area)
	}

	if len(classOfBusiness) > 0 {
		query = query.Where("class_of_business && ?", pq.Array(classOfBusiness))
	}

	// Fetch the results
	var keyIndividuals []models.KeyIndividualProfile
	if err := query.Find(&keyIndividuals).Error; err != nil {
		log.Printf("Error fetching key individuals: %v", err)
		http.Error(w, "Error fetching key individuals", http.StatusInternalServerError)
		return
	}

	// Respond with the search results
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(keyIndividuals)
}

// RegisterAdmin registers a new admin user and sends an OTP for MFA
func RegisterAdmin(w http.ResponseWriter, r *http.Request) {
	log.Printf("inside admin register")
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

	otpRecord := models.OTP{
		Email:     req.Email,
		Code:      otpCode,
		ExpiresAt: otpExpiry,
	}

	if err := config.DB.Create(&otpRecord).Error; err != nil {
		log.Printf("Error saving OTP: %v", err) // Log the error
		http.Error(w, "Failed to generate OTP", http.StatusInternalServerError)
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

func LoginAdmin(w http.ResponseWriter, r *http.Request) {
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Decode the request body into the loginRequest struct
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Retrieve the user from the database using the provided email
	var user models.User
	if err := config.DB.Where("email = ?", loginRequest.Email).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Check if the user has the "admin" role
	if user.Role != "admin" {
		http.Error(w, "Unauthorized: Admin access only", http.StatusForbidden)
		return
	}

	// Verify the provided password with the stored hashed password
	if !utils.CheckPasswordHash(loginRequest.Password, user.Password) {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// Check if MFA is enabled for the admin
	if user.MFAEnabled {
		// Generate an OTP and set its expiry
		otpCode := utils.GenerateOTP()
		user.OTP = otpCode
		user.OTPExpiry = time.Now().Add(5 * time.Minute) // OTP expires in 5 minutes
		config.DB.Save(&user)                            // Save the OTP and expiry to the database

		// Send the OTP to the admin's email
		if err := utils.SendOTPViaEmail(user.Email, otpCode); err != nil {
			http.Error(w, "Failed to send OTP email", http.StatusInternalServerError)
			return
		}

		// Respond to the client that an OTP has been sent
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "OTP sent to your email for MFA verification",
		})
		return
	}

	// If MFA is not required, generate a JWT token and respond
	token, err := utils.GenerateJWT(user.ID, user.Username, user.Role)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Respond with the generated token
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

func UpdateStartupProfile(w http.ResponseWriter, r *http.Request) {
	// Retrieve the user ID from the request context
	userID := r.Context().Value("userID").(uint)

	// Fetch the existing profile from the database
	var profile models.StartupProfile
	if err := config.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		http.Error(w, "Profile not found", http.StatusNotFound)
		return
	}

	// Decode the request payload
	var updatedProfile models.StartupProfile
	if err := json.NewDecoder(r.Body).Decode(&updatedProfile); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Update only the fields that are not empty in the payload
	if updatedProfile.Name != "" {
		profile.Name = updatedProfile.Name
	}
	if updatedProfile.Industry != "" {
		profile.Industry = updatedProfile.Industry
	}
	if updatedProfile.Website != "" {
		profile.Website = updatedProfile.Website
	}
	if updatedProfile.ContactInformation != "" {
		profile.ContactInformation = updatedProfile.ContactInformation
	}
	if updatedProfile.Area != "" {
		profile.Area = updatedProfile.Area
	}

	// Save the updated profile
	if err := config.DB.Save(&profile).Error; err != nil {
		http.Error(w, "Failed to update profile", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Profile updated successfully"))
}

func UpdateKeyIndividualProfile(w http.ResponseWriter, r *http.Request) {
	// Retrieve the user ID from the context
	userID := r.Context().Value("userID").(uint)

	// Fetch the existing profile from the database
	var profile models.KeyIndividualProfile
	if err := config.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		http.Error(w, "Profile not found", http.StatusNotFound)
		return
	}

	// Decode the incoming request body into a temporary struct
	var updatedProfile struct {
		FullName        string   `json:"full_name"`
		Qualifications  []string `json:"qualifications"`
		Experience      []string `json:"experience"`
		ContactDetails  string   `json:"contact_details"`
		Area            string   `json:"area"`
		AssetTypes      []string `json:"asset_types"`
		ClassOfBusiness []string `json:"class_of_business"`
		REExams         []string `json:"re_exams"`
		CPDPoints       *int     `json:"cpd_points"` // Use a pointer to distinguish between 0 and no update
	}
	if err := json.NewDecoder(r.Body).Decode(&updatedProfile); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Update only the fields that are not empty in the payload
	if updatedProfile.FullName != "" {
		profile.FullName = updatedProfile.FullName
	}
	if len(updatedProfile.Qualifications) > 0 {
		profile.Qualifications = updatedProfile.Qualifications
	}
	if len(updatedProfile.Experience) > 0 {
		profile.Experience = updatedProfile.Experience
	}
	if updatedProfile.ContactDetails != "" {
		profile.ContactDetails = updatedProfile.ContactDetails
	}
	if updatedProfile.Area != "" {
		profile.Area = updatedProfile.Area
	}
	if len(updatedProfile.AssetTypes) > 0 {
		profile.AssetTypes = updatedProfile.AssetTypes
	}
	if len(updatedProfile.ClassOfBusiness) > 0 {
		profile.ClassOfBusiness = updatedProfile.ClassOfBusiness
	}
	if len(updatedProfile.REExams) > 0 {
		profile.REExams = updatedProfile.REExams
	}
	if updatedProfile.CPDPoints != nil {
		profile.CPDPoints = *updatedProfile.CPDPoints
	}

	// Save the updated profile
	if err := config.DB.Save(&profile).Error; err != nil {
		http.Error(w, "Failed to update profile", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Profile updated successfully"))
}

// func ViewMessages(w http.ResponseWriter, r *http.Request) {
// 	// Retrieve the user ID from the context (set by middleware)
// 	userID, ok := r.Context().Value("userID").(uint)
// 	if !ok || userID == 0 {
// 		http.Error(w, "Invalid user ID", http.StatusUnauthorized)
// 		return
// 	}

// 	// Fetch the Key Individual Profile associated with the user ID
// 	var keyIndividualProfile models.KeyIndividualProfile
// 	if err := config.DB.Where("user_id = ?", userID).First(&keyIndividualProfile).Error; err != nil {
// 		http.Error(w, "Key Individual Profile not found", http.StatusNotFound)
// 		return
// 	}

// 	// Use the keyIndividualProfile.ID to query the messages table
// 	var messages []models.Message
// 	if err := config.DB.Where("recipient_id = ?", keyIndividualProfile.ID).Find(&messages).Error; err != nil {
// 		http.Error(w, "Failed to fetch messages", http.StatusInternalServerError)
// 		return
// 	}

//		// Respond with the messages in JSON format
//		w.Header().Set("Content-Type", "application/json")
//		json.NewEncoder(w).Encode(messages)
//	}
type MessageResponse struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
	From    string `json:"from"` // Startup name
	SentAt  string `json:"sent_at"`
}

func ViewMessages(w http.ResponseWriter, r *http.Request) {
	// Retrieve the Key Individual's user ID from the context
	keyIndividualUserID, ok := r.Context().Value("userID").(uint)
	if !ok || keyIndividualUserID == 0 {
		http.Error(w, "Invalid user ID", http.StatusUnauthorized)
		return
	}

	// Get the Key Individual's profile ID
	var keyIndividualProfile models.KeyIndividualProfile
	if err := config.DB.Where("user_id = ?", keyIndividualUserID).First(&keyIndividualProfile).Error; err != nil {
		http.Error(w, "Key Individual profile not found", http.StatusNotFound)
		return
	}

	// Query messages and include the startup name
	var messages []MessageResponse
	if err := config.DB.
		Table("messages").
		Select("messages.subject, messages.body, startup_profiles.name AS from, messages.sent_at").
		Joins("JOIN startup_profiles ON messages.sender_id = startup_profiles.user_id").
		Where("messages.recipient_id = ?", keyIndividualProfile.ID).
		Scan(&messages).Error; err != nil {
		log.Printf("Error fetching messages: %v", err)
		http.Error(w, "Error fetching messages", http.StatusInternalServerError)
		return
	}

	// Respond with the messages
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(messages)
}
