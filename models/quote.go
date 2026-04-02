package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Quote struct {
	ID             uuid.UUID      `gorm:"type:uuid;primary_key;" json:"id"`
	OrganizationID uuid.UUID      `gorm:"type:uuid;not null;index" json:"organization_id"`
	CustomerID     uuid.UUID      `gorm:"type:uuid;not null;index" json:"customer_id"`
	Number         string         `gorm:"unique;not null" json:"number"`
	TotalAmount    float64        `gorm:"type:decimal(10,2);not null" json:"total_amount"`
	Status         string         `gorm:"default:pending" json:"status"` // pending, accepted, rejected, converted
	ValidUntil     time.Time      `json:"valid_until"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`

	Customer Customer    `gorm:"foreignKey:CustomerID" json:"customer"`
	Items    []QuoteItem `gorm:"foreignKey:QuoteID" json:"items"`
}

type QuoteItem struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	QuoteID   uuid.UUID `gorm:"type:uuid;not null;index" json:"quote_id"`
	ProductID uuid.UUID `gorm:"type:uuid;not null" json:"product_id"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	Price     float64   `gorm:"type:decimal(10,2);not null" json:"price"`

	Product Product `gorm:"foreignKey:ProductID" json:"product"`
}

func (q *Quote) BeforeCreate(tx *gorm.DB) (err error) {
	q.ID = uuid.New()
	return
}

func (qi *QuoteItem) BeforeCreate(tx *gorm.DB) (err error) {
	qi.ID = uuid.New()
	return
}
