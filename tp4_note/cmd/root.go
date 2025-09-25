package cmd

import (
	"fmt"
	"os"

	"mini-crm/internal/config"
	"mini-crm/internal/store"

	"github.com/spf13/cobra"
)

var (
	// Variables globales pour l'injection de dépendances
	storer  store.Storer
	cfgFile string
)

// rootCmd représente la commande de base
var rootCmd = &cobra.Command{
	Use:   "mini-crm",
	Short: "Un gestionnaire de contacts simple et efficace",
	Long: `Mini-CRM CLI est un gestionnaire de contacts en ligne de commande
écrit en Go. Il supporte SQLite avec GORM (recommandé), Fichier JSON et Stockage en mémoire (temporaire)`,
}

// Execute lance l'applciation
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Initialiser la configuration au démarrage
	cobra.OnInitialize(initConfig)

	// Flag global pour le fichier de configuration
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yaml",
		"fichier de configuration (défaut: config.yaml)")
}

// initConfig lit le fichier de configuration et initialise le storer
func initConfig() {
	// Charger la configuration
	cfg, err := config.LoadConfig(cfgFile)
	if err != nil {
		fmt.Printf("Erreur chargement configuration: %v\n", err)
		fmt.Println("Utilisation du stockage en mémoire par défaut")
		storer = store.NewMemoryStore()
		return
	}

	// Créeer le storer basé sur la configuration
	storer, err = cfg.CreateStore()
	if err != nil {
		fmt.Printf("Erreur initialisation stockage: %v\n", err)
		fmt.Println("Utilisation du stockage en mémoire par défaut")
		storer = store.NewMemoryStore()
		return
	}

	fmt.Printf("Stockage initialisé: %s\n", cfg.Storage.Type)
}
