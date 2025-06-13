package domain

import "time"

type Product struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	SKU         string `gorm:"type:varchar(10)" json:"sku"`
	Barcode     string `gorm:"type:varchar(20)" json:"barcode"`
	Name        string `gorm:"type:varchar(100)" json:"name"`
	Type        string `gorm:"type:varchar(20)" json:"type"`
	Brand       string `gorm:"type:varchar(20)" json:"brand"`
	IsSN        bool
	CategoryID  int
	PriceLevel1 float64 `gorm:"not null" json:"price"`
	PriceLevel2 float64
	PriceLevel3 float64
	Cost        float64
	Stock       int
	Unit        string `gorm:"type:varchar(10)"`
	ImageURL    string `gorm:"type:varchar(255)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
