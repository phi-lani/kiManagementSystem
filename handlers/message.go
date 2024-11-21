// package handlers

// import (
// 	"encoding/json"
// 	"net/http"
// 	"time"

// 	"github.com/phi-lani/kimanagementsystem/config"
// 	"github.com/phi-lani/kimanagementsystem/models"
// 	"github.com/phi-lani/kimanagementsystem/utils"
// )

// type SendMessageRequest struct {
// 	RecipientID uint   `json:"recipient_id"` // ID of the Key Individual
// 	Subject     string `json:"subject"`
// 	Body        string `json:"body"`
// }

// func SendMessage(w http.ResponseWriter, r *http.Request) {
// 	// Retrieve the sender's user ID from the context (set by TokenValidationMiddleware)
// 	senderID := r.Context().Value("userID").(uint)

// 	var req SendMessageRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		http.Error(w, "Invalid input", http.StatusBadRequest)
// 		return
// 	}

// 	// Check if the recipient exists and is a Key Individual
// 	var recipient models.User
// 	if err := config.DB.Where("id = ? AND role = ?", req.RecipientID, "key_individual").First(&recipient).Error; err != nil {
// 		http.Error(w, "Recipient not found or not a Key Individual", http.StatusNotFound)
// 		return
// 	}

// 	// Create and save the message
// 	message := models.Message{
// 		SenderID:    senderID,
// 		RecipientID: req.RecipientID,
// 		Subject:     req.Subject,
// 		Body:        req.Body,
// 		SentAt:      time.Now(),
// 	}

// 	if err := config.DB.Create(&message).Error; err != nil {
// 		http.Error(w, "Error sending message", http.StatusInternalServerError)
// 		return
// 	}

// 	// Prepare email notification
// 	subject := "New Message from a Startup" + req.Subject
// 	body := "You have received a new message from a Startup.\n\n" + req.Body + "\n\n" +
// 		"Please log in to your account to view and respond to the message."

// 	// Send the email notification
// 	err := utils.SendEmail(recipient.Email, subject, body)
// 	if err != nil {
// 		http.Error(w, "Failed to send email notification", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode("Message sent successfully and notification email sent to the recipient")
// }

// package handlers

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/phi-lani/kimanagementsystem/config"
// 	"github.com/phi-lani/kimanagementsystem/models"
// 	"github.com/phi-lani/kimanagementsystem/utils"
// )

// type SendMessageRequest struct {
// 	RecipientID uint   `json:"recipient_id"`
// 	Subject     string `json:"subject"`
// 	Body        string `json:"body"`
// }

// func SendMessage(w http.ResponseWriter, r *http.Request) {
// 	// Retrieve the sender's user ID from the context (set by TokenValidationMiddleware)
// 	senderID, ok := r.Context().Value("userID").(uint)
// 	if !ok {
// 		log.Println("Unauthorized: Invalid sender ID in context")
// 		http.Error(w, "Unauthorized: Invalid sender ID", http.StatusUnauthorized)
// 		return
// 	}

// 	// Decode the request payload
// 	var req SendMessageRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		log.Printf("Error decoding input: %v", err)
// 		http.Error(w, "Invalid input", http.StatusBadRequest)
// 		return
// 	}

// 	// Validate input
// 	if req.RecipientID == 0 || req.Subject == "" || req.Body == "" {
// 		log.Println("Invalid input: Missing required fields")
// 		http.Error(w, "Invalid input: Missing required fields", http.StatusBadRequest)
// 		return
// 	}

// 	// Check if the recipient exists and is a Key Individual
// 	var recipient models.User
// 	if err := config.DB.Where("id = ? AND role = ?", req.RecipientID, "key_individual").First(&recipient).Error; err != nil {
// 		log.Printf("Recipient not found or not a Key Individual: %v", err)
// 		http.Error(w, "Recipient not found or not a Key Individual", http.StatusNotFound)
// 		return
// 	}

// 	// Log recipient details for debugging
// 	log.Printf("Sending message to recipient ID: %d, email: %s", recipient.ID, recipient.Email)

