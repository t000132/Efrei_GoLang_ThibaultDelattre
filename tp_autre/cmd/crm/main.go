package main

import (
	"tp3/internal/app"
)

// la fonction main orchestre l'application. Elle depend de l'interface storer
//  pas de MemoryStore directement, c'est ça l'injection de dépendances ! 
func main() {
	application := app.NewApp()
	application.Run()
}