package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

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

// Expressions régulières pour la validation
var regexMap = map[string]*regexp.Regexp{
	"nickname":  regexp.MustCompile(`^[a-zA-Z0-9]{4,8}$`),
	"firstName": regexp.MustCompile(`^[a-zA-Z]{2,}$`),
	"lastName":  regexp.MustCompile(`^[a-zA-Z]{2,}$`),
	"age":       regexp.MustCompile(`^(1[4-9]|[2-5][0-9]|60)$`),
	"email":     regexp.MustCompile(`^[\w-]+(\.[\w-]+)*@([\w-]+\.)+[a-zA-Z]{2,7}$`),
	"password":  regexp.MustCompile(`^[!-~]{4,}$`),
	"gender":    regexp.MustCompile(`^(Male|Female)$`),
}

// Fonction pour valider les données utilisateur
func validateUserData(user models.Users) error {
	if !regexMap["nickname"].MatchString(user.Nickname) {
		return fmt.Errorf("invalid nickname format ❌")
	}
	if !regexMap["firstName"].MatchString(user.Firstname) {
		return fmt.Errorf("invalid first name format ❌")
	}
	if !regexMap["lastName"].MatchString(user.Lastname) {
		return fmt.Errorf("invalid last name format ❌")
	}
	if !regexMap["age"].MatchString(user.Age) {
		return fmt.Errorf("invalid age format ❌")
	}
	if !regexMap["email"].MatchString(user.Email) {
		return fmt.Errorf("invalid email format ❌")
	}
	if !regexMap["password"].MatchString(user.Password) {
		return fmt.Errorf("invalid password format ❌")
	}
	if !regexMap["gender"].MatchString(user.Gender) {
		return fmt.Errorf("invalid gender value ❌")
	}
	return nil
}

