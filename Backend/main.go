package main

import (
	"forum/database"
	"forum/api"
)

func main() {
	// Initialiser la base de données
	database.InitDB()

	// Démarrer le serveur
	api.Router()
}
