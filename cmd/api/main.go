package main

import (
	"log"

	"github.com/desirorie/job-tracker-api/internal/database"
	"github.com/desirorie/job-tracker-api/internal/models"
	"github.com/desirorie/job-tracker-api/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load .env file")
	}

	database.Connect()
	database.DB.AutoMigrate(&models.Application{})

	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":8080")
}
