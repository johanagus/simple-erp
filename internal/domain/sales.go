package domain

import "time"

type SalesOrder struct {
	ID          uint    `gorm:"primaryKey;autoIncrement"`
	CustomerID  uint    `gorm:"index"` // Foreign key to Customer
	TotalAmount float64 `gorm:"not null"`
	Notes       string  `gorm:"type:varchar(255)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type SalesOrderItem struct {
	ID           uint    `gorm:"primaryKey;autoIncrement"`
	SalesOrderId uint    `gorm:"index"` // Foreign key to Sales Order
	ProductID    uint    `gorm:"index"` // Foreign key to Product
	Quantity     int     `gorm:"not null"`
	Price        float64 `gorm:"not null"`
	Subtotal     float64 `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Sales struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	InvoiceNo    string    `gorm:"type:varchar(30);uniqueIndex"`
	Date         time.Time `gorm:"type:date" json:"date"`
	CustomerID   uint      `gorm:"index"` // Foreign key to Customer
	TotalAmount  float64   `gorm:"not null"`
	PaidAmount   float64   `gorm:"not null"`
	ChangeAmount float64   `gorm:"not null"`
	Notes        string    `gorm:"type:varchar(255)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type SalesItem struct {
	ID        uint    `gorm:"primaryKey;autoIncrement"`
	SalesID   uint    `gorm:"index"` // Foreign key to SalesTransaction
	ProductID uint    `gorm:"index"` // Foreign key to Product
	Quantity  int     `gorm:"not null"`
	Price     float64 `gorm:"not null"`
	Subtotal  float64 `gorm:"not null"`
}
