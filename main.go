
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
//test
func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nMini-CRM")
		fmt.Println("1. Ajouter contact")
		fmt.Println("2. Lister contacts")
		fmt.Println("3. Supprimer contact")
		fmt.Println("4. Quitter")
		fmt.Print("Votre choix : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			ajouterContact(reader)
		case "2":
			listerContacts()
		case "3":
			supprimerContact(reader)
		case "4":
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
