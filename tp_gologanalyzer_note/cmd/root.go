package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go_loganizer",
	Short: "Outil d'analyse de logs",
	Long: `Go_loganizer est un outil pour analyser des fichiers de logs et en extraire des informations utiles.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Message de test pour toggle")
}


