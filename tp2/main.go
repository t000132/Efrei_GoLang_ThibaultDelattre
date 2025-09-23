package main

import (
	"fmt"
	"strings"
	"time"
)

// Interface Notifier => tous les types de notifications
type Notifier interface {
	Send(message string) error
	GetType() string
}

// Interface Storer => archivage des notificaitons
type Storer interface {
	Store(notification NotificationRecord) error
	GetAll() []NotificationRecord
}

// Structure pour enregistrer les notifications
type NotificationRecord struct {
	Type      string
	Message   string
	Timestamp time.Time
}

// Méthode String pour afficher un enregistrement
func (nr NotificationRecord) String() string {
	return fmt.Sprintf("[%s] %s: %s", 
		nr.Timestamp.Format("2006-01-02 15:04:05"), 
		nr.Type, 
		nr.Message)
}

// EmailNotifier => pour les notifications email
type EmailNotifier struct {
	Address string
}

// Send => implémentation de l'interface Notifier pour Email
func (e EmailNotifier) Send(message string) error {
	fmt.Printf("📧 Email envoyé à %s: %s\n", e.Address, message)
	return nil
}

// GetType => obtention du type de notificateur
func (e EmailNotifier) GetType() string {
	return "Email"
}

// SMSNotifier => pour les notifications SMS
type SMSNotifier struct {
	PhoneNumber string
}

// Send => implémente l'interface Notifier pour SMS
func (s SMSNotifier) Send(message string) error {
	// Validation du numéro de téléphone
	if !strings.HasPrefix(s.PhoneNumber, "06") {
		return fmt.Errorf("numéro de téléphone invalide: %s (doit commencer par 06)", s.PhoneNumber)
	}
	
	fmt.Printf("📱 SMS envoyé au %s: %s\n", s.PhoneNumber, message)
	return nil
}

// GetType => retourne le type de notificateur
func (s SMSNotifier) GetType() string {
	return "SMS"
}

// PushNotifier - pour les notifications push
type PushNotifier struct {
	DeviceID string
}

// Send =>implémente l'interface Notifier pour Push
func (p PushNotifier) Send(message string) error {
	fmt.Printf("� Push envoyée à %s: %s\n", p.DeviceID, message)
	return nil
}

// GetType => retourne le type de notificateur
func (p PushNotifier) GetType() string {
	return "Push"
}

// MemoryStorer => système d'archivage en mémoire
type MemoryStorer struct {
	records []NotificationRecord
}

// Store => enregistre une notification
func (m *MemoryStorer) Store(notification NotificationRecord) error {
	m.records = append(m.records, notification)
	return nil
}

// GetAll => retourne tous les enregistrements
func (m *MemoryStorer) GetAll() []NotificationRecord {
	return m.records
}

func main() {
	fmt.Println("Système de Notifications")
	
	// Création du système d'archivage
	storer := &MemoryStorer{}
	
	// Liste des notificateurs à tester
	notifiers := []Notifier{
		EmailNotifier{Address: "tibo@example.com"},
		SMSNotifier{PhoneNumber: "0610203040"},     // Valide
		SMSNotifier{PhoneNumber: "0612345678"},     // Valide
		SMSNotifier{PhoneNumber: "0687654321"},     // Valide
		SMSNotifier{PhoneNumber: "0712345678"},     // Invalide (ne commence pas par 06)
		SMSNotifier{PhoneNumber: "0512345678"},     // Invalide (ne commence pas par 06)
		PushNotifier{DeviceID: "device123"},
		PushNotifier{DeviceID: "iphone456"},
		PushNotifier{DeviceID: "android789"},
		EmailNotifier{Address: "tibo@test.fr"},
		EmailNotifier{Address: "admin@company.com"},
		EmailNotifier{Address: "support@app.fr"},
	}
	
	message := "Notification de test système"
	
	// Traitement de chaque notificateur
	for _, notifier := range notifiers {
		fmt.Printf("\n--- Test %s ---\n", notifier.GetType())
		
		err := notifier.Send(message)
		if err != nil {
			// Affichage de l'erreur sans arrêter le programme
			fmt.Printf("Erreur: %v\n", err)
		} else {
			// Enregistrement si succès
			record := NotificationRecord{
				Type:      notifier.GetType(),
				Message:   message,
				Timestamp: time.Now(),
			}
			storer.Store(record)
			fmt.Printf("Notification enregistrée\n")
		}
	}
	
	// Affichage de l'historique
	fmt.Println("\nHistorique des notifications :")
	records := storer.GetAll()
	if len(records) == 0 {
		fmt.Println("Aucune notification enregistrée")
	} else {
		for _, record := range records {
			fmt.Println(record)
		}
	}
	
	fmt.Printf("\nTotal: %d notifications réussies\n", len(records))
}