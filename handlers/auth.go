package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/phi-lani/kimanagementsystem/config"
	"github.com/phi-lani/kimanagementsystem/models"
	"github.com/phi-lani/kimanagementsystem/utils"
)

type RegistrationRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
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
	otpExpiry := time.Now().Add(10 * time.Minute)

	if err := utils.SendOTPViaEmail(req.Email, otpCode); err != nil {
		http.Error(w, "Failed to send OTP email", http.StatusInternalServerError)
		return
	}

	user := models.User{
		Username:   req.Username,
		Email:      req.Email,
		Password:   hashedPassword,
		Role:       "user",
		MFAEnabled: true,
		OTP:        otpCode,
		OTPExpiry:  otpExpiry,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
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
	token, err := utils.GenerateJWT(user.ID, user.Username)
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

// Helper function to check if username or email already exists
func userExists(username, email string) bool {
	var user models.User
	if err := config.DB.Where("username = ? OR email = ?", username, email).First(&user).Error; err == nil {
		return true // User with username or email already exists
	}
	return false
}
