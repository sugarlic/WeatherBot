package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"example.com/m/configs"
	"example.com/m/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var appID string = "db7172c6a976b06502762e915d239656"

func ReplyToCityForecast(db *sql.DB, update tgbotapi.Update, msg *tgbotapi.MessageConfig) {
	reply, err := MakeRequestByCity(db, update.Message.Text)
	if err == nil {
		msg.Text = reply
	} else {
		msg.Text = err.Error()
	}
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
}

func ReplyToCoordsForecast(update tgbotapi.Update, msg *tgbotapi.MessageConfig) {
	reply, err := MakeRequestByCoords(update.Message.Text)
	if err == nil {
		msg.Text = reply
	} else {
		msg.Text = err.Error()
	}
}

func MakeRequestByCity(db *sql.DB, city string) (string, error) {
	forecast := make(map[string]interface{})
	city_exist, err := utils.CheckRowExists(db, city)
	if err != nil {
		return "", err
	}
	if city_exist {
		log.Println("reading from db")
		forecast, err = utils.ReadFromDb(db, city)
		if err != nil {
			return "", err
		}
	} else {
		configs.Start = time.Now()
		url := utils.MakeUrlForCity(city, appID)
		forecast, err = MakeRequestToOpWether(url)
		if err != nil {
			return "", err
		}

		err = utils.InsertIntoDb(db, forecast)
		if err != nil {
			return "", err
		}
	}

	if time.Since(configs.Start).Seconds() > 60 {
		configs.Start = time.Now()
		utils.DeleteFromDb(db, city)
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
