package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/FranciscoZapa/TestMobileAppTeamcoreBackEnd/src/configs"
	"github.com/FranciscoZapa/TestMobileAppTeamcoreBackEnd/src/handlers"
	"github.com/FranciscoZapa/TestMobileAppTeamcoreBackEnd/src/models"
)

func GetData() (response models.Response) {
	httpClient := &http.Client{}

	req, err := http.NewRequest("GET", configs.BaseUrl+configs.GetDataUrl, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Authorization", configs.SecurityToken)
	res, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	var data models.Data
	json.NewDecoder(res.Body).Decode(&data)

	return handlers.CreateResponse(data)
}
