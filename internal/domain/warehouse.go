package domain

import "time"

type Warehouse struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	CompanyID uint   `gorm:"not null"`                     // Foreign key to Company
	StoreID   uint   `gorm:"not null" validate:"required"` // Foreign key to Store
	Name      string `gorm:"type:varchar(100)" validate:"required"`
	Address   string `gorm:"type:varchar(255)" validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
