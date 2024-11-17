package models

import (
	"time"

	"github.com/lib/pq"
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
	Address    string    `gorm:"type:varchar(255)"`
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

// KeyIndividualProfile defines the schema for key individuals
type KeyIndividualProfile struct {
	ID              uint           `gorm:"primaryKey"`
	UserID          uint           `gorm:"not null"`          // Foreign key linking to the User model
	User            User           `gorm:"foreignKey:UserID"` // Association with the User model
	FullName        string         `gorm:"not null"`          // Full name of the key individual
	Qualifications  pq.StringArray `gorm:"type:text[]"`       // Array of qualifications
	Experience      pq.StringArray `gorm:"type:text[]"`       // Array of experiences
	ContactDetails  string         `gorm:"not null"`          // Contact details
	Area            string         `gorm:"not null"`          // Area where the key individual operates
	AssetTypes      pq.StringArray `gorm:"type:text[]"`       // Array of asset types managed
	ClassOfBusiness pq.StringArray `gorm:"type:text[]"`       // Array of class of business
	REExams         pq.StringArray `gorm:"type:text[]"`       // Array of RE exam results
	CPDPoints       int            `gorm:"default:0"`         // CPD points
}

// StartupProfile defines the schema for startups
type StartupProfile struct {
	ID                 uint   `gorm:"primaryKey"`
	UserID             uint   `gorm:"not null"`          // Foreign key linking to the User model
	User               User   `gorm:"foreignKey:UserID"` // Association with the User model
	Name               string `gorm:"not null"`          // Name of the startup
	Industry           string // Industry of the startup
	Website            string // Optional website URL
	ContactInformation string `gorm:"not null"` // Contact information
	Area               string `gorm:"not null"` // Area where the startup is located
}

// UserDocument model represents uploaded documents for verification purposes
type UserDocument struct {
	ID           uint      `gorm:"primaryKey"`
	UserID       uint      `gorm:"not null"`
	DocumentType string    `gorm:"type:varchar(50)"`
	FileData     []byte    `gorm:"type:bytea"`                // Use bytea for PostgreSQL
	Hash         string    `gorm:"type:varchar(64);not null"` // SHA-256 hash of the document
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
	SentAt      time.Time `gorm:"autoCreateTime"`
}
type OTP struct {
	ID        uint      `gorm:"primaryKey"`
	Email     string    `gorm:"unique;not null"` // Email associated with the OTP
	Code      string    `gorm:"not null"`        // The OTP code
	ExpiresAt time.Time // Expiration time for the OTP
}
