package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/phi-lani/kimanagementsystem/config"
	"github.com/phi-lani/kimanagementsystem/models"
)

type OTPVerificationRequest struct {
	Email   string `json:"email"`
	OTPCode string `json:"otpCode"`
}

// VerifyOTP handles OTP verification and checks for expiry
func VerifyOTP(w http.ResponseWriter, r *http.Request) {
	var req OTPVerificationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if time.Now().After(user.OTPExpiry) {
		http.Error(w, "OTP has expired", http.StatusUnauthorized)
		return
	}

	if req.OTPCode != user.OTP {
		http.Error(w, "Invalid OTP code", http.StatusUnauthorized)
		return
	}

	user.MFAEnabled = true
	user.OTP = ""
	user.OTPExpiry = time.Time{}

	if err := config.DB.Save(&user).Error; err != nil {
		http.Error(w, "Failed to verify user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("OTP verified successfully. Your account is now secured with MFA.")
}
