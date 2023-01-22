package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/amanuel15/fiber_server/pkg/configs"
	"github.com/amanuel15/fiber_server/pkg/models"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	DB *gorm.DB
}

var DB DBInstance

func ConnectDB() {
	createDB()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", configs.DB_HOST, configs.DB_USER, configs.DB_PASSWORD, configs.DB_NAME, configs.DB_PORT)
	log.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database! \n", err)
		os.Exit(2)
	}
	log.Println("Connected to database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations...")

	db.AutoMigrate(&models.User{}, &models.Blog{}, &models.Comment{}, &models.Reference{})

	DB = DBInstance{
		DB: db,
	}
}

// Create database if it doesn't exist
func createDB() {
	url := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable",
		configs.DB_HOST, configs.DB_PORT, configs.DB_USER, configs.DB_PASSWORD)
	log.Println(url)
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Failed to connect to postgres to create database \n", err)
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s;", configs.DB_NAME))
	if err != nil && err.Error() == "pq: database \"fiberone\" already exists" {
		log.Println("Database already exists")
	} else if err != nil {
		log.Fatal("Failed to create database\n", err.Error())
	} else {
		log.Println("Successfully created database: ", configs.DB_NAME)
	}
}
