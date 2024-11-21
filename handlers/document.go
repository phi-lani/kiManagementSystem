package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/phi-lani/kimanagementsystem/config"
	"github.com/phi-lani/kimanagementsystem/models"
	"github.com/phi-lani/kimanagementsystem/utils"
)

// type VerifyDocumentRequest struct {
// 	Verified bool `json:"verified"`
// }

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

	// Generate a hash for the document
	documentHash := utils.GenerateHash(fileData)

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
		FileData:     fileData, // Storing file content as binary data
		UploadedAt:   time.Now(),
		Verified:     false,
		Hash:         documentHash, // Store the generated hash
	}
	if err := config.DB.Create(&document).Error; err != nil {
		http.Error(w, "Error saving document information", http.StatusInternalServerError)
		return
	}

	contract := config.GetContractInstance()
	auth := config.GetAuth()

	// Call the smart contract's UploadDocument method
	//documentHash := "DOCUMENT_HASH" // Replace with the actual hash of the uploaded document
	tx, err := contract.UploadDocument(auth, documentHash)
	if err != nil {
		http.Error(w, "Failed to upload document on the blockchain", http.StatusInternalServerError)
		return
	}

	log.Printf("Transaction sent: %s", tx.Hash().Hex())
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File uploaded and stored in the database successfully with hash: " + documentHash))
}

func DownloadDocument(w http.ResponseWriter, r *http.Request) {
	// Retrieve user ID from context (set by TokenValidationMiddleware)
	//userID := r.Context().Value("userID").(uint)

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
	if err := config.DB.Where("id = ?", uint(documentID)).First(&document).Error; err != nil {
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

// AdminOnly ensures that only users with the "admin" role can access the endpoint
func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value("user").(*models.User)
		if !ok || user.Role != "admin" {
			http.Error(w, "Forbidden: Admins only", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
