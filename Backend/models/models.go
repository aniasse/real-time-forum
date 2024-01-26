package models

import "time"

type Users struct {
	ID               uint      // Identifiant de l'utilisateur
	Nickname             string    // Rôle de l'utilisateur
	Firstname              string    // Nom de l'utilisateur
	Lastname           string    // Prénom de l'utilisateur
	Email            string    // Adresse e-mail de l'utilisateur
	Gender            string    // URL de la photo de l'utilisateur
	Age					int
	Password         string    // Mot de passe de l'utilisateur
	Session      time.Time // Dernière session de l'utilisateur
}

type Register struct {
	Email    string
	Password string
}

type Post struct {
	ID           uint       // Identifiant du post
	UserID       uint       // La clé étrangère faisant référence à l'utilisateur
	Title        string     // Titre du post
	PostContent  string     // Contenu du post
	Image        string     // Image du post
	Categories   []Category // Les catégories du post
}

type Category struct {
	ID           uint   // Identifiant de la catégorie
	CategoryName string // Nom de la catégorie
	Posts        []Post // Les posts de la catégorie
}

type Message struct {
	ID         uint   // Identifiant du message
	SenderID   uint   // La clé étrangère faisant référence à l'utilisateur effectuant le paiement
	ReceiverID uint   // La clé étrangère faisant référence à l'utilisateur recevant le paiement
	Content    string // Contenu du message
}

type Comment struct {
	ID             uint   // Identifiant du commentaire
	UserID         uint   // La clé étrangère faisant référence à l'utilisateur
	PostID         uint   // La clé étrangère faisant référence au post
	CommentContent string // Contenu du commentaire
}
