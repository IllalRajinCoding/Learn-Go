package database

import (
	"IllalRajinCoding/backend-api/config"
	"IllalRajinCoding/backend-api/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB is a global variable that holds the database connection atau sngkatnya koneksi database
var DB *gorm.DB

func InitDB() {
	// konfigurasi database
	dbUser := config.GetEnv("DB_USER", "root")
	dbPassword := config.GetEnv("DB_PASSWORD", "")
	dbHost := config.GetEnv("DB_HOST", "localhost")
	dbPort := config.GetEnv("DB_PORT", "3306")
	dbName := config.GetEnv("DB_NAME", "")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Connected to database successfully")

	// Migrate the User model
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate User model:", err)
		}
}