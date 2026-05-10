package services

import (
	"errors"

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
