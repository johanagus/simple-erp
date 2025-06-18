package domain

import "time"

type Role struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(100)" json:"name" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ModuleRole struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	RoleID     uint      `json:"role_id"`
	ModuleID   uint      `json:"module_id"`
	ModuleName string    `gorm:"type:varchar(100)" json:"module_name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
