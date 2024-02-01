package api

import (
	"fmt"
	"net/http"
)

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCORS(w)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Router() {
	// Serveur HTTP
	server := http.NewServeMux()

	// Ajouter le middleware CORS
	http.Handle("/", corsMiddleware(server))

	// Endpoints
	server.HandleFunc("/api/users/", handleUserRequest)
	server.HandleFunc("/api/login", HandleLogin)
	server.HandleFunc("/api/register", HandleRegister)
	server.HandleFunc("/api/checksession", HandleCheckSession)

	fmt.Println("Server started on port 8080")

	// Serveur sur le port 8080
	http.ListenAndServe(":8080", nil)
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