// Gestionnaire pour l'inscription des utilisateurs
func handleRegister(w http.ResponseWriter, r *http.Request) {
	var newUser models.Users

	if r.Method != http.MethodPost {
		jsonResponse(w, http.StatusMethodNotAllowed, "Method Not Allow")
		return
	}
	// Lire les données JSON de la requête
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		jsonResponse(w, http.StatusBadRequest, "Bad Request")
		fmt.Println("Les données d'inscription sont invalides: ", http.StatusBadRequest)
		return
	}

	// Valider les données de l'utilisateur
	if err := validateUserData(newUser); err != nil {
		jsonResponse(w, http.StatusBadRequest, err.Error())
		fmt.Println("Erreur de validation des données d'utilisateur:", err)
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

func VerifyPost(post newPosts) (string, bool) {

	validCategories := []string{"News", "Tech", "Computing", "Sport", "Gaming"}
	catIsValid := false

	for _, cat := range validCategories {
		if cat == post.Category {
			catIsValid = true
			break
		}
	}

	if !catIsValid {
		return "Invalid Category ❌", false
	}

	if strings.TrimSpace(post.PostContent) == "" {
		return "Empty Post ❌", false
	}

	return "Posted ✅", true
}

type newPosts struct {
	UserId      string
	Category    string
	PostContent string
}

// Creating Post
func handleCreatingPost(w http.ResponseWriter, r *http.Request) {

	var newPost newPosts
	if r.Method != http.MethodPost {
		jsonResponse(w, http.StatusMethodNotAllowed, "Method Not Allow")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&newPost); err != nil {
		jsonResponse(w, http.StatusBadRequest, "Bad Request")
		fmt.Println("Les données d'inscription sont invalides: ", http.StatusBadRequest)
		return
	}

	message, isValid := VerifyPost(newPost)

	if !isValid {
		jsonResponse(w, http.StatusNonAuthoritativeInfo, message)
		fmt.Println("Les données du post sont invalides: ", http.StatusNonAuthoritativeInfo)
		return
	}

	date := time.Now().Format(time.RFC3339)

	_, err := database.DB.Exec("INSERT INTO posts (UserId, Category, Content, Date) VALUES (?, ?, ?, ?)",
		newPost.UserId, newPost.Category, newPost.PostContent, date)

	if err != nil {
		fmt.Println(err)
		jsonResponse(w, http.StatusInternalServerError, "Internal Server Error")
		fmt.Println("Erreur lors de l'enregistrement de l'utilisateur: ", err)
		return
	}

	jsonResponse(w, http.StatusCreated, "Posted ✅")
	fmt.Println("Post réussi: ")

}

func handleGetPosts(w http.ResponseWriter, r *http.Request) {
	// Requête avec une jointure pour inclure le Nickname du User

	rows, err := database.DB.Query(`
			SELECT 
				posts.Id, posts.UserId, posts.Category, posts.Content, posts.Date,
				users.Nickname
			FROM 
				posts
			LEFT JOIN 
				users ON posts.UserId = users.Id
		`)
	if err != nil {
		fmt.Println("error 1")
		jsonResponse(w, http.StatusInternalServerError, "Internal Server Error")
		fmt.Println("Erreur lors de la recuperation de posts: ", err)
		return
	}
	defer rows.Close()

	var posts []PostWithUser
	for rows.Next() {
		var post PostWithUser
		err := rows.Scan(&post.ID, &post.UserID, &post.Category, &post.Content, &post.Date, &post.Nickname)
		if err != nil {
			fmt.Println("error 2")
			jsonResponse(w, http.StatusInternalServerError, "Internal Server Error")
			fmt.Println("Erreur lors de la recuperation de posts: ", err)
			return
		}
		posts = append(posts, post)
	}

	jsonResponse2(w, http.StatusOK, posts)
}

type PostID struct {
	PostId string `json:"PostId"`
}

// Structure de commentaire
type Comment struct {
	Content  string `json:"content"`
	Username string `json:"username"`
}

func handleGetComments(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		jsonResponse(w, http.StatusMethodNotAllowed, "Method Not Allow")
		return
	}

	// Décoder le PostId reçu
	var postId PostID
	if err := json.NewDecoder(r.Body).Decode(&postId); err != nil {
		jsonResponse(w, http.StatusBadRequest, "Bad Request")
		fmt.Println("Les données d'inscription sont invalides: ", http.StatusBadRequest)
		return
	}

	query := `
		SELECT comments.Content, users.Nickname
		FROM comments
		JOIN users ON comments.UserId = users.Id
		WHERE comments.PostId = $1
	`

	rows, err := database.DB.Query(query, postId.PostId)
	if err != nil {
		fmt.Println("error 1")
		fmt.Println("Erreur lors de la récupération des commentaires:", err)
		jsonResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	defer rows.Close()

	comments := []Comment{}
	for rows.Next() {
		var comment Comment
		err := rows.Scan(&comment.Content, &comment.Username)
		if err != nil {
			fmt.Println("error 2")
			fmt.Println("Erreur lors de la lecture d'un commentaire:", err)
			jsonResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		comments = append(comments, comment)
	}

	jsonResponse2(w, http.StatusOK, comments)
}

// Structure de commentaire pour stockage en base de données
type Commentary struct {
	UserID  string `json:"UserId"`
	PostID  string `json:"PostId"`
	Content string `json:"Content"`
}

// Fonction pour traiter les commentaires envoyés depuis le frontend
func handleComment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("enter")
	// Vérifier que la méthode HTTP est POST
	if r.Method != http.MethodPost {
		jsonResponse(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	// Déclarer une variable pour stocker les données du commentaire
	var comment Commentary

	// Décoder les données du commentaire depuis le corps de la requête
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, "Bad Request")
		fmt.Println("Erreur lors de la lecture des données du commentaire:", err)
		return
	}

	// Expression régulière pour un entier positif
	re := regexp.MustCompile(`^\d+$`)

	// Vérifier si la chaîne correspond à l'expression régulière
	isInteger := re.MatchString(comment.PostID)

	var count int
	database.DB.QueryRow("SELECT COUNT(*) FROM posts WHERE id = $1", comment.PostID).Scan(&count)

	// Vérifier que le contenu du commentaire n'est pas vide ou ne contient que des espaces
	if strings.TrimSpace(comment.Content) == "" || !isInteger || count != 1 {
		jsonResponse2(w, http.StatusBadRequest, "Bad Request")
		fmt.Println("Empty com")
		return
	}

	// Insérer le commentaire dans la base de données
	_, err = database.DB.Exec(`
		INSERT INTO comments (UserId, PostId, Content)
		VALUES ($1, $2, $3)
	`, comment.UserID, comment.PostID, comment.Content)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, "Internal Server Error")
		fmt.Println("Erreur lors de l'insertion du commentaire en base de données:", err)
		return
	}

	// Répondre avec un message de succès
	jsonResponse(w, http.StatusCreated, "Commented ✅.")
}

