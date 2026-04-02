package handlers

import (
	"e_commerce/database"
	"e_commerce/models"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type RegisterRequest struct {
	OrganizationName string `json:"organization_name"`
	AdminName        string `json:"admin_name"`
	AdminEmail       string `json:"admin_email"`
	AdminPassword    string `json:"admin_password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Format de requête invalide"})
	}

	// Création de l'organisation
	org := models.Organization{
		Name:  req.OrganizationName,
		Email: req.AdminEmail,
	}

	if err := database.DB.Create(&org).Error; err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "Organisation déjà existante ou erreur"})
	}

	// Création de l'admin
	user := models.User{
		OrganizationID: org.ID,
		Name:           req.AdminName,
		Email:          req.AdminEmail,
		Password:       req.AdminPassword,
		Role:           "admin",
	}

	if err := database.DB.Create(&user).Error; err != nil {
		database.DB.Delete(&org) // Rollback manuel basique
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "Utilisateur déjà existant ou erreur"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Organisation et admin créés avec succès",
		"organization_id": org.ID,
	})
}

func Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Format de requête invalide"})
	}

	var user models.User
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Identifiants incorrects"})
	}

	if !user.CheckPassword(req.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Identifiants incorrects"})
	}

	// Génération JWT
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "votre_secret_tres_securise"
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":         user.ID.String(),
		"organization_id": user.OrganizationID.String(),
		"role":            user.Role,
		"exp":             time.Now().Add(time.Hour * 72).Unix(),
	})

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"token": t,
		"user":  user,
	})
}

func GetMe(c *fiber.Ctx) error {
	userIDStr := c.Locals("user_id").(string)
	userID, _ := uuid.Parse(userIDStr)

	var user models.User
	if err := database.DB.Preload("Organization").Where("id = ?", userID).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Utilisateur introuvable"})
	}

	return c.JSON(user)
}
