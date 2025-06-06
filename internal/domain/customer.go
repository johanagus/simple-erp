package domain

import "time"

type Customer struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"type:varchar(100)"`
	Address   string `gorm:"type:varchar(255)"`
	Phone     string `gorm:"type:varchar(20)"`
	Email     string `gorm:"type:varchar(100)"`
	IsMember  bool   `gorm:"default:false"`
	MemberNo  string `gorm:"type:varchar(30)"`
	IsActive  bool   `gorm:"default:true"`
	IsDeleted bool   `gorm:"default:false"`
	CreatedBy int
	CreatedAt time.Time
	UpdatedBy int
	UpdatedAt time.Time
	DeletedBy int
	DeletedAt time.Time
}
