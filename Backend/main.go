package main

import (
	"fmt"
	"forum/api"
	"forum/database"
	"net/http"
)

func main() {
	// Initialiser la base de données
	database.InitDB()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	// Configurer le routeur de l'API
	api.Router()

	// Démarrer la fonction pour diffuser les messages
	go api.HandleMessages()

	// Serveur sur le port 8080
	fmt.Println("Listening in http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
