package database

import (
	"database/sql"
	"log"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3" // Pilote SQLite

	"forum/models"
)

var DB *sql.DB

func InitDB() {
	// Spécifiez le chemin absolu pour le fichier de base de données
	dbPath, err := filepath.Abs("database.db")
	if err != nil {
		log.Fatal(err)
	}

	// Ouvrir la connexion à la base de données SQLite
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	// Créer les tables si elles n'existent pas déjà
	createTables(db)

	// Affectez la variable globale DB avec la connexion à la base de données
	DB = db
}

func createTables(db *sql.DB) {
	// Créer les tables si elles n'existent pas déjà
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			Id TEXT PRIMARY KEY ,
			Nickname TEXT NOT NULL,
			Firstname TEXT NOT NULL,
			Lastname TEXT NOT NULL,
			Email TEXT NOT NULL,
			Gender TEXT NOT NULL,
			Age TEXT NOT NULL,
			Password VARCHAR(254) NOT NULL,
			SessionExpiry DATETIME NOT NULL
		);

		CREATE TABLE IF NOT EXISTS posts (
			Id INTEGER PRIMARY KEY AUTOINCREMENT,			
			UserId TEXT NOT NULL,
			Category TEXT NOT NULL,
            Content TEXT NOT NULL,
            Date TEXT NOT NULL,
			FOREIGN KEY(UserId) REFERENCES users(Id)
			ON DELETE CASCADE
			ON UPDATE CASCADE
		);

		CREATE TABLE IF NOT EXISTS comments (
			Id INTEGER PRIMARY KEY AUTOINCREMENT,            
			UserId TEXT NOT NULL,
            PostId TEXT NOT NULL,
            Content TEXT NOT NULL,
			FOREIGN KEY(UserId) REFERENCES users(Id)
			FOREIGN KEY(PostId) REFERENCES posts(Id)
			ON DELETE CASCADE
			ON UPDATE CASCADE
		);

		CREATE TABLE IF NOT EXISTS sessions (
			Id INTEGER PRIMARY KEY AUTOINCREMENT,            
			UserId TEXT NOT NULL,
			SessionExpiry DATETIME NOT NULL,
			FOREIGN KEY(UserId) REFERENCES users(Id) ON DELETE CASCADE ON UPDATE CASCADE
		);

		CREATE TABLE IF NOT EXISTS messages (
			Id INTEGER PRIMARY KEY AUTOINCREMENT,
			SenderNickname TEXT NOT NULL,
			ReceiverNickname TEXT NOT NULL,
			Content TEXT NOT NULL,
			Type TEXT NOT NULL,
			Date DATETIME NOT NULL
		);

	`)

	if err != nil {
		log.Fatal(err)
	}
}

// GetUserByID récupère un utilisateur par son ID depuis la base de données
func GetUserByID(id string) (*models.Users, error) {
	var user models.Users

	query := "SELECT Id, Nickname, Firstname, Lastname, Email, Gender, Age, Password, SessionExpiry FROM users WHERE Id = ?"
	err := DB.QueryRow(query, id).
		Scan(&user.ID, &user.Nickname, &user.Firstname, &user.Lastname, &user.Email, &user.Gender, &user.Age, &user.Password, &user.SessionExpiry)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetAllUsers récupère tous les utilisateurs depuis la base de données
func GetAllUsers() ([]*models.Users, error) {
	var users []*models.Users

	query := "SELECT Id, Nickname, Firstname, Lastname, Email, Gender, Age, Password, SessionExpiry FROM users"
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.Users
		err := rows.Scan(&user.ID, &user.Nickname, &user.Firstname, &user.Lastname, &user.Email, &user.Gender, &user.Age, &user.Password, &user.SessionExpiry)

		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

// CreateUser crée un nouvel utilisateur dans la base de données
func CreateUser(user *models.Users) (int64, error) {
	query := "INSERT INTO users (Nickname, Firstname, Lastname, Email, Gender, Age, Password, SessionExpiry) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	result, err := DB.Exec(query, user.Nickname, user.Firstname, user.Lastname, user.Email, user.Gender, user.Age, user.Password, user.SessionExpiry)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

// UpdateUser met à jour les informations d'un utilisateur dans la base de données
func UpdateUser(user *models.Users) error {
	query := "UPDATE users SET Nickname=?, Firstname=?, Lastname=?, Email=?, Gender=?, Age=?, Password=?, SessionExpiry=? WHERE Id=?"

	_, err := DB.Exec(query, user.Nickname, user.Firstname, user.Lastname, user.Email, user.Gender, user.Age, user.Password, user.SessionExpiry, user.ID)
	return err
}

// DeleteUser supprime un utilisateur de la base de données
func DeleteUser(id string) error {
	query := "DELETE FROM users WHERE Id = ?"
	_, err := DB.Exec(query, id)
	return err
}

// func getUserByID(userID string) (*models.Users, error) {
// 	var user models.Users
// 	query := "SELECT * FROM users WHERE ID = ?"
// 	err := DB.QueryRow(query, userID).Scan(
// 		&user.ID,
// 		&user.Nickname,
// 		&user.Firstname,
// 		&user.Lastname,
// 		&user.Email,
// 		&user.Gender,
// 		&user.Age,
// 		&user.Password,
// 		&user.SessionExpiry,
// 	)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &user, nil
// }
