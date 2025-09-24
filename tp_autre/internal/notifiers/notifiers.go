package notifiers

import (
	"fmt"
	"strings"
)

// EmailNotifier - pour les notifications email
type EmailNotifier struct {
	Address string
}

// Send - implémente l'interface Notifier pour Email
func (e EmailNotifier) Send(message string) error {
	fmt.Printf("📧 Email envoyé à %s: %s\n", e.Address, message)
	return nil
}

// GetType - retourne le type de notificateur
func (e EmailNotifier) GetType() string {
	return "Email"
}

// SMSNotifier - pour les notifications SMS
type SMSNotifier struct {
	PhoneNumber string
}

// Send - implémente l'interface Notifier pour SMS avec validation
func (s SMSNotifier) Send(message string) error {
	// Validation du numéro de téléphone
	if !strings.HasPrefix(s.PhoneNumber, "06") {
		return fmt.Errorf("numéro de téléphone invalide: %s (doit commencer par 06)", s.PhoneNumber)
	}
	
	fmt.Printf("📱 SMS envoyé au %s: %s\n", s.PhoneNumber, message)
	return nil
}

// GetType - retourne le type de notificateur
func (s SMSNotifier) GetType() string {
	return "SMS"
}

// PushNotifier - pour les notifications push
type PushNotifier struct {
	DeviceID string
}

// Send - implémente l'interface Notifier pour Push
func (p PushNotifier) Send(message string) error {
	fmt.Printf("Push envoyée à %s: %s\n", p.DeviceID, message)
	return nil
}

// GetType - retourne le type de notificateur
func (p PushNotifier) GetType() string {
	return "Push"
}