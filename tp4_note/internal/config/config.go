package config

import (
	"fmt"
	"mini-crm/internal/store"

	"github.com/spf13/viper"
)

type Config struct {
	Storage StorageConfig `mapstructure:"storage"`
}

type StorageConfig struct {
	Type string `mapstructure:"type"`
	Path string `mapstructure:"path"`
}

// LoadConfig charge la configuration depuis le fichier config.yaml
func LoadConfig(configPath string) (*Config, error) {
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	// Valeurs par défaut
	viper.SetDefault("storage.type", "memory")
	viper.SetDefault("storage.path", "contacts.db")

	// Lire le fichier de configuration
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("erreur lecture fichier config: %v", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("erreur parsing config: %v", err)
	}

	return &config, nil
}

// CreateStore crée une instance de Storer basée sur la configuration
func (c *Config) CreateStore() (store.Storer, error) {
	switch c.Storage.Type {
	case "memory":
		return store.NewMemoryStore(), nil
	case "json":
		return store.NewJSONStore(c.Storage.Path)
	case "gorm":
		return store.NewGORMStore(c.Storage.Path)
	default:
		return nil, fmt.Errorf("type de stockage non supporté: %s", c.Storage.Type)
	}
}
