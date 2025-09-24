package cmd

import (
	"fmt"
	"os"

	"github.com/axellelanca/go_loganizer/internal/analyzer"
	"github.com/axellelanca/go_loganizer/internal/config"
	"github.com/axellelanca/go_loganizer/internal/reporter"
	"github.com/spf13/cobra"
)

var (
	configPath string
	outputPath string
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyse des fichiers de logs en parallèle",
	Long: `Analyse plusieurs fichiers de logs de façon concurrente.
			Prend un fichier de config JSON en entrée et peut exporter les résultats dans un fichier JSON.
			Exemple:
  			loganalyzer analyze -c config.json -o rapport.json`,
	Run: executeAnalysis,
}

func executeAnalysis(cmd *cobra.Command, args []string) {
	if configPath == "" {
		fmt.Println("Erreur: le flag --config (-c) est obligatoire")
		cmd.Help()
		os.Exit(1)
	}

	fmt.Printf("Début de l'analyse avec: %s\n", configPath)

	// Chargement config
	logConfigs, err := config.LoadConfig(configPath)
	if err != nil {
		fmt.Printf("Erreur config: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Config chargée: %d fichiers de logs\n", len(logConfigs))

	// Lancement analyse en parallèle
	fmt.Println("Analyse en cours...")
	results := analyzer.AnalyzeLogsConcurrently(logConfigs)

	// Affichage résultats
	reporter.PrintResults(results)

	// Export JSON si demandé
	if outputPath != "" {
		// Nom avec timestamp si c'est un nom générique
		finalOutputPath := outputPath
		if outputPath == "report.json" || outputPath == "rapport.json" {
			finalOutputPath = reporter.GenerateTimestampedFilename(outputPath)
		}

		fmt.Printf("Export vers: %s\n", finalOutputPath)
		if err := reporter.ExportResults(results, finalOutputPath); err != nil {
			fmt.Printf("Erreur export: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Export réussi!\n")
	}

	fmt.Println("Analyse terminée!")
}

func init() {
	rootCmd.AddCommand(analyzeCmd)

	// Flags
	analyzeCmd.Flags().StringVarP(&configPath, "config", "c", "", 
		"Fichier de config JSON (obligatoire)")
	analyzeCmd.Flags().StringVarP(&outputPath, "output", "o", "", 
		"Fichier de sortie JSON (optionnel)")
	
	// Config obligatoire
	analyzeCmd.MarkFlagRequired("config")
}
