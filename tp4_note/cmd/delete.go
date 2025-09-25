package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Supprimer contact",
	Long: `Supprimer contact du CRM par son ID.
	Une confirmation est demandée avant la suppression.

Exemples d'utilisation:
  mini-crm delete --id 5
  mini-crm delete -i 3
  mini-crm delete --id 1 --force  # Supprimer sans confirmation`,
	RunE: runDelete,
}

var (
	// Flags de la commande delete
	deleteID    string
	deleteForce bool
)

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Flags
	deleteCmd.Flags().StringVarP(&deleteID, "id", "i", "", "ID du contact à supprimer (à mettre)")
	deleteCmd.Flags().BoolVarP(&deleteForce, "force", "f", false, "supprimer sans confirmation")

	// Marquer l'ID comme obligatoire
	deleteCmd.MarkFlagRequired("id")
}

// runDelete exécute la commande delete
func runDelete(cmd *cobra.Command, args []string) error {
	// Convertir l'ID en uint afin de l'utiliser pour la suppression
	id, err := strconv.ParseUint(deleteID, 10, 32)
	if err != nil {
		return fmt.Errorf("ID invalide: %s", deleteID)
	}

	// Récupere le contact à supprimer pour affichage
	contact, err := storer.GetByID(uint(id))
	if err != nil {
		return err
	}

	// Affiche le contact qui va être supprimé
	fmt.Printf("Contact à supprimer:\n")
	fmt.Printf("ID: %d | %s (%s) | %s | %s\n", 
		contact.ID, contact.Name, contact.Email, contact.Phone, contact.Company)

	// Demander confirmation si --force n'est pas utilisé car la suppression est irréversible
	if !deleteForce {
		if !confirmDelete() {
			fmt.Println("Suppression annulée.")
			return nil
		}
	}

	// Supprimer le contact
	if err := storer.Delete(uint(id)); err != nil {
		return fmt.Errorf("erreur suppression contact: %v", err)
	}

	fmt.Printf("Contact %d supprimé avec succès!\n", id)
	return nil
}

// confirmDelete demande confirmation à l'utilisateur
func confirmDelete() bool {
	fmt.Print("\n Êtes-vous sûr de vouloir supprimer ce contact ? (oui/non): ")
	
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		return false
	}

	response = strings.TrimSpace(strings.ToLower(response))
	return response == "oui" || response == "o" || response == "yes" || response == "y"
}