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

// Gestionnaire pour la vérification de la session
func HandleCheckSession(w http.ResponseWriter, r *http.Request) {
	// Extraire l'ID de session à partir du cookie
	sessionCookie, err := r.Cookie("sessionID")
	if err != nil {
		jsonResponse(w, http.StatusUnauthorized, "Session not found")
		return
	}

	// Récupérez l'ID de session à partir du cookie
	sessionID := sessionCookie.Value
	var userID string
	var sessionExpiry time.Time

	// Recherche de la session dans la base de données
	err = database.DB.QueryRow("SELECT UserId, SessionExpiry FROM sessions WHERE UserId = ? AND SessionExpiry > CURRENT_TIMESTAMP", sessionID).
		Scan(&userID, &sessionExpiry)

	if err != nil {
		if err == sql.ErrNoRows {
			// Session non trouvée ou expirée
			jsonResponse(w, http.StatusInternalServerError, "Error fetching user ID")
		}
		// Erreur lors de la recherche de la session
		jsonResponse(w, http.StatusInternalServerError, "Error fetching user ID")
	}

	// Session trouvée et valide
	// La session est valide
	jsonResponse2(w, http.StatusOK, map[string]interface{}{
		"Status":  http.StatusOK,
		"Message": "Session is valid",
		"UserID":  userID,
	})
}

// Gestionnaire pour l'inscription des utilisateurs
func HandleRegister(w http.ResponseWriter, r *http.Request) {
	var newUser models.Users

	// Lire les données JSON de la requête
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		jsonResponse(w, http.StatusBadRequest, "Bad Request ")
		fmt.Println("Les données d'inscription sont invalides: ", http.StatusBadRequest)
		return
	}

	// Vérifier si le nickname existe déjà dans la base de données
	err := database.DB.QueryRow("SELECT * FROM users WHERE Nickname = ?", newUser.Nickname).
		Scan(&newUser.ID, &newUser.Nickname, &newUser.Firstname, &newUser.Lastname, &newUser.Email, &newUser.Gender, &newUser.Age, &newUser.Password, &newUser.SessionExpiry)

	if err == nil {
		jsonResponse(w, http.StatusConflict, "Nickname already exist ❌")
		fmt.Println("Cet nickname est déjà enregistré")
		return
	} else if err != sql.ErrNoRows {
		fmt.Println(err) // Journalisation de l'erreur pour le débogage
		jsonResponse(w, http.StatusInternalServerError, "Internal Server Error")
		fmt.Println("Erreur lors de la recherche de l'utilisateur existant: ", err)
		return
	}

	// Vérifier si l'email existe déjà dans la base de données
	err = database.DB.QueryRow("SELECT * FROM users WHERE Email = ?", newUser.Email).
		Scan(&newUser.ID, &newUser.Nickname, &newUser.Firstname, &newUser.Lastname, &newUser.Email, &newUser.Gender, &newUser.Age, &newUser.Password, &newUser.SessionExpiry)

	if err == nil {
		jsonResponse(w, http.StatusConflict, "Email already exist ❌")
		fmt.Println("Cet e-mail est déjà enregistré")
		return
	} else if err != sql.ErrNoRows {
		fmt.Println(err) // Journalisation de l'erreur pour le débogage
		jsonResponse(w, http.StatusInternalServerError, "Internal Server Error")
		fmt.Println("Erreur lors de la recherche de l'utilisateur existant: ", err)
		return
	}

	// Hasher le mot de passe
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err) // Journalisation de l'erreur pour le débogage
		jsonResponse(w, http.StatusInternalServerError, "Internal Server Error")
		fmt.Println("Erreur lors du hachage du mot de passe: ", err)
		return
	}

	uui, errr := uuid.NewV4()
	if errr != nil {
		fmt.Println(err) // Journalisation de l'erreur pour le débogage
		jsonResponse(w, http.StatusInternalServerError, "Internal Server Error")
		fmt.Println("Erreur lors du génération de l'identifiant: ", err)
		return
	}

	newUser.ID = uui.String()
	// Ajouter le nouvel utilisateur à la base de données avec le mot de passe hashé
	_, err = database.DB.Exec("INSERT INTO users (ID, Nickname, Firstname, Lastname, Email, Age, Gender, Password, SessionExpiry) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		newUser.ID, newUser.Nickname, newUser.Firstname, newUser.Lastname, newUser.Email, newUser.Age, newUser.Gender, string(hashedPassword), newUser.SessionExpiry)

	if err != nil {
		fmt.Println(err) // Journalisation de l'erreur pour le débogage
		jsonResponse(w, http.StatusInternalServerError, "Internal Server Error")
		fmt.Println("Erreur lors de l'enregistrement de l'utilisateur: ", err)
		return
	}

	// Enregistrement réussi
	jsonResponse(w, http.StatusCreated, "Registered ✅")
	fmt.Println("Enregistrement réussi: ")

}
