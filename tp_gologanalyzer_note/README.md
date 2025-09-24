## Fonctionnalités
- Goroutines + WaitGroup + Channels
- Erreurs personnalisées (FileNotFoundError, ParseError)
- CLI avec flags --config/-c et --output/-o
- Import/Export JSON

## Utilisation

### Installation
```bash
go mod download
```

### Tests
```bash
# Analyse basique
go run main.go analyze -c config.json

# Avec export
go run main.go analyze -c config.json -o rapport.json

# Aide
go run main.go analyze --help
```

## Export JSON
```json
[
  {
    "log_id": "web-server-1",
    "file_path": "test_logs/access.log",
    "status": "OK",
    "message": "Analyse terminée avec succès - taille: 154 bytes",
    "error_details": ""
  }
]
```

## Bonus
- Création auto des dossiers d'export
- Horodatage des fichiers (250924_report.json)
- L'ordre des résultats varie (concurrence normale)
