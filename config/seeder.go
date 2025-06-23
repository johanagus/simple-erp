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

	// Seed modules
	modules := []domain.Module{
		{ID: 1, Name: "get_user", Description: "Get User"},
		{ID: 2, Name: "create_user", Description: "Create User"},
		{ID: 3, Name: "update_user", Description: "Update User"},
		{ID: 4, Name: "delete_user", Description: "Delete User"},
		{ID: 5, Name: "get_product", Description: "Get Product"},
		{ID: 6, Name: "create_product", Description: "Create Product"},
		{ID: 7, Name: "update_procuct", Description: "Update Product"},
		{ID: 8, Name: "delete_product", Description: "Delete Product"},
		{ID: 9, Name: "get_customer", Description: "Get Customer"},
		{ID: 10, Name: "create_customer", Description: "Create Customer"},
		{ID: 11, Name: "update_customer", Description: "Update Customer"},
		{ID: 12, Name: "delete_customer", Description: "Delete Customer"},
		{ID: 13, Name: "get_sales", Description: "Get Sales"},
		{ID: 14, Name: "get_sales_all", Description: "Get Sales All Outlet"},
		{ID: 15, Name: "create_sales", Description: "Create Sales"},
		{ID: 16, Name: "refund_sales", Description: "Refund Sales"},
		{ID: 17, Name: "get_inventori", Description: "Get Inventori"},
		{ID: 18, Name: "create_inventori", Description: "Create Inventori"},
		{ID: 19, Name: "update_inventori", Description: "Update Inventori"},
		{ID: 20, Name: "delete_inventori", Description: "Delete Inventori"},
		{ID: 21, Name: "get_store", Description: "Get Store"},
		{ID: 22, Name: "create_store", Description: "Create Store"},
		{ID: 23, Name: "update_store", Description: "Update Store"},
		{ID: 24, Name: "delete_store", Description: "Delete Store"},
		{ID: 25, Name: "get_company", Description: "Get Company"},
		{ID: 26, Name: "create_company", Description: "Create Company"},
		{ID: 27, Name: "update_company", Description: "Update Company"},
		{ID: 28, Name: "delete_company", Description: "Delete Company"},
		{ID: 29, Name: "get_warehouse", Description: "Get Warehouse"},
		{ID: 30, Name: "create_warehouse", Description: "Create Warehouse"},
		{ID: 31, Name: "update_warehouse", Description: "Update Warehouse"},
		{ID: 32, Name: "delete_warehouse", Description: "Delete Warehouse"},
		{ID: 33, Name: "get_supplier", Description: "Get Supplier"},
		{ID: 34, Name: "create_supplier", Description: "Create Supplier"},
		{ID: 35, Name: "update_supplier", Description: "Update Supplier"},
		{ID: 36, Name: "delete_supplier", Description: "Delete Supplier"},
	}

	for _, module := range modules {
		if err := db.Save(&module).Error; err != nil {
			log.Fatalf("Failed to seed module %s: %v", module.Name, err)
		}
	}

	role := domain.Role{
		ID:   1,
		Name: "Admin",
	}

	db.Model(&domain.Role{}).Count(&count)
	if count > 0 {
		log.Println("Skip Role seeding, already exists")
		return
	}

	if err := db.Create(&role).Error; err != nil {
		log.Fatalf("Failed to seed role %s: %v", role.Name, err)
	}

	moduleRoles := []domain.ModuleRole{
		{RoleID: role.ID, ModuleID: 1, ModuleName: "get_user"},
		{RoleID: role.ID, ModuleID: 2, ModuleName: "create_user"},
		{RoleID: role.ID, ModuleID: 3, ModuleName: "update_user"},
		{RoleID: role.ID, ModuleID: 4, ModuleName: "delete_user"},
		{RoleID: role.ID, ModuleID: 5, ModuleName: "get_product"},
		{RoleID: role.ID, ModuleID: 6, ModuleName: "create_product"},
		{RoleID: role.ID, ModuleID: 7, ModuleName: "update_procuct"},
		{RoleID: role.ID, ModuleID: 8, ModuleName: "delete_user"},
		{RoleID: role.ID, ModuleID: 9, ModuleName: "get_customer"},
		{RoleID: role.ID, ModuleID: 10, ModuleName: "create_customer"},
		{RoleID: role.ID, ModuleID: 11, ModuleName: "update_customer"},
		{RoleID: role.ID, ModuleID: 12, ModuleName: "delete_customer"},
		{RoleID: role.ID, ModuleID: 13, ModuleName: "get_sales"},
		{RoleID: role.ID, ModuleID: 14, ModuleName: "get_sales_all"},
		{RoleID: role.ID, ModuleID: 15, ModuleName: "create_sales"},
		{RoleID: role.ID, ModuleID: 16, ModuleName: "refund_sales"},
		{RoleID: role.ID, ModuleID: 17, ModuleName: "get_inventori"},
		{RoleID: role.ID, ModuleID: 18, ModuleName: "create_inventori"},
		{RoleID: role.ID, ModuleID: 19, ModuleName: "update_inventori"},
		{RoleID: role.ID, ModuleID: 20, ModuleName: "delete_inventori"},
		{RoleID: role.ID, ModuleID: 21, ModuleName: "get_store"},
		{RoleID: role.ID, ModuleID: 22, ModuleName: "create_store"},
		{RoleID: role.ID, ModuleID: 23, ModuleName: "update_store"},
		{RoleID: role.ID, ModuleID: 24, ModuleName: "delete_store"},
		{RoleID: role.ID, ModuleID: 25, ModuleName: "get_company"},
		{RoleID: role.ID, ModuleID: 26, ModuleName: "create_company"},
		{RoleID: role.ID, ModuleID: 27, ModuleName: "update_company"},
		{RoleID: role.ID, ModuleID: 28, ModuleName: "delete_company"},
		{RoleID: role.ID, ModuleID: 29, ModuleName: "get_warehouse"},
		{RoleID: role.ID, ModuleID: 30, ModuleName: "create_warehouse"},
		{RoleID: role.ID, ModuleID: 31, ModuleName: "update_warehouse"},
		{RoleID: role.ID, ModuleID: 32, ModuleName: "delete_warehouse"},
		{RoleID: role.ID, ModuleID: 33, ModuleName: "get_supplier"},
		{RoleID: role.ID, ModuleID: 34, ModuleName: "create_supplier"},
		{RoleID: role.ID, ModuleID: 35, ModuleName: "update_supplier"},
		{RoleID: role.ID, ModuleID: 36, ModuleName: "delete_supplier"},
	}

	for _, moduleRole := range moduleRoles {
		if err := db.Save(&moduleRole).Error; err != nil {
			log.Fatalf("Failed to seed module role %s: %v", moduleRole.ModuleName, err)
		}
	}

	db.Model(&domain.User{}).Count(&count)
	if count > 0 {
		log.Println("Skip User seeding, already exists")
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
