package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateOTP creates a random 6-digit OTP
func GenerateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

// SendOTPViaEmail sends the OTP to the user's email address
func SendOTPViaEmail(email, otpCode string) error {
	subject := "Your OTP Code"
	body := "Your OTP code is: " + otpCode
	return SendEmail(email, subject, body)
}
