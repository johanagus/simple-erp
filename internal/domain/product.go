package domain

import "time"

type Product struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	SKU         string `gorm:"type:varchar(10)"`
	Barcode     string `gorm:"type:varchar(20)"`
	Name        string `gorm:"type:varchar(100)"`
	Type        string `gorm:"type:varchar(20)"`
	Brand       string `gorm:"type:varchar(20)"`
	IsSN        bool
	CategoryID  int
	PriceLevel1 float64
	PriceLevel2 float64
	PriceLevel3 float64
	Cost        float64
	Stock       int
	Unit        string `gorm:"type:varchar(10)"`
	ImageURL    string `gorm:"type:varchar(255)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
