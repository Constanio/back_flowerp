package handlers

import (
	"e_commerce/database"
	"e_commerce/models"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetInvoices(c *fiber.Ctx) error {
	orgID := c.Locals("org_uuid").(uuid.UUID)

	var invoices []models.Invoice
	if err := database.DB.Preload("Customer").Where("organization_id = ?", orgID).Find(&invoices).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Erreur lors de la récupération des factures"})
	}

	return c.JSON(invoices)
}

func GetInvoice(c *fiber.Ctx) error {
	orgID := c.Locals("org_uuid").(uuid.UUID)
	id, _ := uuid.Parse(c.Params("id"))

	var invoice models.Invoice
	if err := database.DB.Preload("Customer").Preload("Items.Product").Where("id = ? AND organization_id = ?", id, orgID).First(&invoice).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Facture introuvable"})
	}

	return c.JSON(invoice)
}

func CreateInvoice(c *fiber.Ctx) error {
	orgID := c.Locals("org_uuid").(uuid.UUID)

	var invoice models.Invoice
	if err := c.BodyParser(&invoice); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Format de requête invalide"})
	}

	invoice.OrganizationID = orgID
	// Génération basique d'un numéro de facture (FAC-YYYY-MM-DD-RAND)
	invoice.Number = fmt.Sprintf("FAC-%s-%s", time.Now().Format("20060102"), uuid.New().String()[:8])

	if err := database.DB.Create(&invoice).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Erreur lors de la création de la facture"})
	}

	return c.Status(fiber.StatusCreated).JSON(invoice)
}

func UpdateInvoiceStatus(c *fiber.Ctx) error {
	orgID := c.Locals("org_uuid").(uuid.UUID)
	id, _ := uuid.Parse(c.Params("id"))

	type StatusUpdate struct {
		Status string `json:"status"`
	}
	var req StatusUpdate
	c.BodyParser(&req)

	var invoice models.Invoice
	if err := database.DB.Where("id = ? AND organization_id = ?", id, orgID).First(&invoice).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Facture introuvable"})
	}

	invoice.Status = req.Status
	if req.Status == "paid" {
		now := time.Now()
		invoice.PaidAt = &now
	}

	database.DB.Save(&invoice)
	return c.JSON(invoice)
}
