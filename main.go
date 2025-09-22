
package main

import (
	"bufio"
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

// Map pour le stockjage des contacts
var contacts = make(map[int]Contact)
var nextID = 1

func main() {
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
	nom = strings.TrimSpace(nom)
	fmt.Print("Email : ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	contact := Contact{ID: nextID, Nom: nom, Email: email}
	contacts[nextID] = contact
	fmt.Printf("Contact ajouté : %+v\n", contact)
	nextID++
}

func listerContacts() {
	if len(contacts) == 0 {
		fmt.Println("Aucun contact")
		return
	}
	fmt.Println("Liste des contacts :")
	for _, c := range contacts {
		fmt.Printf("ID: %d, Nom: %s, Email: %s\n", c.ID, c.Nom, c.Email)
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
	
	fmt.Printf("Contact actuel - ID: %d, Nom: %s, Email: %s\n", contact.ID, contact.Nom, contact.Email)
	
	fmt.Print("Nouveau nom (veuillez appuyer sur Entrée pour garder l'actuel) : ")
	nom, _ := reader.ReadString('\n')
	nom = strings.TrimSpace(nom)
	if nom != "" {
		contact.Nom = nom
	}
	
	fmt.Print("Nouvel email (veuillez appuyer sur Entrée pour garder l'actuel) : ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	if email != "" {
		contact.Email = email
	}
	
	contacts[id] = contact
	fmt.Printf("Contact mis à jour : %+v\n", contact)
}