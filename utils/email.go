package utils

import (
	"os"

	"gopkg.in/gomail.v2"
)

// SendEmail sends an email to a specified address with a subject and body.
func SendEmail(to string, subject string, body string) error {
	from := os.Getenv("SMTP_USER")     // Email address from environment variables
	password := os.Getenv("SMTP_PASS") // Email password from environment variables
	smtpHost := os.Getenv("SMTP_HOST") // SMTP host from environment variables
	smtpPort := 587                    // SMTP port (commonly 587 for TLS)

	// Create a new Gomail message
	message := gomail.NewMessage()
	message.SetHeader("From", from)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody("text/plain", body)

	// Create a new Gomail dialer with SMTP settings
	dialer := gomail.NewDialer(smtpHost, smtpPort, from, password)

	// Send the email
	return dialer.DialAndSend(message)
}
