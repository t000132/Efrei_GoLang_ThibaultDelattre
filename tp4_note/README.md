# Mini-CRM CLI

Un gestionnaire de contacts simple et efficace en ligne de commande, écrit en Go.

## Fonctionnalités

- **Gestion complète des contacts (CRUD)** : Ajouter, Lister, Mettre à jour et Supprimer
- **Interface en ligne de commande** avec Cobra
- **Configuration externe** avec Viper
- **Persistance multi-backend** :
  - **SQLite avec GORM** (recommandé)
  - **Fichier JSON** (simple et lisible)
  - **Stockage en mémoire** (temporaire pour tests)

## Architecture

Le projet suit les bonnes pratiques Go avec :
- **Architecture en packages découplés**
- **Injection de dépendances via interfaces**
- **Pattern Repository** pour l'abstraction du stockage
- **Configuration externalisée**

```
mini-crm/
├── cmd/                    # Commandes Cobra
│   ├── root.go            # Commande racine et injection dépendances  
│   ├── add.go             # Ajouter un contact
│   ├── list.go            # Lister les contacts
│   ├── update.go          # Mettre à jour un contact
│   └── delete.go          # Supprimer un contact
├── internal/
│   ├── models/            # Modèles de données
│   │   └── contact.go     # Struct Contact avec annotations GORM
│   ├── store/             # Couche de persistance
│   │   ├── interface.go   # Interface Storer (abstraction)
│   │   ├── memory.go      # Implémentation mémoire
│   │   ├── json.go        # Implémentation JSON
│   │   └── gorm.go        # Implémentation GORM/SQLite
│   └── config/            # Configuration
│       └── config.go      # Chargement config avec Viper
├── config.yaml            # Fichier de configuration
├── go.mod                 # Dépendances Go
├── main.go                # Point d'entrée
└── README.md              # Documentation
```

## Installation et Utilisation

### Prérequis
- Go 1.21 ou supérieur

### Installation des dépendances
```bash
go mod download
```

### Configuration

Modifiez le fichier `config.yaml` pour choisir le type de stockage :

```yaml
# SQLite avec GORM (recommandé)
storage:
  type: "gorm"
  path: "contacts.db"

# Ou stockage JSON
# storage:
#   type: "json"  
#   path: "contacts.json"

# Ou stockage en mémoire (temporaire)
# storage:
#   type: "memory"
#   path: ""
```

### Compilation
```bash
go build -o mini-crm main.go
```

### Utilisation

#### Aide générale
```bash
./mini-crm --help
```

#### Ajouter un contact
```bash
./mini-crm add --name "Jean Dupont" --email "jean@example.com" --phone "0123456789" --company "ACME Corp"
./mini-crm add -n "Marie Martin" -e "marie@test.fr"
```

#### Lister les contacts
```bash
./mini-crm list                # Tous les contacts
./mini-crm list --id 5         # Contact spécifique par ID
```

#### Mettre à jour un contact
```bash
./mini-crm update --id 5 --name "Thibault Dt"
./mini-crm update -i 3 -e "nouveau@email.com" -p "0933443344"
```

#### Supprimer un contact
```bash
./mini-crm delete --id 5       # Avec confirmation
./mini-crm delete -i 3 --force # Sans confirmation
```

## Backends de Stockage

### SQLite avec GORM (Recommandé)
- Persistence robuste en base de données
- Migrations automatiques
- Contraintes d'unicité
- Fichier `contacts.db` créé automatiquement

### JSON
- Stockage simple en fichier texte
- Lisible et éditable manuellement
- Idéal pour petites quantités de données

### Mémoire
- Stockage temporaire (perdu au redémarrage)
- Très rapide
- Idéal pour les tests et développement

## Basculer entre les Backends

Il suffit de modifier le fichier `config.yaml` et relancer l'application :

```yaml
# Passer de SQLite à JSON
storage:
  type: "json"
  path: "mes-contacts.json"
```

Aucune recompilation nécessaire !

## Tests

Le projet inclut une architecture testable avec injection de dépendances. Chaque backend implémente la même interface `Storer`.

## Dépendances

- `github.com/spf13/cobra` - CLI framework
- `github.com/spf13/viper` - Configuration management
- `gorm.io/gorm` - ORM pour Go
- `gorm.io/driver/sqlite` - Driver SQLite pour GORM

## Concepts Go Illustrés

- **Interfaces** pour l'abstraction et l'injection de dépendances
- **Packages** pour l'organisation modulaire  
- **Struct tags** pour GORM et JSON
- **Error handling** avec messages explicites
- **Concurrency safety** avec mutex dans MemoryStore
- **File I/O** et manipulation JSON
- **CLI development** avec Cobra
- **Configuration management** avec Viper

## Exemple d'utilisation complète

```bash
# Configurer pour utiliser SQLite
./mini-crm --config config.yaml

# Ajouter quelques contacts
./mini-crm add -n "Alice Durand" -e "alice@test.fr" -p "0123456789" -c "TechCorp"
./mini-crm add -n "Bob Martin" -e "bob@example.com" -p "0987654321"

# Lister tous les contacts
./mini-crm list

# Mettre à jour un contact
./mini-crm update --id 1 --company "New TechCorp"

# Supprimer un contact  
./mini-crm delete --id 2

# Lister à nouveau
./mini-crm list
```