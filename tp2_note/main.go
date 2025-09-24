package main

import (
	"bufio"
	"fmt"
	"os"
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
	fmt.Println("=== Système de Notifications et Logging ===")
	
	// Création du système d'archivage
	storer := &MemoryStorer{}
	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Println("\nMenu :")
		fmt.Println("1. Envoyer Email")
		fmt.Println("2. Envoyer SMS") 
		fmt.Println("3. Envoyer Push")
		fmt.Println("4. Voir historique")
		fmt.Println("5. Quitter")
		fmt.Print("Votre choix : ")
		
		choix, _ := reader.ReadString('\n')
		choix = strings.TrimSpace(choix)
		
		switch choix {
		case "1":
			envoyerEmail(reader, storer)
		case "2":
			envoyerSMS(reader, storer)
		case "3":
			envoyerPush(reader, storer)
		case "4":
			afficherHistorique(storer)
		case "5":
			fmt.Println("\nHistorique final des notifications :")
			afficherHistorique(storer)
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide")
		}
	}
}

func envoyerEmail(reader *bufio.Reader, storer *MemoryStorer) {
	fmt.Print("Adresse email : ")
	address, _ := reader.ReadString('\n')
	address = strings.TrimSpace(address)
	
	fmt.Print("Message : ")
	message, _ := reader.ReadString('\n')
	message = strings.TrimSpace(message)
	
	notifier := EmailNotifier{Address: address}
	envoyerNotification(notifier, message, storer)
}

func envoyerSMS(reader *bufio.Reader, storer *MemoryStorer) {
	fmt.Print("Numéro de téléphone : ")
	phone, _ := reader.ReadString('\n')
	phone = strings.TrimSpace(phone)
	
	fmt.Print("Message : ")
	message, _ := reader.ReadString('\n')
	message = strings.TrimSpace(message)
	
	notifier := SMSNotifier{PhoneNumber: phone}
	envoyerNotification(notifier, message, storer)
}

func envoyerPush(reader *bufio.Reader, storer *MemoryStorer) {
	fmt.Print("ID de l'appareil : ")
	deviceID, _ := reader.ReadString('\n')
	deviceID = strings.TrimSpace(deviceID)
	
	fmt.Print("Message : ")
	message, _ := reader.ReadString('\n')
	message = strings.TrimSpace(message)
	
	notifier := PushNotifier{DeviceID: deviceID}
	envoyerNotification(notifier, message, storer)
}

func envoyerNotification(notifier Notifier, message string, storer *MemoryStorer) {
	fmt.Printf("\n--- Envoi %s ---\n", notifier.GetType())
	
	err := notifier.Send(message)
	if err != nil {
		fmt.Printf("Erreur: %v\n", err)
	} else {
		// Enregistrement si succès
		record := NotificationRecord{
			Type:      notifier.GetType(),
			Message:   message,
			Timestamp: time.Now(),
		}
		storer.Store(record)
		fmt.Printf("Notification envoyé et enregistrée\n")
	}
}

func afficherHistorique(storer *MemoryStorer) {
	records := storer.GetAll()
	if len(records) == 0 {
		fmt.Println("Aucune notification enregistrée")
		return
	}
	
	fmt.Printf("\nHistorique (%d notifications) :\n", len(records))
	for _, record := range records {
		fmt.Println(record)
	}
}