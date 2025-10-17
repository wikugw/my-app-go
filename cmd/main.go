package main

import (
	"log"

	"my-app/database"
	"my-app/models"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	database.Connect()

	err := database.DB.AutoMigrate(&models.Employee{})
	if err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}

	log.Println("🚀 Migration successful!")
}
