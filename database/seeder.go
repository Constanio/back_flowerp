package database

import (
	"e_commerce/models"
	"fmt"
	"time"
	"github.com/google/uuid"
)

func Seed() {
	var count int64
	database := DB

	// 1. SEED ORGANIZATION
	var org models.Organization
	database.Model(&models.Organization{}).First(&org)
	if org.ID == uuid.Nil {
		org = models.Organization{
			Name:    "FlowERP Corp",
			Email:   "contact@flowerp.com",
			Address: "123 Rue de l'Innovation, Paris",
		}
		database.Create(&org)
		fmt.Println("✅ Organisation FlowERP créée")
	}

	// 2. SEED USERS
	var userCount int64
	database.Model(&models.User{}).Count(&userCount)
	if userCount == 0 {
		admin := models.User{
			OrganizationID: org.ID,
			Name:           "Admin Global",
			Email:          "admin@flowerp.com",
			Password:       "admin123",
			Role:           "admin",
		}
		database.Create(&admin)
		fmt.Println("✅ Admin créé (admin@flowerp.com / admin123)")
	}

	// 3. SEED CUSTOMERS
	database.Model(&models.Customer{}).Count(&count)
	if count == 0 {
		c1 := models.Customer{
			OrganizationID: org.ID,
			Name:           "Client Alpha SARL",
			Email:          "achat@alpha.fr",
			Phone:          "0145678910",
			Address:        "Paris, France",
			Status:         "active",
		}
		c2 := models.Customer{
			OrganizationID: org.ID,
			Name:           "Boutique Beta",
			Email:          "contact@beta.com",
			Phone:          "0612345678",
			Address:        "Lyon, France",
			Status:         "active",
		}
		database.Create(&c1)
		database.Create(&c2)
		fmt.Println("✅ Clients de test créés")
	}

	// 4. SEED PRODUCTS
	database.Model(&models.Product{}).Count(&count)
	if count == 0 {
		p1 := models.Product{
			OrganizationID: org.ID,
			Name:           "Ordinateur Portable Pro",
			Price:          1200.00,
			Stock:          15,
			MinStock:       5,
			Category:       "Informatique",
		}
		p2 := models.Product{
			OrganizationID: org.ID,
			Name:           "Écran 27 Pouces 4K",
			Price:          350.00,
			Stock:          3,
			MinStock:       5,
			Category:       "Périphériques",
		}
		database.Create(&p1)
		database.Create(&p2)
		fmt.Println("✅ Produits de test créés")
	}

	// Fetch references for quotes/invoices
	var customer models.Customer
	database.Where("organization_id = ?", org.ID).First(&customer)
	var product models.Product
	database.Where("organization_id = ?", org.ID).First(&product)

	// 5. SEED QUOTES
	database.Model(&models.Quote{}).Count(&count)
	if count == 0 && customer.ID != uuid.Nil && product.ID != uuid.Nil {
		quote := models.Quote{
			OrganizationID: org.ID,
			CustomerID:     customer.ID,
			Number:         "DEV-2026-001",
			TotalAmount:    product.Price,
			Status:         "pending",
			ValidUntil:     time.Now().AddDate(0, 1, 0),
		}
		database.Create(&quote)
		database.Create(&models.QuoteItem{QuoteID: quote.ID, ProductID: product.ID, Quantity: 1, Price: product.Price})
		fmt.Println("✅ Devis de test créé")
	}

	// 6. SEED INVOICES
	database.Model(&models.Invoice{}).Count(&count)
	if count == 0 && customer.ID != uuid.Nil && product.ID != uuid.Nil {
		now := time.Now()
		invoice := models.Invoice{
			OrganizationID: org.ID,
			CustomerID:     customer.ID,
			Number:         "FAC-2026-001",
			TotalAmount:    product.Price,
			Status:         "paid",
			PaidAt:         &now,
			DueDate:        now,
		}
		database.Create(&invoice)
		database.Create(&models.InvoiceItem{InvoiceID: invoice.ID, ProductID: product.ID, Quantity: 1, Price: product.Price})
		fmt.Println("✅ Factures de test créées")
	}
}
