package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func TenantScope(c *fiber.Ctx) error {
	orgIDRaw := c.Locals("organization_id")
	if orgIDRaw == nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Organisation non identifiée",
		})
	}

	orgIDStr, ok := orgIDRaw.(string)
	if !ok {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Format d'ID organisation invalide",
		})
	}

	orgID, err := uuid.Parse(orgIDStr)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "ID organisation invalide",
		})
	}

	// Injecter l'UUID parsé dans le contexte pour utilisation dans les handlers
	c.Locals("org_uuid", orgID)

	return c.Next()
}
