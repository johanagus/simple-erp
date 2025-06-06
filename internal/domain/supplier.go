package domain

import "time"

type Supplier struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:varchar(100)"`
	Address     string `gorm:"type:varchar(255)"`
	District    string `gorm:"type:varchar(10)"`
	Regency     string `gorm:"type:varchar(10)"`
	Province    string `gorm:"type:varchar(10)"`
	Country     string `gorm:"type:varchar(10)"`
	Phone       string `gorm:"type:varchar(20)"`
	Email       string `gorm:"type:varchar(100)"`
	CompanyName string `gorm:"type:varchar(50)"`
	TaxNumber   string `gorm:"type:varchar(18)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
