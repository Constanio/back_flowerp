package routes

import (
	"e_commerce/handlers"
	"e_commerce/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Setup(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173", // URL par défaut de Vite
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	api := app.Group("/api")

	// AUTH
	auth := api.Group("/auth")
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)
	auth.Get("/me", middleware.AuthRequired, handlers.GetMe)

	// PROTECTED ROUTES (WITH TENANT ISOLATION)
	protected := api.Group("/")
	protected.Use(middleware.AuthRequired)
	protected.Use(middleware.TenantScope)

	// ORGANIZATION
	api.Get("/organization", middleware.AuthRequired, middleware.TenantScope, handlers.GetOrganization)
	api.Put("/organization", middleware.AuthRequired, middleware.TenantScope, handlers.UpdateOrganization)

	// CUSTOMERS
	customers := protected.Group("/customers")
	customers.Get("/", handlers.GetCustomers)
	customers.Post("/", handlers.CreateCustomer)
	customers.Put("/:id", handlers.UpdateCustomer)
	customers.Delete("/:id", handlers.DeleteCustomer)

	// SUPPLIERS
	suppliers := protected.Group("/suppliers")
	suppliers.Get("/", handlers.GetSuppliers)
	suppliers.Post("/", handlers.CreateSupplier)
	suppliers.Put("/:id", handlers.UpdateSupplier)
	suppliers.Delete("/:id", handlers.DeleteSupplier)

	// PRODUCTS
	products := protected.Group("/products")
	products.Get("/", handlers.GetProducts)
	products.Post("/", handlers.CreateProduct)
	products.Put("/:id", handlers.UpdateProduct)
	products.Delete("/:id", handlers.DeleteProduct)
	products.Get("/low-stock", handlers.GetLowStockProducts)

	// INVOICES
	invoices := protected.Group("/invoices")
	invoices.Get("/", handlers.GetInvoices)
	invoices.Get("/:id", handlers.GetInvoice)
	invoices.Post("/", handlers.CreateInvoice)
	invoices.Patch("/:id/status", handlers.UpdateInvoiceStatus)

	// QUOTES
	quotes := protected.Group("/quotes")
	quotes.Get("/", handlers.GetQuotes)
	quotes.Get("/:id", handlers.GetQuote)
	quotes.Post("/", handlers.CreateQuote)
	quotes.Post("/:id/convert", handlers.ConvertQuoteToInvoice)

	// DASHBOARD
	dashboard := protected.Group("/dashboard")
	dashboard.Get("/stats", handlers.GetDashboardStats)
	dashboard.Get("/revenue", handlers.GetMonthlyRevenue)
}
