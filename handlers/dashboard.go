package handlers

import (
	"e_commerce/database"
	"e_commerce/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type DashboardStats struct {
	MonthlyRevenue   float64 `json:"monthly_revenue"`
	CustomerCount    int64   `json:"customer_count"`
	UnpaidInvoices   int64   `json:"unpaid_invoices"`
	LowStockProducts int64   `json:"low_stock_products"`
}

func GetDashboardStats(c *fiber.Ctx) error {
	orgID := c.Locals("org_uuid").(uuid.UUID)
	
	now := time.Now()
	firstDayOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	var stats DashboardStats

	// Chiffre d'affaires du mois
	database.DB.Model(&models.Invoice{}).
		Where("organization_id = ? AND status = ? AND paid_at >= ?", orgID, "paid", firstDayOfMonth).
		Select("COALESCE(SUM(total_amount), 0)").Scan(&stats.MonthlyRevenue)

	// Nombre de clients
	database.DB.Model(&models.Customer{}).
		Where("organization_id = ?", orgID).
		Count(&stats.CustomerCount)

	// Factures impayées
	database.DB.Model(&models.Invoice{}).
		Where("organization_id = ? AND status = ?", orgID, "unpaid").
		Count(&stats.UnpaidInvoices)

	// Stock faible
	database.DB.Model(&models.Product{}).
		Where("organization_id = ? AND stock <= min_stock", orgID).
		Count(&stats.LowStockProducts)

	return c.JSON(stats)
}

func GetMonthlyRevenue(c *fiber.Ctx) error {
	orgID := c.Locals("org_uuid").(uuid.UUID)

	type Result struct {
		Month   string  `json:"month"`
		Revenue float64 `json:"revenue"`
	}

	var results []Result
	// Requête simplifiée pour Chart.js (PostgreSQL)
	database.DB.Model(&models.Invoice{}).
		Where("organization_id = ? AND status = ?", orgID, "paid").
		Select("TO_CHAR(paid_at, 'YYYY-MM') as month, COALESCE(SUM(total_amount), 0) as revenue").
		Group("month").
		Order("month DESC").
		Limit(12).
		Scan(&results)

	return c.JSON(results)
}
