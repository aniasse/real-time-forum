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
	// serveur HTTP
	server := http.NewServeMux()

	fmt.Println("Server started on port 8080")
	// Endpoints 
	server.HandleFunc("/api/users/", handleUserRequest)
	server.HandleFunc("/api/login", HandleLogin)
	server.HandleFunc("/api/register", HandleRegister)

	//serveur sur le port 8080
	http.ListenAndServe(":8080", server)

}

func handleUserRequest(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
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

func enableCors(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
}