// 	// Create and save the message in the database
// 	message := models.Message{
// 		SenderID:    senderID,
// 		RecipientID: req.RecipientID,
// 		Subject:     req.Subject,
// 		Body:        req.Body,
// 		SentAt:      time.Now(),
// 	}

// 	if err := config.DB.Create(&message).Error; err != nil {
// 		log.Printf("Error saving message: %v", err)
// 		http.Error(w, "Error sending message", http.StatusInternalServerError)
// 		return
// 	}

// 	// Prepare email notification
// 	subject := "New Message from a Startup: " + req.Subject
// 	body := "You have received a new message from a Startup.\n\n" + req.Body + "\n\n" +
// 		"Please log in to your account to view and respond to the message."

// 	// Send the email notification
// 	if err := utils.SendEmail(recipient.Email, subject, body); err != nil {
// 		log.Printf("Error sending email notification: %v", err)
// 		http.Error(w, "Failed to send email notification", http.StatusInternalServerError)
// 		return
// 	}

// 	// Respond with success
// 	log.Printf("Message sent successfully to recipient ID: %d", req.RecipientID)
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode("Message sent successfully and notification email sent to the recipient")
// }

package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/phi-lani/kimanagementsystem/config"
	"github.com/phi-lani/kimanagementsystem/models"
	"github.com/phi-lani/kimanagementsystem/utils"
)

type SendMessageRequest struct {
	RecipientID uint   `json:"recipient_id"` // ID of the Key Individual
	Subject     string `json:"subject"`
	Body        string `json:"body"`
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	// Retrieve the sender's user ID from the context (set by TokenValidationMiddleware)
	senderID, ok := r.Context().Value("userID").(uint)
	if !ok {
		log.Println("Unauthorized: Invalid sender ID in context")
		http.Error(w, "Unauthorized: Invalid sender ID", http.StatusUnauthorized)
		return
	}

	// Decode the request payload
	var req SendMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding input: %v", err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Validate input
	if req.RecipientID == 0 || req.Subject == "" || req.Body == "" {
		log.Println("Invalid input: Missing required fields")
		http.Error(w, "Invalid input: Missing required fields", http.StatusBadRequest)
		return
	}

	// Check if the recipient exists in the KeyIndividualProfile table
	var keyIndividualProfile models.KeyIndividualProfile
	if err := config.DB.Where("id = ?", req.RecipientID).First(&keyIndividualProfile).Error; err != nil {
		log.Printf("Recipient Key Individual not found: %v", err)
		http.Error(w, "Recipient Key Individual not found", http.StatusNotFound)
		return
	}

	// Fetch the associated user for the Key Individual
	var recipientUser models.User
	if err := config.DB.Where("id = ?", keyIndividualProfile.UserID).First(&recipientUser).Error; err != nil {
		log.Printf("Associated user for Key Individual not found: %v", err)
		http.Error(w, "Associated user for Key Individual not found", http.StatusNotFound)
		return
	}

	// Log recipient details for debugging
	log.Printf("Sending message to Key Individual ID: %d, email: %s", keyIndividualProfile.ID, recipientUser.Email)

	// Create and save the message in the database
	message := models.Message{
		SenderID:    senderID,
		RecipientID: req.RecipientID, // Key Individual Profile ID
		Subject:     req.Subject,
		Body:        req.Body,
		SentAt:      time.Now(),
	}

	if err := config.DB.Create(&message).Error; err != nil {
		log.Printf("Error saving message: %v", err)
		http.Error(w, "Error sending message", http.StatusInternalServerError)
		return
	}

	// Prepare email notification
	subject := "New Message from a Startup: " + req.Subject
	body := "You have received a new message from a Startup.\n\n" + req.Body + "\n\n" +
		"Please log in to your account to view and respond to the message."

	// Send the email notification
	if err := utils.SendEmail(recipientUser.Email, subject, body); err != nil {
		log.Printf("Error sending email notification: %v", err)
		http.Error(w, "Failed to send email notification", http.StatusInternalServerError)
		return
	}

	// Respond with success
	log.Printf("Message sent successfully to Key Individual ID: %d", req.RecipientID)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Message sent successfully and notification email sent to the recipient")
}
