package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"forum/database"
	"forum/models"
	"time"

	"github.com/gofrs/uuid"
)

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

// Gestionnaire pour la connexion des utilisateurs
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var user models.Users
	var login models.Register

	// Lire les données JSON de la requête
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		jsonResponse(w, http.StatusBadRequest, "Les données d'identification sont invalides")
		fmt.Println("Les données d'identification sont invalides: ", http.StatusBadRequest)
		return
	}

	// Recherche de l'utilisateur dans la base de données
	err := database.DB.QueryRow("SELECT * FROM users WHERE Email = ?", login.Email).
		Scan(&user.ID, &user.Nickname, &user.Firstname, &user.Lastname, &user.Email, &user.Gender, &user.Age, &user.Password, &user.SessionExpiry)

	if err != nil {
		if err == sql.ErrNoRows {
			jsonResponse(w, http.StatusUnauthorized, "Identifiants incorrects")
			return
		} else {
			fmt.Println(err) // Journalisation de l'erreur pour le débogage
			jsonResponse(w, http.StatusInternalServerError, "Erreur lors de la recherche de l'utilisateur")
			return
		}
	}

	// Vérification du mot de passe
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	fmt.Println("Vérification du mot de passe: ", user.Password, login.Password)
	if err != nil {
		jsonResponse(w, http.StatusUnauthorized, "Identifiants incorrects")
		return
	}

	// Connexion réussie
	jsonResponse(w, http.StatusOK, "Connexion réussie")
	fmt.Println("connexion réussie: ")

	// Session de l'utilisateur

	sessionID, err1 := uuid.NewV4()
	if err1 != nil {
		// ...
	}

	// Calcul de l'heure d'expiration de la session (15 minutes plus tard)
	sessionExpiry := time.Now().Add(15 * time.Minute)

	// Mise à jour de l'identifiant de session et de l'heure d'expiration dans la base de données
	_, err = database.DB.Exec("UPDATE users SET SessionExpiry = ? WHERE Id = ?", sessionExpiry, user.ID)
	if err != nil {
		// ...
	}

	// Insertion de la session dans la table sessions
	_, err = database.DB.Exec("INSERT INTO sessions (ID, UserId, SessionExpiry) VALUES (?, ?, ?)", sessionID.String(), user.ID, sessionExpiry)
	if err != nil {
		// ...
	}
}

// Gestionnaire pour l'inscription des utilisateurs
func HandleRegister(w http.ResponseWriter, r *http.Request) {
	var newUser models.Users

	// Lire les données JSON de la requête
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		jsonResponse(w, http.StatusBadRequest, "Les données d'inscription sont invalides")
		fmt.Println("Les données d'inscription sont invalides: ", http.StatusBadRequest)
		return
	}

	// Vérifier si l'utilisateur existe déjà dans la base de données
	err := database.DB.QueryRow("SELECT * FROM users WHERE Email = ?", newUser.Email).
		Scan(&newUser.ID, &newUser.Nickname, &newUser.Firstname, &newUser.Lastname, &newUser.Email, &newUser.Gender, &newUser.Age, &newUser.Password, &newUser.SessionExpiry)

	if err == nil {
		jsonResponse(w, http.StatusConflict, "Cet e-mail est déjà enregistré")
		fmt.Println("Cet e-mail est déjà enregistré")
		return
	} else if err != sql.ErrNoRows {
		fmt.Println(err) // Journalisation de l'erreur pour le débogage
		jsonResponse(w, http.StatusInternalServerError, "Erreur lors de la recherche de l'utilisateur existant")
		fmt.Println("Erreur lors de la recherche de l'utilisateur existant: ", err)
		return
	}

	// Hasher le mot de passe
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err) // Journalisation de l'erreur pour le débogage
		jsonResponse(w, http.StatusInternalServerError, "Erreur lors du hachage du mot de passe")
		fmt.Println("Erreur lors du hachage du mot de passe: ", err)
		return
	}

	uui, errr := uuid.NewV4()
	if errr != nil {
		fmt.Println(err) // Journalisation de l'erreur pour le débogage
		jsonResponse(w, http.StatusInternalServerError, "Erreur lors du génération de l'identifiant")
		fmt.Println("Erreur lors du génération de l'identifiant: ", err)
		return
	}

	newUser.ID = uui.String()
	// Ajouter le nouvel utilisateur à la base de données avec le mot de passe hashé
	_, err = database.DB.Exec("INSERT INTO users (ID, Nickname, Firstname, Lastname, Email, Age, Gender, Password, SessionExpiry) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		newUser.ID, newUser.Nickname, newUser.Firstname, newUser.Lastname, newUser.Email, newUser.Age, newUser.Gender, string(hashedPassword), newUser.SessionExpiry)

	if err != nil {
		fmt.Println(err) // Journalisation de l'erreur pour le débogage
		jsonResponse(w, http.StatusInternalServerError, "Erreur lors de l'enregistrement de l'utilisateur")
		fmt.Println("Erreur lors de l'enregistrement de l'utilisateur: ", err)
		return
	}

	// Enregistrement réussi
	jsonResponse(w, http.StatusCreated, "Enregistrement réussi")
	fmt.Println("Enregistrement réussi: ")

}
