package database

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"roadmap/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateProjectRootPath() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "../")
}

func Connect() string {

	err := godotenv.Load(CreateProjectRootPath() + "/.env")
	if err != nil {
		return fmt.Sprintf("Error loading .env file: %s", err)
	}

	dbUser := os.Getenv("USER")
	dbPassword := os.Getenv("PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("HOST")
	dbPort := os.Getenv("PORT")

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", dbUser, dbPassword, dbName, dbHost, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Sprintf("Failed to connect to database: %s", err)
	}

	db.AutoMigrate(&model.User{}, &model.Project{}, &model.Progress{})

	Db, err := db.DB()
	
	if err != nil {
		return fmt.Sprintf("Error loading .env file: %s", err)
	}
	defer Db.Close()
	
	return "OK"
}
