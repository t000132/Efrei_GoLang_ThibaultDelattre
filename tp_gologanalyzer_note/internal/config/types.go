package config

// Config d'un fichier de log depuis le JSON
type LogConfig struct {
	ID   string `json:"id"`
	Path string `json:"path"`
	Type string `json:"type"`
}

// Résultat après analyse d'un log
type AnalysisResult struct {
	LogID        string `json:"log_id"`
	FilePath     string `json:"file_path"`
	Status       string `json:"status"`
	Message      string `json:"message"`
	ErrorDetails string `json:"error_details"`
}

// Status possibles
const (
	StatusOK     = "OK"
	StatusFailed = "FAILED"
)