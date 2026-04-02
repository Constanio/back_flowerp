package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthRequired(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "En-tête Authorization manquant",
		})
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "votre_secret_tres_securise"
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Token invalide ou expiré",
		})
	}

	claims := token.Claims.(jwt.MapClaims)
	c.Locals("user_id", claims["user_id"])
	c.Locals("organization_id", claims["organization_id"])
	c.Locals("role", claims["role"])

	return c.Next()
}
