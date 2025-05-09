package service

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gopkg.in/mail.v2"
)

// EmailService handles sending emails
type EmailService struct {
	SMTPHost string
	SMTPPort int
	Username string
	Password string
}

// NewEmailService creates a new email service with the provided configuration
func NewEmailService() *EmailService {
	smtpHost := os.Getenv("SMTP_HOST")
	if smtpHost == "" {
		smtpHost = "smtp.gmail.com"
	}

	smtpPortStr := os.Getenv("SMTP_PORT")
	if smtpPortStr == "" {
		smtpPortStr = "587"
	}
	smtpPort, _ := strconv.Atoi(smtpPortStr)

	return &EmailService{
		SMTPHost: smtpHost,
		SMTPPort: smtpPort,
		Username: os.Getenv("EMAIL_USER"),
		Password: os.Getenv("EMAIL_PASS"),
	}
}

// SendEmail sends an email with the provided details
func (s *EmailService) SendEmail(to, subject, body string) error {
	m := mail.NewMessage()
	m.SetHeader("From", s.Username)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := mail.NewDialer(s.SMTPHost, s.SMTPPort, s.Username, s.Password)

	if err := d.DialAndSend(m); err != nil {
		log.Printf("Failed to send email: %v", err)
		return fmt.Errorf("failed to send email: %w", err)
	}

	log.Printf("Email sent successfully to %s", to)
	return nil
}
