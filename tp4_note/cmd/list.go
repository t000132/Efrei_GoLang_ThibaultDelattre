package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lister les contacts",
	Long: `Afficher la liste de tous les contacts ou un contact spécifique par ID.

		Exemples d'utilisation:
  mini-crm list          # Va lister tous les contacts
  mini-crm list --id 5   # Afficher le contact avec l'ID 5
  mini-crm list -i 3     # Afficher le contact avec l'ID 3`,
	RunE: runList,
}

var (
	// Flag pour l'ID spécifique
	listID string
)

func init() {
	rootCmd.AddCommand(listCmd)

	// Flag optionnel pour afficher un contact spécifique
	listCmd.Flags().StringVarP(&listID, "id", "i", "", "ID du contact à afficher")
}

// runList exécute la commande list
func runList(cmd *cobra.Command, args []string) error {
	// Si un ID est spécifié, afficher ce contact uniquement
	if listID != "" {
		return showContactByID()
	}

	// Sinon, afficher tous les contacts
	return showAllContacts()
}

// showContactByID affiche un contact spécifique par son ID
func showContactByID() error {
	// Convertir l'ID en uint
	id, err := strconv.ParseUint(listID, 10, 32)
	if err != nil {
		return fmt.Errorf("ID invalide: %s", listID)
	}

	// Récupérer le contact
	contact, err := storer.GetByID(uint(id))
	if err != nil {
		return err
	}

	fmt.Printf("Contact trouvé:\n")
	fmt.Printf("ID: %d\n", contact.ID)
	fmt.Printf("Nom: %s\n", contact.Name)
	fmt.Printf("Email: %s\n", contact.Email)
	fmt.Printf("Téléphone: %s\n", contact.Phone)
	fmt.Printf("Entreprise: %s\n", contact.Company)
	fmt.Printf("Créé le: %s\n", contact.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("Modifié le: %s\n", contact.UpdatedAt.Format("2006-01-02 15:04:05"))

	return nil
}

// showAllContacts affiche tous les contacts
func showAllContacts() error {
	// Récupérer tous les contacts
	contacts, err := storer.GetAll()
	if err != nil {
		return fmt.Errorf("erreur récupération contacts: %v", err)
	}

	if len(contacts) == 0 {
		fmt.Println("Aucun contact trouvé.")
		return nil
	}

	fmt.Printf("%d contact(s) trouvé(s):\n\n", len(contacts))
	
	for _, contact := range contacts {
		fmt.Printf("ID: %d | %s (%s) | %s | %s\n", 
			contact.ID, contact.Name, contact.Email, contact.Phone, contact.Company)
	}

	return nil
}