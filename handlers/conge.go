package handlers

import (
	"sirh/database"
	"sirh/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Liste toutes les demandes (pour Manager/RH)
func GetAllDemandes(c *gin.Context) {
	var demandes []models.DemandeConge
	err := database.DB.Preload("Utilisateur").Preload("TypeConge").Find(&demandes).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, demandes)
}

// TYPES DE CONGÉS
func GetTypesConges(c *gin.Context) {
	var types []models.TypeConge
	database.DB.Find(&types)
	c.JSON(http.StatusOK, types)
}

// DEMANDES DE CONGÉS DE L'UTILISATEUR CONNECTÉ
func GetMesDemandes(c *gin.Context) {
	userID, _ := c.Get("user_id")
	var demandes []models.DemandeConge
	err := database.DB.Preload("TypeConge").Where("utilisateur_id = ?", userID).Find(&demandes).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, demandes)
}

func CreateDemandeConge(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Utilisateur non identifié"})
		return
	}

	var input struct {
		TypeCongeID   uint      `json:"type_conge_id" binding:"required"`
		DateDebut     time.Time `json:"date_debut" binding:"required"`
		DateFin       time.Time `json:"date_fin" binding:"required"`
		JoursDemandes float64   `json:"jours_demandes" binding:"required"`
		Motif         string    `json:"motif"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides : " + err.Error()})
		return
	}

	demande := models.DemandeConge{
		UtilisateurID: userID.(uint),
		TypeCongeID:   input.TypeCongeID,
		DateDebut:     input.DateDebut,
		DateFin:       input.DateFin,
		JoursDemandes: input.JoursDemandes,
		Motif:         input.Motif,
		Statut:        models.StatutEnAttente,
	}

	if err := database.DB.Create(&demande).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la création de la demande : " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, demande)
}

func ApprouverDemande(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	approuvePar, _ := c.Get("user_id")

	var demande models.DemandeConge
	if err := database.DB.First(&demande, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Demande non trouvée"})
		return
	}

	demande.Statut = models.StatutApprouve
	demande.ApprouveParID = new(uint)
	*demande.ApprouveParID = approuvePar.(uint)

	database.DB.Save(&demande)

	// Mettre à jour le solde de l'utilisateur
	database.DB.Model(&models.SoldeConge{}).
		Where("utilisateur_id = ? AND type_conge_id = ?", demande.UtilisateurID, demande.TypeCongeID).
		UpdateColumn("jours_utilises", database.DB.Raw("jours_utilises + ?", demande.JoursDemandes))

	c.JSON(http.StatusOK, gin.H{"message": "Demande approuvée"})
}

func RefuserDemande(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	database.DB.Model(&models.DemandeConge{}).Where("id = ?", id).Update("statut", models.StatutRefuse)
	c.JSON(http.StatusOK, gin.H{"message": "Demande refusée"})
}

// SOLDES
func GetMesSoldes(c *gin.Context) {
	userID, _ := c.Get("user_id")
	var soldes []models.SoldeConge
	err := database.DB.Preload("TypeConge").Where("utilisateur_id = ?", userID).Find(&soldes).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, soldes)
}
