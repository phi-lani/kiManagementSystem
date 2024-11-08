package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/phi-lani/kimanagementsystem/config"
	"github.com/phi-lani/kimanagementsystem/models"
	"github.com/phi-lani/kimanagementsystem/utils"
)

// ViewProfile handles retrieving the user's profile
func ViewProfile(w http.ResponseWriter, r *http.Request) {
	// Retrieve the token from the request cookie
	cookie, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "Unauthorized: No token provided", http.StatusUnauthorized)
		return
	}

	// Parse and verify the token
	claims, err := utils.VerifyJWT(cookie.Value)
	if err != nil {
		http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
		return
	}

	// Retrieve the username from the claims
	username := claims.Username

	// Find the user by username
	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Respond with the user profile in JSON format
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// UpdateProfile handles updating the user's profile
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	// Retrieve the token from the request cookie
	cookie, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "Unauthorized: No token provided", http.StatusUnauthorized)
		return
	}

	// Parse and verify the token to extract claims
	claims, err := utils.VerifyJWT(cookie.Value)
	if err != nil {
		http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
		return
	}

	// Use the extracted userID from the claims
	userID := claims.UserID

	// Decode the incoming request body into an update struct
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Find the user in the database
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Check for duplicate username
	if req.Username != "" && req.Username != user.Username {
		var existingUser models.User
		if err := config.DB.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
			http.Error(w, "Username already taken", http.StatusConflict)
			return
		}
		user.Username = req.Username
	}

	// Check for duplicate email
	if req.Email != "" && req.Email != user.Email {
		var existingUser models.User
		if err := config.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
			http.Error(w, "Email already taken", http.StatusConflict)
			return
		}
		user.Email = req.Email
	}

	// Save the updated user information
	if err := config.DB.Save(&user).Error; err != nil {
		log.Printf("error updating profile: %v", err)
		http.Error(w, "Error updating profile", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	json.NewEncoder(w).Encode("Profile updated successfully")
}
