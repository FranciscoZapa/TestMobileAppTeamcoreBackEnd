package handlers

import (
	"github.com/FranciscoZapa/TestMobileAppTeamcoreBackEnd/src/models"
)

func CreateResponse(data models.Data) (response models.Response) {
	return models.Response{
		Title:      "Encuesta",
		Day:        data.Date,
		Info:       data.Data,
		ApiVersion: 1,
	}
}
