package handlers

import (
	"encoding/json"
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
	// Retrieve the user ID from the token or session
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

	var updatedUser models.User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	user.Username = updatedUser.Username
	user.Email = updatedUser.Email

	if err := config.DB.Save(&user).Error; err != nil {
		http.Error(w, "Error updating profile", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Profile updated successfully")
}
