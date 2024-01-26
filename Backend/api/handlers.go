package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"forum/database"
	"forum/models"
)

// Fonctions CRUD pour Users
func GetUser(w http.ResponseWriter, r *http.Request, id int) {

	user, err := database.GetUserByID(id)
	if err != nil {
		fmt.Println("Error getting user: ", err)
		http.Error(w, "Utilisateur non trouvé", http.StatusNotFound)
		return
	}

	jsonResponse2(w, http.StatusOK, user)
	fmt.Println("status: ", http.StatusOK)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := database.GetAllUsers()
	if err != nil {
		fmt.Println("Error getting users: ", err)
		http.Error(w, "Erreur lors de la récupération des utilisateurs", http.StatusInternalServerError)
		return
	}

	jsonResponse2(w, http.StatusOK, users)
	fmt.Println("status: ", http.StatusOK)
	fmt.Println("*** Liste des utilisateurs ENVOYER AVEC SUCCÈS ***")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.Users

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	_, err := database.CreateUser(&newUser)
	if err != nil {
		fmt.Println("Error creating user: ", err)
		http.Error(w, "Erreur lors de la création de l'utilisateur", http.StatusInternalServerError)
		return
	}

	jsonResponse2(w, http.StatusCreated, newUser)
	fmt.Println("status: ", http.StatusCreated)
}

func UpdateUser(w http.ResponseWriter, r *http.Request, id int) {

	user, err := database.GetUserByID(id)
	if err != nil {
		http.Error(w, "Utilisateur non trouvé", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = database.UpdateUser(user)
	if err != nil {
		fmt.Println("Error updating user: ", err)
		http.Error(w, "Erreur lors de la mise à jour de l'utilisateur", http.StatusInternalServerError)
		return
	}

	jsonResponse2(w, http.StatusOK, user)
	fmt.Println("status: ", http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request, id int) {

	err := database.DeleteUser(id)
	if err != nil {
		http.Error(w, "Utilisateur non trouvé", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	fmt.Println("status: ", http.StatusNoContent)
}

// Fonction utilitaire pour extraire l'ID à partir du chemin de l'URL
func extractIDFromPath(path string) int {
	parts := strings.Split(path, "/")
	id, _ := strconv.Atoi(parts[len(parts)-1])
	return id
}
