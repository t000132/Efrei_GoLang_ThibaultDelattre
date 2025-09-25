package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// updateCmd représente la commande update
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Mettre à jour un contact existant",
	Long: `Mettre à jour les informations d'un contact existant.
L'ID est obligatoire, les autres champs sont optionnels.

Exemples d'utilisation:
  mini-crm update --id 5 --name "Jean Dupont Jr"
  mini-crm update -i 3 -e "nouveau@email.com" -p "0987654321"
  mini-crm update --id 1 --company "Nouvelle Entreprise"`,
	RunE: runUpdate,
}

var (
	// Flags pour la commande update
	updateID      string
	updateName    string
	updateEmail   string
	updatePhone   string
	updateCompany string
)

func init() {
	rootCmd.AddCommand(updateCmd)

	// Flag obligatoire
	updateCmd.Flags().StringVarP(&updateID, "id", "i", "", "ID du contact à mettre à jour (obligatoire)")
	
	// Flags optionnels pour les champs à modifier
	updateCmd.Flags().StringVarP(&updateName, "name", "n", "", "nouveau nom")
	updateCmd.Flags().StringVarP(&updateEmail, "email", "e", "", "nouvel email")
	updateCmd.Flags().StringVarP(&updatePhone, "phone", "p", "", "nouveau téléphone")
	updateCmd.Flags().StringVarP(&updateCompany, "company", "c", "", "nouvelle entreprise")

	// Marquer l'ID comme obligatoire
	updateCmd.MarkFlagRequired("id")
}

// runUpdate exécute la commande update
func runUpdate(cmd *cobra.Command, args []string) error {
	// Convertir l'ID en uint
	id, err := strconv.ParseUint(updateID, 10, 32)
	if err != nil {
		return fmt.Errorf("ID invalide: %s", updateID)
	}

	// Récupérer le contact existant
	contact, err := storer.GetByID(uint(id))
	if err != nil {
		return err
	}

	// Afficher le contact avant modification
	fmt.Printf("📋 Contact actuel:\n")
	fmt.Printf("ID: %d | %s (%s) | %s | %s\n", 
		contact.ID, contact.Name, contact.Email, contact.Phone, contact.Company)

	// Mettre à jour seulement les champs spécifiés
	modified := false
	
	if updateName != "" && updateName != contact.Name {
		contact.Name = updateName
		modified = true
	}
	
	if updateEmail != "" && updateEmail != contact.Email {
		contact.Email = updateEmail
		modified = true
	}
	
	if updatePhone != contact.Phone {
		contact.Phone = updatePhone
		modified = true
	}
	
	if updateCompany != contact.Company {
		contact.Company = updateCompany
		modified = true
	}

	// Vérifier qu'au moins un champ a été modifié
	if !modified {
		fmt.Println("⚠️  Aucune modification détectée.")
		return nil
	}

	// Sauvegarder les modifications
	if err := storer.Update(contact); err != nil {
		return fmt.Errorf("erreur mise à jour contact: %v", err)
	}

	fmt.Printf("\n✅ Contact mis à jour avec succès!\n")
	fmt.Printf("ID: %d | %s (%s) | %s | %s\n", 
		contact.ID, contact.Name, contact.Email, contact.Phone, contact.Company)

	return nil
}