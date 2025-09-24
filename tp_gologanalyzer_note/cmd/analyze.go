package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Pour analyser des logs",
	Long: `Truc d'analyse de logs`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("analyze called")
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)

	// a completer avec flags etc
}
