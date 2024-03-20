package api

import (
	"encoding/json"
	"net/http"
)

// Structure de réponse pour la connexion réussie
type LoginSuccessResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	// SessionID     string    `json:"sessionID"`
	UserID string `json:"userID"`
	// SessionExpiry time.Time `json:"sessionExpiry"`
	HomePage string `json:"homePage"`
	HomeHead string `json:"homeHead"`
}

// Fonction utilitaire pour envoyer des réponses JSON standardisées
func jsonResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  status,
		"message": message,
	})
}

func jsonResponse2(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
