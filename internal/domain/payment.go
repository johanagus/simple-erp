package domain

import "time"

type SalesPayment struct {
	ID        uint    `gorm:"primaryKey;autoIncrement"`
	SalesID   uint    `gorm:"index"`             // Foreign key to Sales
	Method    string  `gorm:"type:varchar(30)"`  // Payment method (e.g., cash, debit, credit)
	Amount    float64 `gorm:"not null"`          // Amount paid with this method
	Notes     string  `gorm:"type:varchar(100)"` // Optional notes
	CreatedAt time.Time
	UpdatedAt time.Time
}
