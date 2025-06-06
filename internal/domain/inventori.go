package domain

import "time"

type Inventory struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	WarehouseID int  `gorm:"type:varchar(50)"` // Warehouse or store location
	ProductID   uint `gorm:"not null"`         // Foreign key to Product
	Quantity    int  `gorm:"not null"`         // Current stock quantity
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
