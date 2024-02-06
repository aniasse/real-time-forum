package main

import (
	"fmt"
	"forum/api"
	"forum/database"
	"net/http"
)

// func main() {
// 	// Initialiser la base de données
// 	database.InitDB()

// 	// Configurer le routeur de l'API
// 	api.Router()
// }

func main() {
	// Initialiser la base de données
	database.InitDB()

	// Configurer le routeur de l'API
	api.Router()

	// Serveur sur le port 8080
	fmt.Println("Listening in http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
