package api

import (
	"fmt"
	"net/http"
	"strconv"
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

	// Servir les fichiers statiques du dossier "Frontend"
	server.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./Frontend"))))

	// Endpoints
	server.HandleFunc("/api/activeSession", handleActiveSession)
	server.HandleFunc("/api/users/", handleUserRequest)
	server.HandleFunc("/api/login", HandleLogin)
	server.HandleFunc("/api/register", HandleRegister)
	// server.HandleFunc("/api/checksession", HandleCheckSession)
}

func handleUserRequest(w http.ResponseWriter, r *http.Request) {
	// Extraire l'ID du chemin de l'URL
	idInt := extractIDFromPath(r.URL.Path)
	id := strconv.Itoa(idInt)

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
