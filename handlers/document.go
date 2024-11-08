package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/phi-lani/kimanagementsystem/config"
	"github.com/phi-lani/kimanagementsystem/models"
	"github.com/phi-lani/kimanagementsystem/utils"
)

type VerifyDocumentRequest struct {
	Verified bool `json:"verified"`
}

func UploadDocument(w http.ResponseWriter, r *http.Request) {
	// Retrieve user ID from context (set by TokenValidationMiddleware)
	userID := r.Context().Value("userID").(uint)

	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB max upload size
	if err != nil {
		http.Error(w, "File size too large", http.StatusBadRequest)
		return
	}

	// Get the uploaded file
	file, handler, err := r.FormFile("document")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Validate file type (for example, allowing only PDF and image files)
	allowedExtensions := map[string]bool{
		".pdf":  true,
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}
	ext := filepath.Ext(handler.Filename)
	if !allowedExtensions[ext] {
		http.Error(w, "Invalid file type. Only PDF and image files are allowed.", http.StatusBadRequest)
		return
	}

	// Read the file content
	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading the file", http.StatusInternalServerError)
		return
	}

	// Get the document type from the form data
	documentType := r.FormValue("documentType")
	if documentType == "" {
		http.Error(w, "Document type is required", http.StatusBadRequest)
		return
	}

	// Save the file data in the database
	document := models.UserDocument{
		UserID:       userID,
		DocumentType: documentType,
		//FilePath:     "",
		FileData:   fileData, // Storing file content as binary data
		UploadedAt: time.Now(),
		Verified:   false,
	}
	if err := config.DB.Create(&document).Error; err != nil {
		http.Error(w, "Error saving document information", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File uploaded and stored in the database successfully"))
}

func DownloadDocument(w http.ResponseWriter, r *http.Request) {
	// Retrieve user ID from context (set by TokenValidationMiddleware)
	userID := r.Context().Value("userID").(uint)

	// Get the document ID from the URL query parameters
	documentIDStr := r.URL.Query().Get("documentID")
	if documentIDStr == "" {
		http.Error(w, "Document ID is required", http.StatusBadRequest)
		return
	}

	// Convert document ID from string to uint
	documentID, err := strconv.ParseUint(documentIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid Document ID", http.StatusBadRequest)
		return
	}

	// Retrieve the document from the database
	var document models.UserDocument
	if err := config.DB.Where("id = ? AND user_id = ?", uint(documentID), userID).First(&document).Error; err != nil {
		http.Error(w, "Document not found", http.StatusNotFound)
		return
	}

	// Set headers for file download
	w.Header().Set("Content-Disposition", "attachment; filename="+document.DocumentType)
	w.Header().Set("Content-Type", "application/octet-stream")

	// Write the file data to the response
	_, err = w.Write(document.FileData)
	if err != nil {
		http.Error(w, "Error writing file to response", http.StatusInternalServerError)
		return
	}
}

// func VerifyDocument(w http.ResponseWriter, r *http.Request) {
// 	// Check if the user has an admin role
// 	role := r.Context().Value("role").(string)
// 	if role != "admin" {
// 		http.Error(w, "Forbidden: Only admins can verify or reject documents", http.StatusForbidden)
// 		return
// 	}

// 	// Get the document ID from the URL query parameters
// 	documentIDStr := r.URL.Query().Get("documentID")
// 	if documentIDStr == "" {
// 		http.Error(w, "Document ID is required", http.StatusBadRequest)
// 		return
// 	}

// 	// Convert document ID from string to uint
// 	documentID, err := strconv.ParseUint(documentIDStr, 10, 32)
// 	if err != nil {
// 		http.Error(w, "Invalid Document ID", http.StatusBadRequest)
// 		return
// 	}

// 	// Parse the request body to get the verification status
// 	var req VerifyDocumentRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		http.Error(w, "Invalid input", http.StatusBadRequest)
// 		return
// 	}

// 	// Retrieve the document from the database
// 	var document models.UserDocument
// 	if err := config.DB.First(&document, uint(documentID)).Error; err != nil {
// 		http.Error(w, "Document not found", http.StatusNotFound)
// 		return
// 	}

// 	// Update the verification status
// 	document.Verified = req.Verified
// 	if err := config.DB.Save(&document).Error; err != nil {
// 		http.Error(w, "Error updating document status", http.StatusInternalServerError)
// 		return
// 	}

// 	// Respond with a success message
// 	w.WriteHeader(http.StatusOK)
// 	if req.Verified {
// 		w.Write([]byte("Document verified successfully"))
// 	} else {
// 		w.Write([]byte("Document rejected successfully"))
// 	}
// }

func ViewUnverifiedDocuments(w http.ResponseWriter, r *http.Request) {
	// Check if the user has an admin role
	role, ok := r.Context().Value("role").(string)
	if !ok || role != "admin" {
		http.Error(w, "Forbidden: Only admins can view unverified documents", http.StatusForbidden)
		return
	}

	// Retrieve all unverified documents from the database
	var unverifiedDocuments []models.UserDocument
	if err := config.DB.Where("verified = ?", false).Find(&unverifiedDocuments).Error; err != nil {
		http.Error(w, "Error fetching unverified documents", http.StatusInternalServerError)
		return
	}

	// Respond with the list of unverified documents
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(unverifiedDocuments)
}

func VerifyDocument(w http.ResponseWriter, r *http.Request) {
	// Check if the user has an admin role
	role := r.Context().Value("role").(string)
	if role != "admin" {
		http.Error(w, "Forbidden: Only admins can verify or reject documents", http.StatusForbidden)
		return
	}

	// Get the document ID from the URL query parameters
	documentIDStr := r.URL.Query().Get("documentID")
	if documentIDStr == "" {
		http.Error(w, "Document ID is required", http.StatusBadRequest)
		return
	}

	// Convert document ID from string to uint
	documentID, err := strconv.ParseUint(documentIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid Document ID", http.StatusBadRequest)
		return
	}

	// Parse the request body to get the verification status
	var req VerifyDocumentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Retrieve the document from the database
	var document models.UserDocument
	if err := config.DB.First(&document, uint(documentID)).Error; err != nil {
		http.Error(w, "Document not found", http.StatusNotFound)
		return
	}

	// Retrieve the email of the user associated with the document
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

	// Prepare email content
	subject := "Document Verification Status Update"
	status := "verified"
	if !req.Verified {
		status = "rejected"
	}
	body := "Dear " + user.Username + ",\n\nYour document has been " + status + ".\n\nThank you."

	// Send the email notification
	err = utils.SendEmail(user.Email, subject, body)
	if err != nil {
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
