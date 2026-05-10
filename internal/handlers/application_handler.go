package handlers

import (
	"net/http"
	"strconv"

	"github.com/desirorie/job-tracker-api/internal/models"
	"github.com/desirorie/job-tracker-api/internal/services"
	"github.com/gin-gonic/gin"
)

func CreateApplication(c *gin.Context) {
	var application models.Application

	if err := c.BindJSON(&application); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message":     "application created",
		"application": application,
	})

}

func GetApplications(c *gin.Context) {
	c.JSON(http.StatusOK, models.Applications)

}

func GetApplicationById(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid id",
		})
		return
	}
	application, err := services.GetApplicationById(id)

	c.JSON(200, application)
}
