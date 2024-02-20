package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"example.com/m/utils"
)

var appID string = "db7172c6a976b06502762e915d239656"

func MakeRequestByCity(city string) (string, error) {
	url := utils.MakeUrlForCity(city, appID)

	forecast, err := MakeRequestToOpWether(url)
	if err != nil {
		return "", err
	}

	return utils.MakeStrFromMap(forecast), nil
}

func MakeRequestByCoords(coords string) (string, error) {
	url := utils.MakeUrlForCoords(coords, appID)

	forecast, err := MakeRequestToOpWether(url)
	if err != nil {
		return "", err
	}

	return utils.MakeStrFromMap(forecast), nil
}

func MakeRequestToOpWether(url string) (map[string]interface{}, error) {
	forecast := make(map[string]interface{})

	response, err := http.Get(url)
	if err != nil {
		return forecast, err
	}
	defer response.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return forecast, err
	}

	if city, ok := data["name"]; ok {
		forecast["name"] = city
	} else {
		return forecast, errors.New("no current place")
	}

	if main, ok := data["main"].(map[string]interface{}); ok {
		for key, elem := range main {
			forecast[key] = elem
		}
	}

	if wind, ok := data["wind"].(map[string]interface{}); ok {
		for key, elem := range wind {
			forecast[key] = elem
		}
	}

	if visibility, ok := data["visibility"]; ok {
		forecast["visibility"] = visibility
	}

	if clouds, ok := data["clouds"].(map[string]interface{}); ok {
		if all, ok := clouds["all"]; ok {
			forecast["clouds"] = all
		}
	}

	return forecast, nil
}
