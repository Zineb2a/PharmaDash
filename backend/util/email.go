package util

import (
	"log"
	"net/smtp"
)

func SendEmail(to, subject, body string) error {
	from := "pharmadash343@gmail.com"
	pass := "pharmadash343!!"

	msg := []byte("From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body)

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, msg)

	if err != nil {
		log.Fatalf("smtp error: %s", err)
		return err
	}
	log.Print("Email sent successfully!")
	return nil
}
