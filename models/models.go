package models

import (
	"time"

	"github.com/phi-lani/kimanagementsystem/config"
)

// User model represents the main user table
type User struct {
	ID         uint      `gorm:"primaryKey"`
	Username   string    `gorm:"unique;not null"`
	Email      string    `gorm:"unique;not null"`
	Password   string    `gorm:"not null"`
	Role       string    `gorm:"type:varchar(20);not null"` // Role of the user (e.g., "user", "admin")
	MFAEnabled bool      `gorm:"default:false"`             // Indicates if MFA is enabled
	MFASecret  string    `gorm:"type:varchar(255)"`         // Secret key for MFA (if using an authenticator app)
	OTP        string    `gorm:"type:varchar(6)"`           // OTP for email-based MFA
	OTPExpiry  time.Time `gorm:""`                          // Expiry time for the OTP
	CreatedAt  time.Time `gorm:"autoCreateTime"`            // Timestamp for when the user was created
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`            // Timestamp for when the user was last updated
}

// CreateUser creates a new user in the database
func (u *User) CreateUser() error {
	return config.DB.Create(&u).Error
}

// GetUserByUsername fetches a user by their username
func GetUserByUsername(username string) (*User, error) {
	var user User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// KeyIndividualProfile model represents additional details for Key Individuals
type KeyIndividualProfile struct {
	UserID              uint `gorm:"primaryKey"`
	FullName            string
	ContactNumber       string
	Location            string
	LinkedInProfile     string
	Qualification       string
	ClassOfBusiness     string
	FSPLicenseNumber    string
	LicenseCategory     string
	REExamResults       string
	CPDPoints           int
	ExperienceHistory   string
	CryptoWalletAddress string
	VerificationStatus  bool `gorm:"default:false"`
}

// StartupProfile model represents additional details for Startups
type StartupProfile struct {
	UserID              uint `gorm:"primaryKey"`
	CompanyName         string
	ContactNumber       string
	Location            string
	Website             string
	Industry            string
	SearchPreferences   string
	CryptoWalletAddress string
	VerificationStatus  bool `gorm:"default:false"`
}

// UserDocument model represents uploaded documents for verification purposes
type UserDocument struct {
	ID           uint   `gorm:"primaryKey"`
	UserID       uint   `gorm:"not null"`
	DocumentType string `gorm:"type:varchar(50)"`
	FilePath     string
	UploadedAt   time.Time `gorm:"autoCreateTime"`
	Verified     bool      `gorm:"default:false"`
}

// Message model to store contact messages from Startups to KIs
type Message struct {
	ID          uint      `gorm:"primaryKey"`
	SenderID    uint      `gorm:"not null"` // Startup's UserID
	RecipientID uint      `gorm:"not null"` // KI's UserID
	Subject     string    `gorm:"type:varchar(100);not null"`
	Body        string    `gorm:"type:text;not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
