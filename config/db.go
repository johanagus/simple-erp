package config

import (
	"fmt"
	"log"
	"time"

	"github.com/johanagus/simple-erp/internal/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() *gorm.DB {

	LoadEnv() // Load ENV File

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		GetEnv("DB_USER", "root"),
		GetEnv("DB_PASSWORD", ""),
		GetEnv("DB_HOST", "localhost"),
		GetEnv("DB_PORT", "3306"),
		GetEnv("DB_NAME", "simple-erp"),
	)

	var err error
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Gagal koneksi ke database: %v", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Gagal mendapatkan instance SQL DB: %v", err)
	}

	DB.AutoMigrate(
		&domain.Category{},
		&domain.User{},
		&domain.Customer{},
		&domain.Sales{},
		&domain.Product{},
		&domain.Inventory{},
		&domain.SalesItem{},
		&domain.SalesOrder{},
		&domain.SalesOrderItem{},
		&domain.SalesPayment{},
		&domain.Supplier{},
		&domain.Warehouse{},
		&domain.Module{},
		&domain.Role{},
		&domain.ModuleRole{},
		&domain.Store{},
	)

	// konfigurasi koneksi pool
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Berhasil koneksi ke database")
	return DB
}
