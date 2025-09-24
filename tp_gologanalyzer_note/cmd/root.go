package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "loganalyzer",
	Short: "Outil d'analyse de logs",
	Long: `loganalyzer est un outil CLI pour analyser des fichiers de logs de diverses sources (serveurs, applications).
			Il analyse plusieurs logs en parallèle et extrait des infos utiles avec gestion d'erreurs robuste.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Pas de flags glovbaux
}


