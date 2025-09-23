package main

import (
	"fmt"
)

// Structure Notification
type Notification struct {
	Type    string
	Message string
}

// Structure Notification Email
type EmailNotifier struct {
	Address string
}

// Structure Notification SMS
type SMSNotifier struct {
	PhoneNumber string
}

// Fonction envoi email
func (e EmailNotifier) SendEmail(message string) {
	fmt.Printf("📧 Email envoyé à %s: %s\n", e.Address, message)
}

// Fonction envoi SMS
func (s SMSNotifier) SendSMS(message string) {
	fmt.Printf("📱 SMS envoyé au %s: %s\n", s.PhoneNumber, message)
}

func main() {
	fmt.Println("=== Système de Notifications - v1 ===")
	
	// Test email
	email := EmailNotifier{Address: "test@example.com"}
	email.SendEmail("Ceci est un test email!")
	
	// Test SMS
	sms := SMSNotifier{PhoneNumber: "0612345678"}
	sms.SendSMS("Message de test SMS")
	
	fmt.Println("Tests terminés")
}