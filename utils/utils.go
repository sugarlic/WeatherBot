package utils

import "fmt"

func MakeUrlForCity(city string, appID string) string {
	return fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric",
		city, appID)
}

func MakeUrlForCoords(coords string, appID string) string {
	var lat string
	var lon string

	fmt.Sscanf(coords, "%s %s", &lat, &lon)

	return fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s&units=metric",
		lat, lon, appID)
}

func MakeStrFromMap(forecast map[string]interface{}) string {
	res := fmt.Sprintf("Прогноз по городу %s:\nТекущая температура %v °C\nМаксимальная температура: %v °C\nМинимальная температура: %v °C\nПо ощущениям: %v °C\nВидимость: %v метров\nСкорость ветра: %v м/с\nДавление: %v hPa",
		forecast["name"], forecast["temp"], forecast["temp_max"], forecast["temp_min"], forecast["feels_like"], forecast["visibility"], forecast["speed"], forecast["pressure"])

	return res
}
