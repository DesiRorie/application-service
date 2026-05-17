package repository

import (
	"github.com/desirorie/job-tracker-api/internal/database"
	"github.com/desirorie/job-tracker-api/internal/models"
)

func GetApplicationByID(id int) (models.Application, error) {
	var application models.Application

	err := database.DB.First(&application, id).Error
	if err != nil {
		return models.Application{}, err
	}

	return application, nil
}

func SaveApplication(application models.Application) (models.Application, error) {
	err := database.DB.Save(&application).Error
	if err != nil {
		return models.Application{}, err
	}

	return application, nil
}
