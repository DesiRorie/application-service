package services

import (
	"errors"
	"strings"

	"github.com/desirorie/job-tracker-api/internal/dto"
	"github.com/desirorie/job-tracker-api/internal/models"
)

func GetApplicationById(id int) (models.Application, error) {
	for _, application := range models.Applications {
		if application.ID == id {
			return application, nil
		}
	}
	return models.Application{}, errors.New("application not found")
}
func GetApplications() []models.Application {
	return models.Applications
}
func CreateApplication(application dto.CreateApplicationRequest) (dto.CreateApplicationRequest, error) {

	newApplication := models.Application{
		ID:       len(models.Applications) + 1,
		Company:  application.Company,
		Position: application.Position,
		Status:   application.Status,
		Notes:    application.Notes,
	}
	models.Applications = append(models.Applications, newApplication)
	return application, nil
}

func DeleteApplication(id int) error {
	for i, application := range models.Applications {

		if application.ID == id {
			models.Applications = append(
				models.Applications[:i],
				models.Applications[i+1:]...,
			)
			return nil
		}
	}
	return errors.New("Delete unsuccessful, application not found")
}
func UpdateApplication(id int, company string) (models.Application, error) {
	for i, application := range models.Applications {
		if application.ID == id {
			models.Applications[i].Company = company
			return models.Applications[i], nil
		}
	}

	return models.Application{}, errors.New("application not found")
}

func ListApplications(company string) []models.Application {
	if company == "" {
		return models.Applications

	}
	var filtered []models.Application

	for _, application := range models.Applications {
		if strings.EqualFold(application.Company, company) {
			filtered = append(filtered, application)
		}
	}

	return filtered
}
