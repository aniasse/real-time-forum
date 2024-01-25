// package api

// import (
// 	"net/http"

// //	"gorm.io/driver/sqlite"
// 	"github.com/gin-gonic/gin"

// 	"forum/database"
// 	"forum/models"
// )

// func HandleLogin(c *gin.Context) {
// 	var user models.Users

// 	// Lire les données JSON de la requête
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Les données d'identification sont invalides"})
// 		return
// 	}

// 	// Recherche de l'utilisateur dans la base de données
// 	result := database.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user)
// 	if result.Error != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Identifiants incorrects"})
// 		return
// 	}

// 	// Connexion réussie
// 	c.JSON(http.StatusOK, gin.H{"message": "Connexion réussie"})
// }

// func HandleRegister(c *gin.Context) {
// 	var newUser models.Users

// 	// Lire les données JSON de la requête
// 	if err := c.ShouldBindJSON(&newUser); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Les données d'inscription sont invalides"})
// 		return
// 	}

// 	// Vérifier si l'utilisateur existe déjà dans la base de données
// 	existingUser := models.Users{}
// 	result := database.DB.Where("email = ?", newUser.Email).First(&existingUser)
// 	if result.RowsAffected > 0 {
// 		c.JSON(http.StatusConflict, gin.H{"error": "Cet e-mail est déjà enregistré"})
// 		return
// 	}

// 	// Ajouter le nouvel utilisateur à la base de données
// 	if err := database.DB.Create(&newUser).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de l'enregistrement de l'utilisateur"})
// 		return
// 	}

// 	// Enregistrement réussi
// 	c.JSON(http.StatusCreated, gin.H{"message": "Enregistrement réussi"})
// }
