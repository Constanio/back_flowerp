package database

import (
	"e_commerce/models"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() (*gorm.DB, error) {
	// Utilisation de variables d'environnement ou valeurs par défaut
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "postgres"
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "password"
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "flow_erp"
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", 
		dbHost, dbUser, dbPassword, dbName, dbPort)

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("impossible de se connecter à la base de données: %w", err)
	}

	DB = connection
	
	// Migration des modèles
	err = connection.AutoMigrate(
		&models.Organization{},
		&models.User{},
		&models.Customer{},
		&models.Supplier{},
		&models.Product{},
		&models.Invoice{},
		&models.InvoiceItem{},
		&models.Quote{},
		&models.QuoteItem{},
	)
	
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la migration: %w", err)
	}

	return connection, nil
}
