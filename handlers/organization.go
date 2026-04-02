package handlers

import (
	"e_commerce/database"
	"e_commerce/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetOrganization(c *fiber.Ctx) error {
	orgID := c.Locals("org_uuid").(uuid.UUID)

	var org models.Organization
	if err := database.DB.Where("id = ?", orgID).First(&org).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Organisation introuvable"})
	}

	return c.JSON(org)
}

func UpdateOrganization(c *fiber.Ctx) error {
	orgID := c.Locals("org_uuid").(uuid.UUID)

	var org models.Organization
	if err := database.DB.Where("id = ?", orgID).First(&org).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Organisation introuvable"})
	}

	if err := c.BodyParser(&org); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Format invalide"})
	}

	database.DB.Save(&org)
	return c.JSON(org)
}
