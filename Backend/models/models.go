package models

// import "time"

type Users struct {
	ID            string
	Nickname      string
	Firstname     string
	Lastname      string
	Email         string
	Gender        string
	Age           int
	Password      string
	SessionExpiry string
}

type Register struct {
	Email    string
	Password string
}

type Post struct {
	ID          int    // Identifiant du post
	UserId      int    // La clé étrangère faisant référence à l'utilisateur
	Date        string // Date du post au format ISO 8601 (YYYY-MM-DDTHH:MM:SS)
	PostContent string // Contenu du post
	Category    string // Les catégories du post
}

type Category struct {
	ID           int    // Identifiant de la catégorie
	CategoryName string // Nom de la catégorie
	Posts        []Post // Les posts de la catégorie
}

type Message struct {
	ID         int    // Identifiant du message
	SenderID   int    // La clé étrangère faisant référence à l'utilisateur effectuant le paiement
	ReceiverID int    // La clé étrangère faisant référence à l'utilisateur recevant le paiement
	Content    string // Contenu du message
}

type Comment struct {
	ID             int    // Identifiant du commentaire
	UserID         int    // La clé étrangère faisant référence à l'utilisateur
	PostID         int    // La clé étrangère faisant référence au post
	CommentContent string // Contenu du commentaire
}
