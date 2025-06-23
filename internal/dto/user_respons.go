package dto

import "github.com/johanagus/simple-erp/internal/domain"

type UserRespons struct {
	ID        uint     `json:"id"`
	Email     string   `json:"email"`
	Phone     string   `json:"phone"`
	Firstname string   `json:"firstname"`
	Lastname  string   `json:"lastname"`
	CompanyId int      `json:"company_id"`
	RoleID    int      `json:"role_id"`
	Roles     []string `json:"roles"`
}

func ToUserRespons(u *domain.User) UserRespons {
	return UserRespons{
		ID:        u.ID,
		Email:     u.Email,
		Phone:     u.Phone,
		Firstname: u.Firstname,
		Lastname:  u.Lastname,
		CompanyId: u.CompanyId,
		RoleID:    u.RoleID,
		Roles:     u.Roles,
	}
}
