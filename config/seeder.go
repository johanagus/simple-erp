package config

import (
	"log"

	"github.com/johanagus/simple-erp/internal/domain"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func SeedData(db *gorm.DB) {
	// Check if the database is empty
	var count int64

	db.Model(&domain.Module{}).Count(&count)
	if count > 0 {
		log.Println("User already seeded")
		return
	}

	// Seed modules
	modules := []domain.Module{
		{ID: 1, Name: "create_user", Description: "Create User"},
		{ID: 2, Name: "get_user_by_id", Description: "Find user By ID User"},
		{ID: 3, Name: "get_users", Description: "Find All Users"},
		{ID: 4, Name: "update_user", Description: "Update User"},
	}

	for _, module := range modules {
		if err := db.Create(&module).Error; err != nil {
			log.Fatalf("Failed to seed module %s: %v", module.Name, err)
		}
	}

	role := domain.Role{
		ID:   1,
		Name: "Admin",
	}

	db.Model(&domain.Role{}).Count(&count)
	if count > 0 {
		log.Println("Role already seeded")
		return
	}

	if err := db.Create(&role).Error; err != nil {
		log.Fatalf("Failed to seed role %s: %v", role.Name, err)
	}

	moduleRoles := []domain.ModuleRole{
		{RoleID: role.ID, ModuleID: 1, ModuleName: "create_user"},
		{RoleID: role.ID, ModuleID: 2, ModuleName: "get_user_by_id"},
		{RoleID: role.ID, ModuleID: 3, ModuleName: "get_users"},
		{RoleID: role.ID, ModuleID: 4, ModuleName: "update_user"},
	}

	db.Model(&domain.ModuleRole{}).Count(&count)
	if count > 0 {
		log.Println("ModuleRole already seeded")
		return
	}

	for _, moduleRole := range moduleRoles {
		if err := db.Create(&moduleRole).Error; err != nil {
			log.Fatalf("Failed to seed module role %s: %v", moduleRole.ModuleName, err)
		}
	}

	db.Model(&domain.User{}).Count(&count)
	if count > 0 {
		log.Println("User already seeded")
		return
	}

	password, err := HashPassword("admin123") // Use a secure password
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	// Seed users
	adminUser := domain.User{
		Firstname: "Admin User",
		Email:     "admin@simple.erp",
		Password:  password, // In a real application, use a hashed password
		RoleID:    int(role.ID),
	}

	db.Create(&adminUser)
	log.Println("Database seeded successfully with admin user and role")

}
