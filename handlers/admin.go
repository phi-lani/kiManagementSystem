package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/phi-lani/kimanagementsystem/config"
	"github.com/phi-lani/kimanagementsystem/models"
	"github.com/phi-lani/kimanagementsystem/utils"
)

type VerifyDocumentRequest struct {
	DocumentID uint `json:"document_id"`
	Verified   bool `json:"verified"`
}

// VerifyDocument allows an admin to verify or reject a document
func VerifyDocument(w http.ResponseWriter, r *http.Request) {
	// Ensure the user is an admin
	role, ok := r.Context().Value("role").(string)
	if !ok || role != "admin" {
		http.Error(w, "Forbidden: Admin access only", http.StatusForbidden)
		return
	}

	var req VerifyDocumentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Retrieve the document
	var document models.UserDocument
	if err := config.DB.First(&document, req.DocumentID).Error; err != nil {
		http.Error(w, "Document not found", http.StatusNotFound)
		return
	}

	// Retrieve the user associated with the document
	var user models.User
	if err := config.DB.First(&user, document.UserID).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Update the verification status
	document.Verified = req.Verified
	if err := config.DB.Save(&document).Error; err != nil {
		http.Error(w, "Error updating document status", http.StatusInternalServerError)
		return
	}

	// Send an email notification to the user
	status := "verified"
	if !req.Verified {
		status = "rejected"
	}
	subject := "Document Verification Status Update"
	body := "Dear " + user.Username + ",\n\nYour document has been " + status + ".\n\nThank you."

	if err := utils.SendEmail(user.Email, subject, body); err != nil {
		http.Error(w, "Failed to send email notification", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	if req.Verified {
		w.Write([]byte("Document verified successfully"))
	} else {
		w.Write([]byte("Document rejected successfully"))
	}
}

// DownloadLogs allows admins to download the server logs
func DownloadLogs(w http.ResponseWriter, r *http.Request) {
	// Ensure the user is an admin
	role, ok := r.Context().Value("role").(string)
	if !ok || role != "admin" {
		http.Error(w, "Forbidden: Admin access only", http.StatusForbidden)
		return
	}

	// Path to the log file
	logFilePath := "path/to/your/logfile.log"

	// Open the log file
	file, err := os.Open(logFilePath)
	if err != nil {
		http.Error(w, "Error opening log file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Set headers for file download
	w.Header().Set("Content-Disposition", "attachment; filename=logs.log")
	w.Header().Set("Content-Type", "application/octet-stream")

	// Serve the file
	http.ServeFile(w, r, logFilePath)
}

// ViewUnverifiedDocuments allows admins to view all unverified documents
func ViewUnverifiedDocuments(w http.ResponseWriter, r *http.Request) {
	// Ensure the user is an admin
	role, ok := r.Context().Value("role").(string)
	if !ok || role != "admin" {
		http.Error(w, "Forbidden: Admin access only", http.StatusForbidden)
		return
	}

	// Fetch all unverified documents
	var unverifiedDocuments []models.UserDocument
	if err := config.DB.Where("verified = ?", false).Find(&unverifiedDocuments).Error; err != nil {
		http.Error(w, "Error fetching unverified documents", http.StatusInternalServerError)
		return
	}

	// Respond with the unverified documents
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(unverifiedDocuments)
}
