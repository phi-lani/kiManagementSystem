package handlers

import (
	"encoding/json"
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
	senderID := r.Context().Value("userID").(uint)

	var req SendMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Check if the recipient exists and is a Key Individual
	var recipient models.User
	if err := config.DB.Where("id = ? AND role = ?", req.RecipientID, "key_individual").First(&recipient).Error; err != nil {
		http.Error(w, "Recipient not found or not a Key Individual", http.StatusNotFound)
		return
	}

	// Create and save the message
	message := models.Message{
		SenderID:    senderID,
		RecipientID: req.RecipientID,
		Subject:     req.Subject,
		Body:        req.Body,
		SentAt:      time.Now(),
	}

	if err := config.DB.Create(&message).Error; err != nil {
		http.Error(w, "Error sending message", http.StatusInternalServerError)
		return
	}

	// Prepare email notification
	subject := "New Message from a Startup" + req.Subject
	body := "You have received a new message from a Startup.\n\n" + req.Body + "\n\n" +
		"Please log in to your account to view and respond to the message."

	// Send the email notification
	err := utils.SendEmail(recipient.Email, subject, body)
	if err != nil {
		http.Error(w, "Failed to send email notification", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Message sent successfully and notification email sent to the recipient")
}
