package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/phi-lani/kimanagementsystem/config"
	"github.com/phi-lani/kimanagementsystem/models"
)

type VerifyOTPRequest struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

// VerifyOTP verifies the OTP from the database
func VerifyOTP(w http.ResponseWriter, r *http.Request) {
	var req VerifyOTPRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	log.Printf("Received VerifyOTP request: email=%s, otp=%s", req.Email, req.OTP)

	// Retrieve the OTP record from the database
	var otpRecord models.OTP
	if err := config.DB.Where("email = ? AND code = ?", req.Email, req.OTP).First(&otpRecord).Error; err != nil {
		http.Error(w, "Invalid OTP", http.StatusUnauthorized)
		return
	}

	// Check if OTP is expired
	if time.Now().After(otpRecord.ExpiresAt) {
		http.Error(w, "OTP has expired", http.StatusUnauthorized)
		return
	}

	// OTP is valid; perform post-verification actions
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OTP verified successfully"))

	// Optional: Delete the OTP after successful verification
	config.DB.Delete(&otpRecord)
}

// func SendOTP(w http.ResponseWriter, r *http.Request) {
// 	email := r.URL.Query().Get("email")
// 	if email == "" {
// 		http.Error(w, "Email is required", http.StatusBadRequest)
// 		return
// 	}

// 	otpCode := utils.GenerateOTP()
// 	expiresAt := time.Now().Add(60 * time.Minute) // Set OTP to expire in 60 minutes

// 	otpRecord := models.OTP{
// 		Email:     email,
// 		Code:      otpCode,
// 		ExpiresAt: expiresAt,
// 	}

// 	// Delete any old OTP for this email to prevent duplicates
// 	if err := config.DB.Where("email = ?", email).Delete(&models.OTP{}).Error; err != nil {
// 		http.Error(w, "Error deleting old OTP", http.StatusInternalServerError)
// 		return
// 	}

// 	// Store the new OTP
// 	if err := config.DB.Create(&otpRecord).Error; err != nil {
// 		http.Error(w, "Failed to store OTP", http.StatusInternalServerError)
// 		return
// 	}

// 	// Send OTP via email
// 	if err := utils.SendOTPViaEmail(email); err != nil {
// 		http.Error(w, "Failed to send OTP", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("OTP sent successfully"))
// }
