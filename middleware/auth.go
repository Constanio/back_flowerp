package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("votre_secret_jwt") // Doit correspondre à celui de handlers/auth.go

type Claims struct {
	UserID uint `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func AuthRequired(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token d'autorisation manquant"})
		return
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("méthode de signature inattendue: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token invalide ou expiré"})
		return
	}

	c.Set("user_id", claims.UserID)
	c.Set("user_role", claims.Role)
	c.Next()
}