type UserNickname struct {
	Nickname string `json:"nickname"`
}

type User struct {
	UserId string `json:"UserId"`
}

func handleGetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getusers")
	if r.Method != http.MethodPost {
		jsonResponse(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		jsonResponse(w, http.StatusBadRequest, "Bad Request")
		fmt.Println("Erreur lors de la réception des données utilisateur:", err)
		return
	}

	// Récupérer tous les nicknames sauf celui de l'utilisateur spécifié
	query := `
		SELECT Nickname
		FROM users
		WHERE Id != $1
	`
	rows, err := database.DB.Query(query, user.UserId)
	if err != nil {
		fmt.Println("Erreur lors de la récupération des utilisateurs:", err)
		jsonResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	defer rows.Close()

	var userNicknames []UserNickname
	for rows.Next() {
		var nickname string
		err := rows.Scan(&nickname)
		if err != nil {
			fmt.Println("Erreur lors de la lecture d'un nickname:", err)
			jsonResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		userNicknames = append(userNicknames, UserNickname{Nickname: nickname})
	}

	jsonResponse2(w, http.StatusOK, userNicknames)
}

type Message struct {
	From string `json:"from"`
	Text string `json:"text"`
	Date string `json:"date"`
}

func handleGettingDiscus(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handleGettingDiscus")

	if r.Method != http.MethodPost {
		jsonResponse(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	var requestBody struct {
		UserId           string `json:"UserId"`
		ReceiverNickname string `json:"ReceiverNickname"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		jsonResponse(w, http.StatusBadRequest, "Bad Request")
		fmt.Println("Erreur lors du decodage", http.StatusBadRequest)
		return
	}

	userId := requestBody.UserId
	receiverNickname := requestBody.ReceiverNickname

	// Vérifier si le ReceiverNickname existe dans la table users
	var receiverID string
	err := database.DB.QueryRow("SELECT Id FROM users WHERE Nickname = $1", receiverNickname).Scan(&receiverID)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, "ReceiverNickname does not exist")
		fmt.Println("ReceiverNickname does not exist:", err)
		return
	}

	// Vérifier si le ReceiverNickname est différent du Nickname du UserId
	if receiverID == userId {
		jsonResponse(w, http.StatusBadRequest, "ReceiverNickname cannot be the same as UserId")
		fmt.Println("ReceiverNickname cannot be the same as UserId")
		return
	}

	// Query pour récupérer les messages de la base de données
	query := `
		SELECT SenderId, ReceiverId, Content, Date
		FROM messages
		WHERE (SenderId = $1 AND ReceiverId = $2)
			OR (ReceiverId = $1 AND SenderId = $2)
		ORDER BY Date
	`
	rows, err := database.DB.Query(query, userId, receiverID)
	if err != nil {
		fmt.Println("Erreur lors de la récupération des messages:", err)
		jsonResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var senderID, receiverID, content, date string
		err := rows.Scan(&senderID, &receiverID, &content, &date)
		if err != nil {
			fmt.Println("Erreur lors de l'insertion des données dans les variables:", err)
			jsonResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		// Déterminer la provenance du message
		from := "notUser"
		if senderID == userId {
			from = "user"
		}

		// Créer un objet Message
		msg := Message{
			From: from,
			Text: content,
			Date: date,
		}

		messages = append(messages, msg)
	}

	jsonResponse2(w, http.StatusOK, messages)
}
