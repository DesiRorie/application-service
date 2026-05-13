package handlers

import (
	"net/http"
	"strconv"

	"github.com/desirorie/job-tracker-api/internal/models"
	"github.com/desirorie/job-tracker-api/internal/services"
	"github.com/gin-gonic/gin"
)

func CreateApplication(c *gin.Context) {
	var req models.Application

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	createdApplication, err := services.CreateApplication(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{
		"message":     "application created",
		"application": createdApplication,
	})

}

func GetApplications(c *gin.Context) {

	applications := services.GetApplications()
	c.JSON(http.StatusOK, applications)

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
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, application)
}

func DeleteApplication(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "bad request",
		})
		return
	}

	err = services.DeleteApplication(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Application has been deleted.",
	})
}
func UpdateApplication(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid application id"})
		return
	}

	var req struct {
		Company string `json:"company"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	updatedApplication, err := services.UpdateApplication(id, req.Company)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "company updated",
		"application": updatedApplication,
	})
}
func ListApplications(c *gin.Context) {

	company := c.Query("company")
	applications := services.ListApplications(company)
	c.JSON(http.StatusOK, applications)

}
