package domain

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"type:varchar(50)"`
	Phone     string `gorm:"type:varchar(14)"`
	Firstname string `gorm:"type:varchar(10)"`
	Lastname  string `gorm:"type:varchar(10)"`
	Password  string `gorm:"type:varchar(255)"`
	CompanyId int
	CreatedAt time.Time
	UpdatedAt time.Time
}
