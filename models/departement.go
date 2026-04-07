package models

import (
	"time"
	"gorm.io/gorm"
)

type Departement struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Nom         string         `gorm:"unique;not null" json:"nom"`
	Code        string         `json:"code"`
	ManagerID   *uint          `json:"manager_id"`
	Manager     *Utilisateur   `gorm:"foreignKey:ManagerID" json:"manager"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
