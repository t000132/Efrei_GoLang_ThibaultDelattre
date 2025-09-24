# TP2 - Système de Notifications et Logging

### Types de notifications supportés
- **Email** : Envoi vers une adresse email personnalisée
- **SMS** : Envoi vers un numéro de téléphone avec validation
- **Push** : Notification push vers un appareil spécifique

### Validation SMS
- Les numéros de téléphone doivent commencer par "06"
- Erreur explicite si le numéro est invalide
- Les envois échoués n'arrêtent pas le programme

### Système d'archivage
- Enregistrement automatique des notifications réussies
- Stockage en mémoire avec timestamp
- Affichage de l'historique complet à la fin

## Structure du code

### Interfaces
- `Notifier` : Interface commune pour tous les types de notifications
- `Storer` : Interface pour le système d'archivage

### Structures
- `EmailNotifier` : Gestion des emails
- `SMSNotifier` : Gestion des SMS avec validation
- `PushNotifier` : Gestion des notifications push
- `MemoryStorer` : Stockage en mémoire
- `NotificationRecord` : Enregistrement d'une notification

## Utilisation

```bash
go run main.go
```

Le programme démarre avec un menu interactif :
1. **Envoyer Email** - Saisir adresse et message
2. **Envoyer SMS** - Saisir numéro (doit commencer par 06) et message  
3. **Envoyer Push** - Saisir ID appareil et message
4. **Voir historique** - Afficher les notifications archivées
5. **Quitter** - Fermer le programme et afficher l'historique final

## Tests

Lancer les tests unitaires :
```bash
go test -v
```

## Concepts Go utilisés

- Interface CLI avec `bufio` pour la saisie utilisateur
- Interfaces et polymorphisme (`Notifier`, `Storer`)
- Structures et méthodes avec récepteur
- Slices pour les collections
- Gestion d'erreurs avec `error` et validation
- Package `time` pour les timestamps
- Boucles et structures de contrôle (`for`, `switch`)
- Saisie interactive avec `os.Stdin`