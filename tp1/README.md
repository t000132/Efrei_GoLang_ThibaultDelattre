## Structure des données

Chaque contact contient :
- ID (entier, généré automatiquement)
- Nom (chaîne de caractères)
- Email (chaîne de caractères)

### Mode interactif CLI

Lancer le programme :
```bash
go run main.go
```

Le menu propose 5 options :
1. Ajouter contact
2. Lister contacts
3. Supprimer contact
4. Mettre à jour un contact
5. Quitter

### Mode ligne de commande flags

Ajouter un contact directement :
```bash
go run main.go --ajouter --nom=Tibo --mail=tibo@gmail.fr
```

Les paramètres `--nom` et `--mail` sont obligatoires avec le flag `--ajouter`.

## Notes techniques

- Stockage en mémoire uniquement
- IDs générés de manière séquentielle
- Gestion basique des erreurs de saisie