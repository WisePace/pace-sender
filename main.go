package main

import (
	"log"
	"pace-sender/pkg/email"
)

func main() {
	// Get the list of valid email addresses
	validEmails, err := email.GetValidEmails()
	if err != nil {
		log.Fatalf("Error validating emails: %v", err)
	}

	// Send emails only to valid email addresses
	err = email.SendEmailsToValidRecipients(validEmails)
	if err != nil {
		log.Fatalf("Error sending emails: %v", err)
	}
}
