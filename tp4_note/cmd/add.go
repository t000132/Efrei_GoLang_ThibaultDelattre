package cmd

import (
	"fmt"

	"mini-crm/internal/models"

	"github.com/spf13/cobra"
)

// addCmd représente la commande add
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Ajouter un nouveau contact",
	Long: `Ajouter un nouveau contact dans le CRM.
	
Exemple d'utilisation:
  mini-crm add --name "Jean Dupont" --email "jean@example.com" --phone "0123456789" --company "ACME Corp"
  mini-crm add -n "Marie Martin" -e "marie@test.fr"`,
	RunE: runAdd,
}

var (
	// Flags pour la commande add
	addName    string
	addEmail   string
	addPhone   string
	addCompany string
)

func init() {
	rootCmd.AddCommand(addCmd)

	// Flags obligatoires
	addCmd.Flags().StringVarP(&addName, "name", "n", "", "nom du contact (obligatoire)")
	addCmd.Flags().StringVarP(&addEmail, "email", "e", "", "email du contact (obligatoire)")
	
	// Flags optionnels
	addCmd.Flags().StringVarP(&addPhone, "phone", "p", "", "numéro de téléphone")
	addCmd.Flags().StringVarP(&addCompany, "company", "c", "", "entreprise")

	// Marquer les flags obligatoires
	addCmd.MarkFlagRequired("name")
	addCmd.MarkFlagRequired("email")
}

// runAdd exécute la commande add
func runAdd(cmd *cobra.Command, args []string) error {
	// Créer le nouveau contact
	contact := &models.Contact{
		Name:    addName,
		Email:   addEmail,
		Phone:   addPhone,
		Company: addCompany,
	}

	// Sauvegarder le contact
	if err := storer.Create(contact); err != nil {
		return fmt.Errorf("erreur ajout contact: %v", err)
	}

	fmt.Printf("✅ Contact ajouté avec succès!\n")
	fmt.Printf("ID: %d | %s (%s) | %s | %s\n", 
		contact.ID, contact.Name, contact.Email, contact.Phone, contact.Company)
	
	return nil
}