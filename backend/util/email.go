package util

import (
	"log"
	"net/smtp"
)

func SendEmail(to, subject, body string) (erroryes error) {
	from := "pharmadash343@gmail.com"
	pass := "jbqr fyfr khhf ypfz"

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
		return err
	}

	log.Printf("Email sent successfully to %s", to)
	return nil
}
