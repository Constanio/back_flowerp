package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Invoice struct {
	ID             uuid.UUID      `gorm:"type:uuid;primary_key;" json:"id"`
	OrganizationID uuid.UUID      `gorm:"type:uuid;not null;index" json:"organization_id"`
	CustomerID     uuid.UUID      `gorm:"type:uuid;not null;index" json:"customer_id"`
	Number         string         `gorm:"unique;not null" json:"number"`
	TotalAmount    float64        `gorm:"type:decimal(10,2);not null" json:"total_amount"`
	Status         string         `gorm:"default:unpaid" json:"status"` // unpaid, paid, cancelled
	PaidAt         *time.Time     `json:"paid_at"`
	DueDate        time.Time      `json:"due_date"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`

	Customer Customer      `gorm:"foreignKey:CustomerID" json:"customer"`
	Items    []InvoiceItem `gorm:"foreignKey:InvoiceID" json:"items"`
}

type InvoiceItem struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	InvoiceID uuid.UUID `gorm:"type:uuid;not null;index" json:"invoice_id"`
	ProductID uuid.UUID `gorm:"type:uuid;not null" json:"product_id"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	Price     float64   `gorm:"type:decimal(10,2);not null" json:"price"`

	Product Product `gorm:"foreignKey:ProductID" json:"product"`
}

func (i *Invoice) BeforeCreate(tx *gorm.DB) (err error) {
	i.ID = uuid.New()
	return
}

func (ii *InvoiceItem) BeforeCreate(tx *gorm.DB) (err error) {
	ii.ID = uuid.New()
	return
}
