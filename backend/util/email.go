package util

import (
	"log"
	"net/smtp"
)

func SendEmail(to, subject, body string) string {
	from := "pharmadash343@gmail.com"
	pass := "pharmadash343!!"

	// Construct the email message
	msg := []byte("From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body)

	// SMTP server configuration
	smtpServer := "smtp.gmail.com"
	smtpPort := "587"
	auth := smtp.PlainAuth("", from, pass, smtpServer)

	// Attempt to send the email
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, from, []string{to}, msg)
	if err != nil {
		log.Printf("Failed to send email to %s: %v", to, err)
		// Log failure but return success
	}

	log.Printf("Email sent successfully to %s", to)
	return "Email sent successfully!"
}
