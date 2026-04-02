package handlers

import (
	"e_commerce/database"
	"e_commerce/models"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetQuotes(c *fiber.Ctx) error {
	orgID := c.Locals("org_uuid").(uuid.UUID)

	var quotes []models.Quote
	if err := database.DB.Preload("Customer").Where("organization_id = ?", orgID).Find(&quotes).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Erreur lors de la récupération des devis"})
	}

	return c.JSON(quotes)
}

func GetQuote(c *fiber.Ctx) error {
	orgID := c.Locals("org_uuid").(uuid.UUID)
	id, _ := uuid.Parse(c.Params("id"))

	var quote models.Quote
	if err := database.DB.Preload("Customer").Preload("Items.Product").Where("id = ? AND organization_id = ?", id, orgID).First(&quote).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Devis introuvable"})
	}

	return c.JSON(quote)
}

func CreateQuote(c *fiber.Ctx) error {
	orgID := c.Locals("org_uuid").(uuid.UUID)

	var quote models.Quote
	if err := c.BodyParser(&quote); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Format de requête invalide"})
	}

	quote.OrganizationID = orgID
	// Génération basique d'un numéro de devis (DEV-YYYY-MM-DD-RAND)
	quote.Number = fmt.Sprintf("DEV-%s-%s", time.Now().Format("20060102"), uuid.New().String()[:8])

	if err := database.DB.Create(&quote).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Erreur lors de la création du devis"})
	}

	return c.Status(fiber.StatusCreated).JSON(quote)
}

func ConvertQuoteToInvoice(c *fiber.Ctx) error {
	orgID := c.Locals("org_uuid").(uuid.UUID)
	quoteID, _ := uuid.Parse(c.Params("id"))

	var quote models.Quote
	if err := database.DB.Preload("Items").Where("id = ? AND organization_id = ?", quoteID, orgID).First(&quote).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Devis introuvable"})
	}

	// Création de la facture basée sur le devis
	invoice := models.Invoice{
		OrganizationID: orgID,
		CustomerID:     quote.CustomerID,
		Number:         fmt.Sprintf("FAC-%s-%s", time.Now().Format("20060102"), uuid.New().String()[:8]),
		TotalAmount:    quote.TotalAmount,
		Status:         "unpaid",
		DueDate:        time.Now().AddDate(0, 0, 30), // 30 jours pour payer
	}

	if err := database.DB.Create(&invoice).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Erreur lors de la conversion"})
	}

	// Copie des items
	for _, qi := range quote.Items {
		item := models.InvoiceItem{
			InvoiceID: invoice.ID,
			ProductID: qi.ProductID,
			Quantity:  qi.Quantity,
			Price:     qi.Price,
		}
		database.DB.Create(&item)
	}

	// Mise à jour du devis
	quote.Status = "converted"
	database.DB.Save(&quote)

	return c.JSON(invoice)
}
