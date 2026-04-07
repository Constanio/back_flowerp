package routes

import (
	"sirh/handlers"
	"sirh/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	api := r.Group("/api")
	{
		// AUTH
		auth := api.Group("/auth")
		{
			auth.POST("/register", handlers.Register)
			auth.POST("/login", handlers.Login)
			auth.GET("/me", middleware.AuthRequired, handlers.GetMe)
		}

		// ROUTES PROTÉGÉES
		protected := api.Group("/")
		protected.Use(middleware.AuthRequired)
		{
			// DASHBOARD
			protected.GET("/dashboard/stats", handlers.GetDashboardStats)
			protected.GET("/dashboard/revenue", handlers.GetMonthlyRevenue)

			// UTILISATEURS
			utilisateurs := protected.Group("/utilisateurs")
			{
				utilisateurs.GET("/", handlers.GetUtilisateurs)
				utilisateurs.GET("/:id", handlers.GetUtilisateur)
				utilisateurs.POST("/", handlers.CreateUtilisateur)
				utilisateurs.PUT("/:id", handlers.UpdateUtilisateur)
				utilisateurs.DELETE("/:id", handlers.DeleteUtilisateur)
			}

			// ORGANISATION
			protected.GET("/departements", handlers.GetDepartements)
			protected.POST("/departements", handlers.CreateDepartement)
			protected.GET("/postes", handlers.GetPostes)
			protected.POST("/postes", handlers.CreatePoste)

			// CONGÉS
			conges := protected.Group("/conges")
			{
				conges.GET("/types", handlers.GetTypesConges)
				conges.GET("/mes-demandes", handlers.GetMesDemandes)
				conges.POST("/demande", handlers.CreateDemandeConge)
				conges.GET("/mes-soldes", handlers.GetMesSoldes)
				
				// Routes RH/Manager
				conges.GET("/toutes-les-demandes", handlers.GetAllDemandes)
				conges.PATCH("/approuver/:id", handlers.ApprouverDemande)
				conges.PATCH("/refuser/:id", handlers.RefuserDemande)
			}
		}

		// PING
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong"})
		})
	}
}
