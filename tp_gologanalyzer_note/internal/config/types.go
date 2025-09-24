package config

// Configuration d'un fichier de log provenant de l'entrée JSON
type LogConfig struct {
	ID   string `json:"id"`
	Path string `json:"path"`
	Type string `json:"type"`
}

// Résultat d'analyse d'un fichier de log
type AnalysisResult struct {
	LogID        string `json:"log_id"`
	FilePath     string `json:"file_path"`
	Status       string `json:"status"`
	Message      string `json:"message"`
	ErrorDetails string `json:"error_details"`
}

// Constantes des status pour les résultats d'analyse
const (
	StatusOK     = "OK"
	StatusFailed = "FAILED"
)