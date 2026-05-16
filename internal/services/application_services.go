package services

import (
	"errors"
	"strings"

	"github.com/desirorie/job-tracker-api/internal/database"
	"github.com/desirorie/job-tracker-api/internal/dto"
	"github.com/desirorie/job-tracker-api/internal/models"
	"gorm.io/gorm"
)

func GetApplicationById(id int) (models.Application, error) {
	var application models.Application

	err := database.DB.First(&application, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Application{}, errors.New("application not found")
		}
		return models.Application{}, err
	}

	return application, nil
}

func GetApplications() ([]models.Application, error) {
	var applications []models.Application

	err := database.DB.Find(&applications).Error
	if err != nil {
		return nil, err
	}

	return applications, nil
}

func CreateApplication(req dto.CreateApplicationRequest) (models.Application, error) {
	newApplication := models.Application{
		Company:  req.Company,
		Position: req.Position,
		Status:   strings.ToLower(req.Status),
		Notes:    req.Notes,
	}

	err := database.DB.Create(&newApplication).Error
	if err != nil {
		return models.Application{}, err
	}

	return newApplication, nil
}

func DeleteApplication(id int) error {
	result := database.DB.Delete(&models.Application{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("application not found")
	}

	return nil
}

func UpdateApplication(id int, req dto.UpdateApplication) (models.Application, error) {
	var application models.Application

	err := database.DB.First(&application, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Application{}, errors.New("application not found")
		}
		return models.Application{}, err
	}

	if err != nil {
		return models.Application{}, err
	}
	if req.Company != "" {

		application.Company = req.Company
	}
	if req.Position != "" {
		application.Position = req.Position
	}
	if req.Status != "" {
		application.Status = strings.ToLower(req.Status)
	}
	if req.Notes != "" {
		application.Notes = req.Notes
	}
	err = database.DB.Save(&application).Error
	return application, nil
}

func ListApplications(company string) ([]models.Application, error) {
	var applications []models.Application

	query := database.DB

	if company != "" {
		query = query.Where("LOWER(company) = LOWER(?)", company)
	}

	err := query.Find(&applications).Error
	if err != nil {
		return nil, err
	}

	return applications, nil
}
