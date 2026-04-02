package handlers

import (
	"e_commerce/database"
	"e_commerce/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetProducts(c *fiber.Ctx) error {
	orgID := c.Locals("org_uuid").(uuid.UUID)

	var products []models.Product
	if err := database.DB.Where("organization_id = ?", orgID).Find(&products).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Erreur lors de la récupération des produits"})
	}

	return c.JSON(products)
}

func CreateProduct(c *fiber.Ctx) error {
	orgID := c.Locals("org_uuid").(uuid.UUID)

	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Format de requête invalide"})
	}

	product.OrganizationID = orgID
	if err := database.DB.Create(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Erreur lors de la création du produit"})
	}

	return c.Status(fiber.StatusCreated).JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	orgID := c.Locals("org_uuid").(uuid.UUID)
	productID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "ID produit invalide"})
	}

	var product models.Product
	if err := database.DB.Where("id = ? AND organization_id = ?", productID, orgID).First(&product).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Produit introuvable"})
	}

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Format de requête invalide"})
	}

	database.DB.Save(&product)
	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	orgID := c.Locals("org_uuid").(uuid.UUID)
	productID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "ID produit invalide"})
	}

	if err := database.DB.Where("id = ? AND organization_id = ?", productID, orgID).Delete(&models.Product{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Erreur lors de la suppression"})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}

func GetLowStockProducts(c *fiber.Ctx) error {
	orgID := c.Locals("org_uuid").(uuid.UUID)

	var products []models.Product
	if err := database.DB.Where("organization_id = ? AND stock <= min_stock", orgID).Find(&products).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Erreur lors de la récupération des produits en stock faible"})
	}

	return c.JSON(products)
}
