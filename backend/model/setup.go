package model

import (
	"fmt"
	"log"
	"os"

	"github.com/ahmadnafi30/monetra/backend/entity"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSLMODE") // tambahkan ini untuk fleksibilitas

	// Format DSN PostgreSQL
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s TimeZone=Asia/Jakarta",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
		sslMode,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Gagal konek ke database: %v", err))
	}

	log.Println(" Berhasil konek ke PostgreSQL")

	DB = database

	err = DB.AutoMigrate(
		&entity.User{},
		&entity.PasswordReset{},
		&entity.Category{},
		// &entity.Transaction{},
		// &entity.Budget{},
	)
	if err != nil {
		log.Fatalf("Gagal migrasi database: %v", err)
	}

	fmt.Println("🚀 Berhasil melakukan migrasi database")
}
