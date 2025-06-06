package domain

import "time"

type Warehouse struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	CompanyID uint   `gorm:"not null"` // Foreign key to Company
	StoreID   uint   `gorm:"not null"` // Foreign key to Store
	Name      string `gorm:"type:varchar(100)"`
	Address   string `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
