package handlers

import (
	"encoding/json"
	"log"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phi-lani/kimanagementsystem/config"
	"github.com/phi-lani/kimanagementsystem/models"
	"github.com/phi-lani/kimanagementsystem/utils"
)

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
	body := "Dear " + user.Username + ",\n\nYour document has been " + status + ".\n\nThank you."
	if err := utils.SendEmail(user.Email, subject, body); err != nil {
		log.Printf("Error sending email: %v", err)
		http.Error(w, "Failed to send email notification", http.StatusInternalServerError)
		return
	}

	// Use the smart contract instance and authorization from the config package
	contract := config.GetContractInstance()
	auth := config.GetAuth()

	// Convert the user's address to a common.Address type
	kiAddress := common.HexToAddress(user.Address)  // Ensure the user.Address contains a valid Ethereum address
	documentIndex := big.NewInt(int64(document.ID)) // Use the document's ID as the index

	// Call the smart contract's VerifyDocument method
	tx, err := contract.VerifyDocument(auth, kiAddress, documentIndex)
	if err != nil {
		http.Error(w, "Failed to verify document on the blockchain", http.StatusInternalServerError)
		return
	}

	log.Printf("Transaction sent: %s", tx.Hash().Hex())

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	if req.Verified {
		w.Write([]byte("Document verified successfully"))
	} else {
		w.Write([]byte("Document rejected successfully"))
	}
}

// ViewUnverifiedDocuments allows admins to view all unverified documents
func ViewUnverifiedDocuments(w http.ResponseWriter, r *http.Request) {
	log.Println("ViewUnverifiedDocuments endpoint hit")
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
