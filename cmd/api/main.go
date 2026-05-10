package main

import (
	"github.com/desirorie/job-tracker-api/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.RegisterRoutes(router)

	router.Run(":8080")
}
