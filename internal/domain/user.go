package domain

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"type:varchar(50)" json:"email" validate:"required,email"`
	Phone     string    `gorm:"type:varchar(14)" json:"phone"`
	Firstname string    `gorm:"type:varchar(10)" json:"firstname"`
	Lastname  string    `gorm:"type:varchar(10)" json:"lastname"`
	Password  string    `gorm:"type:varchar(255)" json:"password,omitempty"`
	CompanyId int       `json:"company_id"`
	RoleID    int       `json:"role_id"`
	Roles     []string  `gorm:"-" json:"roles"` // tidak di simpan di database
	CreatedAt time.Time `json:"-"`              // tidak tampilkan di respon api
	UpdatedAt time.Time `json:"-"`              // tidak tampilkan di respon api
}
