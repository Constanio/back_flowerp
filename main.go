package main

import (
	"e_commerce/database"
	"e_commerce/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Connexion à la base de données
	_, err := database.Connect()
	if err != nil {
		log.Fatalf("Échec de la connexion à la base de données: %v", err)
	}

	// Seeder
	database.Seed()

	app := fiber.New(fiber.Config{
		AppName: "FlowERP API",
	})

	// Middlewares globaux
	app.Use(logger.New())
	app.Use(recover.New())

	// Configuration des routes
	routes.Setup(app)

	// Lancement du serveur
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	log.Printf("Serveur démarré sur le port %s", port)
	log.Fatal(app.Listen(":" + port))
}
