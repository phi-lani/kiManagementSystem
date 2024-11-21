package handlers

import (
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phi-lani/kimanagementsystem/config"
	"github.com/phi-lani/kimanagementsystem/models"
	"github.com/phi-lani/kimanagementsystem/utils"
)

func AdminDashboard(w http.ResponseWriter, r *http.Request) {
	// Example logic for dashboard summary
	var stats struct {
		TotalUsers          int64 `json:"total_users"`
		TotalKeyIndividuals int64 `json:"total_key_individuals"`
		TotalStartups       int64 `json:"total_startups"`
		PendingDocuments    int64 `json:"pending_documents"`
	}

	// Query for stats
	config.DB.Model(&models.User{}).Count(&stats.TotalUsers)
	config.DB.Model(&models.KeyIndividualProfile{}).Count(&stats.TotalKeyIndividuals)
	config.DB.Model(&models.StartupProfile{}).Count(&stats.TotalStartups)
	config.DB.Model(&models.UserDocument{}).Where("verified = ?", false).Count(&stats.PendingDocuments)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func ViewUnverifiedDocuments(w http.ResponseWriter, r *http.Request) {
	type DocumentWithUser struct {
		ID           uint      `json:"id"`
		UserID       uint      `json:"user_id"`
		Username     string    `json:"username"`
		DocumentType string    `json:"document_type"`
		UploadedAt   time.Time `json:"uploaded_at"`
		Verified     bool      `json:"verified"`
	}

	var documents []DocumentWithUser

	// Perform a join between UserDocuments and Users
	err := config.DB.Table("user_documents").
		Select("user_documents.id, user_documents.user_id, users.username, user_documents.document_type, user_documents.uploaded_at, user_documents.verified").
		Joins("JOIN users ON users.id = user_documents.user_id").
		Where("user_documents.verified = ?", false).
		Scan(&documents).Error

	if err != nil {
		http.Error(w, "Failed to fetch unverified documents", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(documents)
}

func GetMessageLogs(w http.ResponseWriter, r *http.Request) {
	var messages []models.Message
	if err := config.DB.Find(&messages).Error; err != nil {
		http.Error(w, "Failed to fetch messages", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

type VerifyDocumentRequest struct {
	DocumentID uint `json:"document_id"`
	Verified   bool `json:"verified"`
}

func VerifyDocument(w http.ResponseWriter, r *http.Request) {
	// Ensure the user is an admin
	role, ok := r.Context().Value("role").(string)
	if !ok || role != "admin" {
		log.Printf("Role check failed: role=%v", role)
		http.Error(w, "Forbidden: Admin access only", http.StatusForbidden)
		return
	}

	var req VerifyDocumentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding input: %v", err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Retrieve the document
	var document models.UserDocument
	if err := config.DB.First(&document, req.DocumentID).Error; err != nil {
		log.Printf("Error finding document: %v", err)
		http.Error(w, "Document not found", http.StatusNotFound)
		return
	}

	// Retrieve the user associated with the document
	var user models.User
	if err := config.DB.First(&user, document.UserID).Error; err != nil {
		log.Printf("Error finding user: %v", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Update the verification status
	document.Verified = req.Verified
	if err := config.DB.Save(&document).Error; err != nil {
		log.Printf("Error updating document: %v", err)
		http.Error(w, "Error updating document status", http.StatusInternalServerError)
		return
	}

	// Send an email notification to the user
	status := "verified"
	if !req.Verified {
		status = "rejected"
	}

	subject := "Document Verification Status Update"
	body := "Dear " + user.Username + ",\n\n" +
		"Your document has been " + status + ".\n\n" +
		"If you have any questions, please contact us.\n\nThank you."
	if err := utils.SendEmail(user.Email, subject, body); err != nil {
		log.Printf("Error sending email: %v", err)
		http.Error(w, "Failed to send email notification", http.StatusInternalServerError)
		return
	}

	// Log the email notification
	log.Printf("Email sent to user: %s with status: %s", user.Email, status)

	// Use the smart contract instance and authorization from the config package
	contract := config.GetContractInstance()
	auth := config.GetAuth()

	// Convert the user's address to a common.Address type
	kiAddress := common.HexToAddress(user.Address)  // Ensure the user.Address contains a valid Ethereum address
	documentIndex := big.NewInt(int64(document.ID)) // Use the document's ID as the index

	// Call the smart contract's VerifyDocument method
	tx, err := contract.VerifyDocument(auth, kiAddress, documentIndex)
	if err != nil {
		log.Printf("Blockchain verification error: %v", err)
		http.Error(w, "Failed to verify document on the blockchain", http.StatusInternalServerError)
		return
	}

	log.Printf("Transaction sent: %s", tx.Hash().Hex())

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	if req.Verified {
		w.Write([]byte("Document verified successfully. Email notification sent."))
	} else {
		w.Write([]byte("Document rejected successfully. Email notification sent."))
	}
}
