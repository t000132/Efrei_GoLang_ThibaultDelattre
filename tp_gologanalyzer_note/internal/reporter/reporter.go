package reporter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/axellelanca/go_loganizer/internal/config"
)

// ExportResults sauve les résultats en JSON
func ExportResults(results []config.AnalysisResult, outputPath string) error {
	// Créer les dossiers si besoin
	if err := createDirectoriesIfNeeded(outputPath); err != nil {
		return fmt.Errorf("impossible de créer les dossiers: %w", err)
	}

	// Convertir en JSON avec indentation
	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return fmt.Errorf("erreur sérialisation JSON: %w", err)
	}

	// Écrire le fichier
	if err := os.WriteFile(outputPath, jsonData, 0644); err != nil {
		return fmt.Errorf("erreur écriture fichier: %w", err)
	}

	return nil
}

// GenerateTimestampedFilename ajoute la date au nom de fichier
func GenerateTimestampedFilename(basePath string) string {
	now := time.Now()
	timestamp := now.Format("060102") // Format AAMMJJ
	
	dir := filepath.Dir(basePath)
	ext := filepath.Ext(basePath)
	nameWithoutExt := filepath.Base(basePath[:len(basePath)-len(ext)])
	
	timestampedName := fmt.Sprintf("%s_%s%s", timestamp, nameWithoutExt, ext)
	return filepath.Join(dir, timestampedName)
}

// createDirectoriesIfNeeded crée les dossiers manquants
func createDirectoriesIfNeeded(filePath string) error {
	dir := filepath.Dir(filePath)
	if dir == "." || dir == "/" {
		return nil
	}
	
	return os.MkdirAll(dir, 0755)
}

// PrintResults affiche un résumé des résultats sur la console
func PrintResults(results []config.AnalysisResult) {
	fmt.Println("\n=== RÉSUMÉ DE L'ANALYSE ===")
	fmt.Printf("Total de fichiers analysés: %d\n\n", len(results))

	successCount := 0
	failedCount := 0

	for _, result := range results {
		if result.Status == config.StatusFailed {
			failedCount++
		} else {
			successCount++
		}

		fmt.Printf("[%s] %s\n", result.LogID, result.FilePath)
		fmt.Printf("   Status: %s\n", result.Status)
		fmt.Printf("   Message: %s\n", result.Message)
		
		if result.ErrorDetails != "" {
			fmt.Printf("   Erreur: %s\n", result.ErrorDetails)
		}
		fmt.Println()
	}

	fmt.Printf("=== BILAN ===\n")
	fmt.Printf("Succès: %d | Échecs: %d\n", successCount, failedCount)
}