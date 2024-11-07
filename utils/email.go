package utils

import (
	"fmt"
	"math/rand"
	"time"

	"gopkg.in/gomail.v2"
)

// GenerateOTP creates a random 6-digit OTP
func GenerateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

// SendOTPViaEmail sends the OTP to the user's email address using Gomail
func SendOTPViaEmail(email string, otpCode string) error {
	from := "philaningumede@gmail.com" // Replace with your email
	password := "bpxlrlyjjdblyame"     // Replace with your email password
	smtpHost := "smtp.gmail.com"       // Replace with your SMTP server
	smtpPort := 587                    // Replace with your SMTP port (e.g., 587 for TLS)

	// Create a new Gomail message
	message := gomail.NewMessage()
	message.SetHeader("From", from)
	message.SetHeader("To", email)
	message.SetHeader("Subject", "MFA OTP Code")
	message.SetBody("text/plain", "Your OTP code is: "+otpCode)

	// Create a new Gomail dialer
	dialer := gomail.NewDialer(smtpHost, smtpPort, from, password)

	// Send the email
	return dialer.DialAndSend(message)
}
