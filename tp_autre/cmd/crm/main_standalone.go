package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// Interfaces
type Notifier interface {
	Send(message string) error
	GetType() string
}

type Storer interface {
	Store(notification NotificationRecord) error
	GetAll() []NotificationRecord
}

// Structures
type NotificationRecord struct {
	Type      string
	Message   string
	Timestamp time.Time
}

func (nr NotificationRecord) String() string {
	return fmt.Sprintf("[%s] %s: %s", 
		nr.Timestamp.Format("2006-01-02 15:04:05"), 
		nr.Type, 
		nr.Message)
}

// Notificateurs
type EmailNotifier struct {
	Address string
}

func (e EmailNotifier) Send(message string) error {
	fmt.Printf("Email envoyé à %s: %s\n", e.Address, message)
	return nil
}

func (e EmailNotifier) GetType() string {
	return "Email"
}

type SMSNotifier struct {
	PhoneNumber string
}

func (s SMSNotifier) Send(message string) error {
	if !strings.HasPrefix(s.PhoneNumber, "06") {
		return fmt.Errorf("numéro de téléphone invalide: %s (doit commencer par 06)", s.PhoneNumber)
	}
	fmt.Printf("SMS envoyé au %s: %s\n", s.PhoneNumber, message)
	return nil
}

func (s SMSNotifier) GetType() string {
	return "SMS"
}

type PushNotifier struct {
	DeviceID string
}

func (p PushNotifier) Send(message string) error {
	fmt.Printf("Push envoyée à %s: %s\n", p.DeviceID, message)
	return nil
}

func (p PushNotifier) GetType() string {
	return "Push"
}

// Stockage
type MemoryStorer struct {
	records []NotificationRecord
}

func NewMemoryStorer() *MemoryStorer {
	return &MemoryStorer{
		records: make([]NotificationRecord, 0),
	}
}

func (m *MemoryStorer) Store(notification NotificationRecord) error {
	m.records = append(m.records, notification)
	return nil
}

func (m *MemoryStorer) GetAll() []NotificationRecord {
	return m.records
}

// Application
type App struct {
	storer Storer
	reader *bufio.Reader
}

func NewApp() *App {
	return &App{
		storer: NewMemoryStorer(),
		reader: bufio.NewReader(os.Stdin),
	}
}

func (a *App) Run() {
	fmt.Println("=== Système de Notifications et Logging ===")
	
	for {
		a.showMenu()
		choix := a.getUserInput("Votre choix : ")
		
		switch choix {
		case "1":
			a.envoyerEmail()
		case "2":
			a.envoyerSMS()
		case "3":
			a.envoyerPush()
		case "4":
			a.afficherHistorique()
		case "5":
			fmt.Println("\n=== Historique final des notifications ===")
			a.afficherHistorique()
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide")
		}
	}
}

func (a *App) showMenu() {
	fmt.Println("\n--- Menu ---")
	fmt.Println("1. Envoyer Email")
	fmt.Println("2. Envoyer SMS") 
	fmt.Println("3. Envoyer Push")
	fmt.Println("4. Voir historique")
	fmt.Println("5. Quitter")
}

func (a *App) getUserInput(prompt string) string {
	fmt.Print(prompt)
	input, _ := a.reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func (a *App) envoyerEmail() {
	address := a.getUserInput("Adresse email : ")
	message := a.getUserInput("Message : ")
	
	notifier := EmailNotifier{Address: address}
	a.envoyerNotification(notifier, message)
}

func (a *App) envoyerSMS() {
	phone := a.getUserInput("Numéro de téléphone : ")
	message := a.getUserInput("Message : ")
	
	notifier := SMSNotifier{PhoneNumber: phone}
	a.envoyerNotification(notifier, message)
}

func (a *App) envoyerPush() {
	deviceID := a.getUserInput("ID de l'appareil : ")
	message := a.getUserInput("Message : ")
	
	notifier := PushNotifier{DeviceID: deviceID}
	a.envoyerNotification(notifier, message)
}

func (a *App) envoyerNotification(notifier Notifier, message string) {
	fmt.Printf("\n--- Envoi %s ---\n", notifier.GetType())
	
	err := notifier.Send(message)
	if err != nil {
		fmt.Printf("Erreur: %v\n", err)
	} else {
		record := NotificationRecord{
			Type:      notifier.GetType(),
			Message:   message,
			Timestamp: time.Now(),
		}
		a.storer.Store(record)
		fmt.Printf("Notification envoyée et enregistrée\n")
	}
}

func (a *App) afficherHistorique() {
	records := a.storer.GetAll()
	if len(records) == 0 {
		fmt.Println("Aucune notification enregistrée")
		return
	}
	
	fmt.Printf("\nHistorique (%d notifications) :\n", len(records))
	for _, record := range records {
		fmt.Println(record)
	}
}

func main() {
	app := NewApp()
	app.Run()
}