package handlers

import (
	"e_commerce/database"
	"e_commerce/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetCustomers(c *fiber.Ctx) error {
	orgID := c.Locals("org_uuid").(uuid.UUID)

	var customers []models.Customer
	if err := database.DB.Where("organization_id = ?", orgID).Find(&customers).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Erreur lors de la récupération des clients"})
	}

	return c.JSON(customers)
}

func CreateCustomer(c *fiber.Ctx) error {
	orgID := c.Locals("org_uuid").(uuid.UUID)

	var customer models.Customer
	if err := c.BodyParser(&customer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Format de requête invalide"})
	}

	customer.OrganizationID = orgID
	if err := database.DB.Create(&customer).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Erreur lors de la création du client"})
	}

	return c.Status(fiber.StatusCreated).JSON(customer)
}

func UpdateCustomer(c *fiber.Ctx) error {
	orgID := c.Locals("org_uuid").(uuid.UUID)
	customerID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "ID client invalide"})
	}

	var customer models.Customer
	if err := database.DB.Where("id = ? AND organization_id = ?", customerID, orgID).First(&customer).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Client introuvable"})
	}

	if err := c.BodyParser(&customer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Format de requête invalide"})
	}

	database.DB.Save(&customer)
	return c.JSON(customer)
}

func DeleteCustomer(c *fiber.Ctx) error {
	orgID := c.Locals("org_uuid").(uuid.UUID)
	customerID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "ID client invalide"})
	}

	if err := database.DB.Where("id = ? AND organization_id = ?", customerID, orgID).Delete(&models.Customer{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Erreur lors de la suppression"})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}
