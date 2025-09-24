package analyzer

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/axellelanca/go_loganizer/internal/config"
)

// AnalyzeLogsConcurrently lance l'analyse en parallèle
func AnalyzeLogsConcurrently(logConfigs []config.LogConfig) []config.AnalysisResult {
	var wg sync.WaitGroup
	results := make(chan config.AnalysisResult, len(logConfigs))
	
	// Une goroutine par fichier
	for _, logConfig := range logConfigs {
		wg.Add(1)
		go func(cfg config.LogConfig) {
			defer wg.Done()
			result := analyzeLogFile(cfg)
			results <- result
		}(logConfig)
	}

	// Fermer le channel quand tout est fini
	go func() {
		wg.Wait()
		close(results)
	}()

	// Récupération des résultats
	var allResults []config.AnalysisResult
	for result := range results {
		allResults = append(allResults, result)
	}

	return allResults
}

// analyzeLogFile analyse un fichier
func analyzeLogFile(logConfig config.LogConfig) config.AnalysisResult {
	result := config.AnalysisResult{
		LogID:    logConfig.ID,
		FilePath: logConfig.Path,
	}

	// Le fichier existe ?
	fileInfo, err := os.Stat(logConfig.Path)
	if err != nil {
		if os.IsNotExist(err) {
			result.Status = config.StatusFailed
			result.Message = "Fichier introuvable"
			result.ErrorDetails = err.Error()
			return result
		}
		// Autre problème (permissions, etc.)
		result.Status = config.StatusFailed  
		result.Message = "Impossible d'accéder au fichier"
		result.ErrorDetails = err.Error()
		return result
	}

	// C'est un dossier ?
	if fileInfo.IsDir() {
		result.Status = config.StatusFailed
		result.Message = "C'est un répertoire, pas un fichier"
		result.ErrorDetails = "path is a directory"
		return result
	}

	// Fichier vide ?
	if fileInfo.Size() == 0 {
		result.Status = config.StatusOK
		result.Message = "Fichier vide - analyse terminée"
		result.ErrorDetails = ""
		return result
	}

	// Simulation du temps d'analyse
	simulateAnalysis()

	// Tentative de lecture
	_, readErr := os.ReadFile(logConfig.Path)
	if readErr != nil {
		result.Status = config.StatusFailed
		result.Message = "Erreur lecture fichier"
		result.ErrorDetails = readErr.Error()
		return result
	}

	// Toutt va bien
	result.Status = config.StatusOK
	result.Message = fmt.Sprintf("Analyse terminée avec succès - taille: %d bytes", fileInfo.Size())
	result.ErrorDetails = ""
	
	return result
}

// simulateAnalysis simule le temps de traitement
func simulateAnalysis() {
	// Délai aléatoire entre 50 et 200ms
	delay := time.Duration(50+rand.Intn(151)) * time.Millisecond
	time.Sleep(delay)
}