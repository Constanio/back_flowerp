package handlers

import (
	"sirh/database"
	"sirh/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("votre_secret_jwt") // À mettre dans un .env normalement

type Claims struct {
	UserID uint `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func Register(c *gin.Context) {
	var input struct {
		Matricule     string `json:"matricule"`
		Email         string `json:"email" binding:"required,email"`
		Password      string `json:"password" binding:"required,min=6"`
		Prenom        string `json:"prenom" binding:"required"`
		Nom           string `json:"nom" binding:"required"`
		Role          string `json:"role"`
		DepartementID *uint  `json:"departement_id"`
		PosteID       *uint  `json:"poste_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	// Note: Les soldes de congés sont créés automatiquement via le hook AfterCreate dans models/utilisateur.go
	user := models.Utilisateur{
		Matricule:      input.Matricule,
		Email:          input.Email,
		MotDePasseHash: string(hashedPassword),
		Prenom:         input.Prenom,
		Nom:            input.Nom,
		Role:           models.Role(input.Role),
		Statut:         models.StatutActif,
		DateEmbauche:   time.Now(),
		DepartementID:  input.DepartementID,
		PosteID:        input.PosteID,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la création de l'utilisateur"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Utilisateur créé avec succès", "user": user})
}

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.Utilisateur
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Identifiants invalides"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.MotDePasseHash), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Identifiants invalides"})
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: user.ID,
		Role:   string(user.Role),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la génération du token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"user": gin.H{
			"id":     user.ID,
			"email":  user.Email,
			"prenom": user.Prenom,
			"nom":    user.Nom,
			"role":   user.Role,
		},
	})
}

func GetMe(c *gin.Context) {
	// Cette fonction sera utilisée avec un middleware d'authentification
	userID, _ := c.Get("user_id")
	
	var user models.Utilisateur
	if err := database.DB.Preload("Departement").Preload("Poste").First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Utilisateur non trouvé"})
		return
	}

	c.JSON(http.StatusOK, user)
}
