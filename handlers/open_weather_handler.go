package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func MakeStrFronMap(forecast map[string]interface{}, city string) string {
	res := fmt.Sprintf("Прогноз по городу %s:\nТекущая температура %v °C\nМаксимальная температура: %v °C\nМинимальная температура: %v °C\nПо ощущениям: %v °C\nВидимость: %v метров\nСкорость ветра: %v м/с\nДавление: %v hPa",
		city, forecast["temp"], forecast["temp_max"], forecast["temp_min"], forecast["feels_like"], forecast["visibility"], forecast["speed"], forecast["pressure"])

	return res
}

func MakeRequestToOpWether(city string) (string, error) {
	forecast := make(map[string]interface{})

	appID := "db7172c6a976b06502762e915d239656"

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric",
		city, appID)

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return "", err
	}

	if main, ok := data["main"].(map[string]interface{}); ok {
		for key, elem := range main {
			forecast[key] = elem
		}
	} else {
		return "", errors.New("no current city")
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

	return MakeStrFronMap(forecast, city), nil
}

// func main() {
// 	forecast, _ := MakeRequestToOpWether("Kazan")

// 	fmt.Println(forecast)
// }
