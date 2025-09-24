package cmd  

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "gowatcher",
	Short: "Gowatcher est un outil pour vérifier l'accessibilité des URLs",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Erreur lors de l'exécution de la commande: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	// Ici, vous pouvez ajouter des sous-commandes au rootCmd si nécessaire
	