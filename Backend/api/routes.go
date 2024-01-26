package api

import (
	// "encoding/json"
	// "fmt"
	"fmt"
	"net/http"
	// "strings"

	// "forum/database"
	// "forum/models"
)

func Router() {
	// Créer un serveur HTTP
	server := http.NewServeMux()

	// Configurer CORS (si nécessaire)
	// ...
	fmt.Println("Server started on port 8080")
	// Endpoints pour Users
	server.HandleFunc("/api/users/", handleUserRequest)
	// Nouveaux endpoints pour le login et le register
	server.HandleFunc("/api/login", HandleLogin)
	server.HandleFunc("/api/register", HandleRegister)

	// Démarrer le serveur sur le port 8080
	http.ListenAndServe(":8080", server)

}

func handleUserRequest(w http.ResponseWriter, r *http.Request) {
	// Extraire l'ID du chemin de l'URL
	id := extractIDFromPath(r.URL.Path)

	switch r.Method {
	case http.MethodGet:
		GetUser(w, r, id)
	case http.MethodPost:
		CreateUser(w, r)
	case http.MethodPut:
		UpdateUser(w, r, id)
	case http.MethodDelete:
		DeleteUser(w, r, id)
	default:
		fmt.Printf("Method %s not allowed\n", r.Method)
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
}