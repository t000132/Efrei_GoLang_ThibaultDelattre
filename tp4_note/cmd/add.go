package cmd

import (
	"fmt"

	"mini-crm/internal/models"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Ajouter nouveau contact",
	Long: `Ajouter nouveau contact dans le CRM.
	
Exemple d'utilisation:
mini-crm add --name "Lionel Messi" --email "messi@example.com" --phone "0110203040" --company "FIFA Corp"
mini-crm add -n "Cristiano Ronaldo" -e "cristiano@test.fr"`,
	RunE: runAdd,
}

var (
	// Flags de la commande add
	addName    string
	addEmail   string
	addPhone   string
	addCompany string
)

func init() {
	rootCmd.AddCommand(addCmd)

	// Flags obligatoires
	addCmd.Flags().StringVarP(&addName, "name", "n", "", "nom du contact (à mettre)")
	addCmd.Flags().StringVarP(&addEmail, "email", "e", "", "email du contact (à mettre)")
	
	// Flags pas obligatoires
	addCmd.Flags().StringVarP(&addPhone, "phone", "p", "", "numéro de téléphone")
	addCmd.Flags().StringVarP(&addCompany, "company", "c", "", "entreprise")

	// Marquer les flags obligatoires
	addCmd.MarkFlagRequired("name")
	addCmd.MarkFlagRequired("email")
}

// runAdd => exécute la commande add
func runAdd(cmd *cobra.Command, args []string) error {
	// Créer nouveau contact
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

	fmt.Printf("Contact ajouté !\n")
	fmt.Printf("ID: %d | %s (%s) | %s | %s\n", 
		contact.ID, contact.Name, contact.Email, contact.Phone, contact.Company)
	
	return nil
}