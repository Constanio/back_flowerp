package handlers

import (
	"e_commerce/database"
	"e_commerce/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetSuppliers(c *fiber.Ctx) error {
	orgID := c.Locals("org_uuid").(uuid.UUID)

	var suppliers []models.Supplier
	if err := database.DB.Where("organization_id = ?", orgID).Find(&suppliers).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Erreur récupération fournisseurs"})
	}

	return c.JSON(suppliers)
}

func CreateSupplier(c *fiber.Ctx) error {
	orgID := c.Locals("org_uuid").(uuid.UUID)

	var supplier models.Supplier
	if err := c.BodyParser(&supplier); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Format invalide"})
	}

	supplier.OrganizationID = orgID
	if err := database.DB.Create(&supplier).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Erreur création fournisseur"})
	}

	return c.Status(fiber.StatusCreated).JSON(supplier)
}

func UpdateSupplier(c *fiber.Ctx) error {
	orgID := c.Locals("org_uuid").(uuid.UUID)
	id, _ := uuid.Parse(c.Params("id"))

	var supplier models.Supplier
	if err := database.DB.Where("id = ? AND organization_id = ?", id, orgID).First(&supplier).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Fournisseur introuvable"})
	}

	if err := c.BodyParser(&supplier); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Format invalide"})
	}

	database.DB.Save(&supplier)
	return c.JSON(supplier)
}

func DeleteSupplier(c *fiber.Ctx) error {
	orgID := c.Locals("org_uuid").(uuid.UUID)
	id, _ := uuid.Parse(c.Params("id"))

	if err := database.DB.Where("id = ? AND organization_id = ?", id, orgID).Delete(&models.Supplier{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Erreur suppression"})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}
