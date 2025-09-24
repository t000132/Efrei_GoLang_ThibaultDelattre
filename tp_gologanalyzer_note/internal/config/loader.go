package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// LoadConfig charge le fichier de config JSON
func LoadConfig(configPath string) ([]LogConfig, error) {
	// Vérif si le fichier existe
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("fichier config introuvable: %s", configPath)
	}

	// Lecture du fichier
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("impossible de lire le config: %w", err)
	}

	// Fichier vide ?
	if len(data) == 0 {
		return nil, fmt.Errorf("fichier config vide")
	}

	// Parsing JSON
	var configs []LogConfig
	if err := json.Unmarshal(data, &configs); err != nil {
		return nil, fmt.Errorf("erreur parsing JSON: %w", err)
	}

	// Au moins une config ?
	if len(configs) == 0 {
		return nil, fmt.Errorf("aucune config trouvée")
	}

	// Validation des champs obligatoires
	for i, config := range configs {
		if config.ID == "" {
			return nil, fmt.Errorf("config %d: ID manquant", i)
		}
		if config.Path == "" {
			return nil, fmt.Errorf("config %d: chemin manquant", i)
		}
		if config.Type == "" {
			return nil, fmt.Errorf("config %d: type manquant", i)
		}
	}

	return configs, nil
}