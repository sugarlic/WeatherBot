package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

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

func InitDb() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=1 dbname=postgres sslmode=disable")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InsertIntoDb(db *sql.DB, forecast map[string]interface{}) error {
	_, err := db.Exec("INSERT INTO forecasts VALUES($1, $2, $3, $4, $5, $6, $7, $8)", forecast["name"],
		forecast["temp"], forecast["temp_max"], forecast["temp_min"], forecast["feels_like"], forecast["visibility"], forecast["speed"], forecast["pressure"])

	if err != nil {
		return err
	}

	return nil
}

func DeleteFromDb(db *sql.DB, city string) error {
	_, err := db.Exec("DELETE FROM forecasts WHERE city_name = $1", city)

	if err != nil {
		return err
	}

	return nil
}

func CheckRowExists(db *sql.DB, city string) (bool, error) {
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM forecasts WHERE city_name = '%s' LIMIT 1)",
		city)
	var exists bool
	err := db.QueryRow(query).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func ReadFromDb(db *sql.DB, city string) (map[string]interface{}, error) {
	row, err := db.Query("SELECT * FROM forecasts WHERE city_name = $1", city)

	if err != nil {
		return nil, err
	}
	defer row.Close()

	forecast := make(map[string]interface{})
	for row.Next() {
		var name, visibility, pressure interface{}
		var temp, temp_max, temp_min, feels_like, speed float32
		err = row.Scan(&name, &temp, &temp_max, &temp_min,
			&feels_like, &visibility, &speed, &pressure)
		if err != nil {
			return nil, err
		}

		forecast["name"] = name
		forecast["temp"] = temp
		forecast["temp_max"] = temp_max
		forecast["temp_min"] = temp_min
		forecast["feels_like"] = feels_like
		forecast["visibility"] = visibility
		forecast["speed"] = speed
		forecast["pressure"] = pressure
	}

	return forecast, nil
}
