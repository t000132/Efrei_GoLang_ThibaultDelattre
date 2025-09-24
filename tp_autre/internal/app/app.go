package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"tp3/internal/models"
	"tp3/internal/notifiers"
	"tp3/internal/storage"
)

// App - structure principale de l'application
type App struct {
	storer models.Storer
	reader *bufio.Reader
}

// NewApp - constructeur pour l'application
func NewApp() *App {
	return &App{
		storer: storage.NewMemoryStorer(),
		reader: bufio.NewReader(os.Stdin),
	}
}

// Run - lance l'application CLI
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

// showMenu - affiche le menu principal
func (a *App) showMenu() {
	fmt.Println("\n--- Menu ---")
	fmt.Println("1. Envoyer Email")
	fmt.Println("2. Envoyer SMS") 
	fmt.Println("3. Envoyer Push")
	fmt.Println("4. Voir historique")
	fmt.Println("5. Quitter")
}

// getUserInput - récupère la saisie utilisateur
func (a *App) getUserInput(prompt string) string {
	fmt.Print(prompt)
	input, _ := a.reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// envoyerEmail - gère l'envoi d'email
func (a *App) envoyerEmail() {
	address := a.getUserInput("Adresse email : ")
	message := a.getUserInput("Message : ")
	
	notifier := notifiers.EmailNotifier{Address: address}
	a.envoyerNotification(notifier, message)
}

// envoyerSMS - gère l'envoi de SMS
func (a *App) envoyerSMS() {
	phone := a.getUserInput("Numéro de téléphone : ")
	message := a.getUserInput("Message : ")
	
	notifier := notifiers.SMSNotifier{PhoneNumber: phone}
	a.envoyerNotification(notifier, message)
}

// envoyerPush - gère l'envoi de notification push
func (a *App) envoyerPush() {
	deviceID := a.getUserInput("ID de l'appareil : ")
	message := a.getUserInput("Message : ")
	
	notifier := notifiers.PushNotifier{DeviceID: deviceID}
	a.envoyerNotification(notifier, message)
}

// envoyerNotification - logique commune d'envoi
func (a *App) envoyerNotification(notifier models.Notifier, message string) {
	fmt.Printf("\n--- Envoi %s ---\n", notifier.GetType())
	
	err := notifier.Send(message)
	if err != nil {
		fmt.Printf("Erreur: %v\n", err)
	} else {
		// Enregistrement si succès
		record := models.NotificationRecord{
			Type:      notifier.GetType(),
			Message:   message,
			Timestamp: time.Now(),
		}
		a.storer.Store(record)
		fmt.Printf("Notification envoyée et enregistrée\n")
	}
}

// afficherHistorique - affiche l'historique des notifications
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