package database

import (
	"sirh/models"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Utilitaire pour transformer un uint en *uint
func toPtr(i uint) *uint {
	return &i
}

func Seed() {
	// 1. TYPES DE CONGÉS
	typesConges := []models.TypeConge{
		{Nom: "Congés Payés", Description: "2.08 jours par mois", JoursParAn: 25, Couleur: "#4CAF50"},
		{Nom: "RTT", Description: "Réduction du temps de travail", JoursParAn: 10, Couleur: "#2196F3"},
		{Nom: "Maladie", Description: "Arrêt maladie justifié", JoursParAn: 0, Couleur: "#F44336"},
	}
	for i := range typesConges {
		DB.FirstOrCreate(&typesConges[i], models.TypeConge{Nom: typesConges[i].Nom})
	}

	// 2. DÉPARTEMENTS
	deps := []models.Departement{
		{Nom: "Direction Générale", Code: "DIR"},
		{Nom: "Ressources Humaines", Code: "RH"},
		{Nom: "Informatique", Code: "IT"},
		{Nom: "Commercial", Code: "COM"},
	}
	for i := range deps {
		DB.FirstOrCreate(&deps[i], models.Departement{Nom: deps[i].Nom})
	}

	// 3. POSTES
	postes := []models.Poste{
		{Titre: "Directeur Général", DepartementID: toPtr(deps[0].ID), SalaireMin: 6000},
		{Titre: "Responsable RH", DepartementID: toPtr(deps[1].ID), SalaireMin: 4000},
		{Titre: "Lead Developer", DepartementID: toPtr(deps[2].ID), SalaireMin: 5000},
		{Titre: "Développeur Junior", DepartementID: toPtr(deps[2].ID), SalaireMin: 3000},
	}
	for i := range postes {
		DB.FirstOrCreate(&postes[i], models.Poste{Titre: postes[i].Titre})
	}

	// 4. UTILISATEURS (AUTHENTIFICATION)
	password, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	passStr := string(password)

	utilisateurs := []models.Utilisateur{
		{
			Matricule:      "ADM001",
			Email:          "admin@sirh.com",
			MotDePasseHash: passStr,
			Prenom:         "Alice",
			Nom:            "Admin",
			Role:           models.RoleAdmin,
			Statut:         models.StatutActif,
			DateEmbauche:   time.Now(),
			DepartementID:  toPtr(deps[0].ID),
			PosteID:        toPtr(postes[0].ID),
		},
		{
			Matricule:      "RH001",
			Email:          "rh@sirh.com",
			MotDePasseHash: passStr,
			Prenom:         "Robert",
			Nom:            "Hess",
			Role:           models.RoleRH,
			Statut:         models.StatutActif,
			DateEmbauche:   time.Now(),
			DepartementID:  toPtr(deps[1].ID),
			PosteID:        toPtr(postes[1].ID),
		},
		{
			Matricule:      "MGR001",
			Email:          "manager@sirh.com",
			MotDePasseHash: passStr,
			Prenom:         "Marc",
			Nom:            "Anger",
			Role:           models.RoleManager,
			Statut:         models.StatutActif,
			DateEmbauche:   time.Now(),
			DepartementID:  toPtr(deps[2].ID),
			PosteID:        toPtr(postes[2].ID),
		},
		{
			Matricule:      "EMP001",
			Email:          "employe@sirh.com",
			MotDePasseHash: passStr,
			Prenom:         "Émilie",
			Nom:            "Ploye",
			Role:           models.RoleEmploye,
			Statut:         models.StatutActif,
			DateEmbauche:   time.Now(),
			DepartementID:  toPtr(deps[2].ID),
			PosteID:        toPtr(postes[3].ID),
		},
	}

	for i := range utilisateurs {
		var u models.Utilisateur
		err := DB.Where("email = ?", utilisateurs[i].Email).First(&u).Error
		if err != nil {
			DB.Create(&utilisateurs[i])
			
			for _, tc := range typesConges {
				DB.Create(&models.SoldeConge{
					UtilisateurID: utilisateurs[i].ID,
					TypeCongeID:   tc.ID,
					Annee:         time.Now().Year(),
					TotalJours:    float64(tc.JoursParAn),
					JoursUtilises: 0,
				})
			}
		}
	}

	log.Println("Seeder SIRH : Données d'authentification créées avec succès.")
}
