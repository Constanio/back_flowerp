package database

import (
	"sirh/models"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() (*gorm.DB, error) {
	var dsn string
	
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL != "" {
		dsn = databaseURL
	} else {
		dbHost := os.Getenv("DB_HOST")
		if dbHost == "" { dbHost = "localhost" }
		dbUser := os.Getenv("DB_USER")
		if dbUser == "" { dbUser = "postgres" }
		dbPassword := os.Getenv("DB_PASSWORD")
		if dbPassword == "" { dbPassword = "password" }
		dbName := os.Getenv("DB_NAME")
		if dbName == "" { dbName = "flow_erp" }
		dbPort := os.Getenv("DB_PORT")
		if dbPort == "" { dbPort = "5432" }

		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", 
			dbHost, dbUser, dbPassword, dbName, dbPort)
	}

	// Configuration pour désactiver les FK lors de la migration initiale si nécessaire
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, fmt.Errorf("impossible de se connecter à la base de données: %w", err)
	}

	DB = connection
	
	// Migration des modèles dans un ordre qui réduit les conflits
	err = connection.AutoMigrate(
		&models.TypeConge{},
		&models.Utilisateur{},
		&models.Departement{},
		&models.Poste{},
		&models.SoldeConge{},
		&models.DemandeConge{},
		&models.SalaireEmploye{},
		&models.FichePaie{},
		&models.EvaluationPerformance{},
	)
	
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la migration: %w", err)
	}

	return connection, nil
}
