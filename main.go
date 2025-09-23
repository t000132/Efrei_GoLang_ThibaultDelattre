
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// memo technique :
// fmt = format / le package de base de Go en gros
// bufio = package pour lire / écrire des flux de données via mémoire tampon
// Stdin = standard input (clavier)
// nil = null
// strconv.Atoi = packagae qui convertit les types => ASCII to integer
// _ = pour ignorer la valeur de retour par ex une erreur

// Contact structure
type Contact struct {
	ID    int
	Nom   string
	Email string
}

// Constructeur pour créer un nouveau contact
func NewContact(nom, email string) (*Contact, error) {
	// Validation
	if strings.TrimSpace(nom) == "" {
		return nil, fmt.Errorf("le nom ne peut pas être vide")
	}
	if strings.TrimSpace(email) == "" {
		return nil, fmt.Errorf("l'email ne peut pas être vide")
	}
	if !strings.Contains(email, "@") {
		return nil, fmt.Errorf("l'email doit contenir @")
	}

	return &Contact{
		ID:    nextID,
		Nom:   strings.TrimSpace(nom),
		Email: strings.TrimSpace(email),
	}, nil
}

// Méthode pour afficher le contact
func (c *Contact) String() string {
	return fmt.Sprintf("ID: %d, Nom: %s, Email: %s", c.ID, c.Nom, c.Email)
}

// Méthode pour mettre à jour le nom
func (c *Contact) UpdateNom(nouveauNom string) error {
	if strings.TrimSpace(nouveauNom) == "" {
		return fmt.Errorf("le nom ne peut pas être vide")
	}
	c.Nom = strings.TrimSpace(nouveauNom)
	return nil
}

// Méthode pour mettre à jour l'email
func (c *Contact) UpdateEmail(nouvelEmail string) error {
	if strings.TrimSpace(nouvelEmail) == "" {
		return fmt.Errorf("l'email ne peut pas être vide")
	}
	if !strings.Contains(nouvelEmail, "@") {
		return fmt.Errorf("l'email doit contenir @")
	}
	c.Email = strings.TrimSpace(nouvelEmail)
	return nil
}

// Map pour le stockage des contacts avec pointeurs
var contacts = make(map[int]*Contact)
var nextID = 1

func main() {
	// Définition des flags
	ajouter := flag.Bool("ajouter", false, "Ajouter un contact via ligne de commande")
	nom := flag.String("nom", "", "Nom du contact")
	email := flag.String("mail", "", "Email du contact")
	flag.Parse()

	// Si le flag --ajouter est utilisé
	if *ajouter {
		ajouterContactViaFlags(*nom, *email)
		return
	}

	// Mode CLI normal
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nMini-CRM")
		fmt.Println("1. Ajouter contact")
		fmt.Println("2. Lister contacts")
		fmt.Println("3. Supprimer contact")
		fmt.Println("4. Mettre à jour un contact")
		fmt.Println("5. Quitter")
		fmt.Print("Votre choix : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			ajouterContact(reader)
		case "2":
			listerContacts()
		case "3":
			supprimerContact(reader)
		case "4":
			mettreAJourContact(reader)
		case "5":
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func ajouterContact(reader *bufio.Reader) {
	fmt.Print("Nom : ")
	nom, _ := reader.ReadString('\n')
	fmt.Print("Email : ")
	email, _ := reader.ReadString('\n')

	contact, err := NewContact(nom, email)
	if err != nil {
		fmt.Printf("Erreur : %v\n", err)
		return
	}

	contact.ID = nextID
	contacts[nextID] = contact
	fmt.Printf("Contact ajouté : %s\n", contact.String())
	nextID++
}

func listerContacts() {
	if len(contacts) == 0 {
		fmt.Println("Aucun contact")
		return
	}
	fmt.Println("Liste des contacts :")
	for _, c := range contacts {
		fmt.Println(c.String())
	}
}

func supprimerContact(reader *bufio.Reader) {
	fmt.Print("ID du contact à supprimer : ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID invalide")
		return
	}
	if _, ok := contacts[id]; ok {
		delete(contacts, id)
		fmt.Println("Contact supprimé")
	} else {
		fmt.Println("Contact non trouvé")
	}
}

func mettreAJourContact(reader *bufio.Reader) {
	fmt.Print("ID du contact à mettre à jour : ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID invalide")
		return
	}
	
	contact, ok := contacts[id]
	if !ok {
		fmt.Println("Contact non trouvé")
		return
	}
	
	fmt.Printf("Contact actuel - %s\n", contact.String())
	
	fmt.Print("Nouveau nom (appuyez sur Entrée pour garder l'actuel) : ")
	nom, _ := reader.ReadString('\n')
	nom = strings.TrimSpace(nom)
	if nom != "" {
		err := contact.UpdateNom(nom)
		if err != nil {
			fmt.Printf("Erreur nom : %v\n", err)
			return
		}
	}
	
	fmt.Print("Nouvel email (veuillez appuyer sur Entrée pour garder l'actuel) : ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	if email != "" {
		err := contact.UpdateEmail(email)
		if err != nil {
			fmt.Printf("Erreur email : %v\n", err)
			return
		}
	}
	
	fmt.Printf("Contact mis à jour : %s\n", contact.String())
}

func ajouterContactViaFlags(nom, email string) {
	contact, err := NewContact(nom, email)
	if err != nil {
		fmt.Printf("Erreur : %v\n", err)
		fmt.Println("Exemple : go run main.go --ajouter --nom=Axelle --mail=axelle@lanca.fr")
		return
	}

	// Attribution de l'ID et ajout à la map
	contact.ID = nextID
	contacts[nextID] = contact
	
	// Affichage du résultat
	fmt.Printf("Contact ajouté parfaitement !\n")
	fmt.Printf("%s\n", contact.String())
	nextID++
}