package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func MakeRequestToOpWether(city string) {
	appID := "db7172c6a976b06502762e915d239656"

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric",
		city, appID)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return
	}
	defer response.Body.Close()

	// body, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// log.Println(string(body))

	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	for key, elem := range data {
		log.Println(key, elem)
	}

	if main, ok := data["main"].([]interface{}); ok {
		for _, elem := range main {
			log.Println(elem)
		}
	}

	// cities := make([]string, 0)
	// if list, ok := data["list"].([]interface{}); ok {
	//     for _, item := range list {
	//         if city, ok := item.(map[string]interface{})["name"].(string); ok {
	//             if sys, ok := item.(map[string]interface{})["sys"].(map[string]interface{}); ok {
	//                 country := sys["country"].(string)
	//                 cities = append(cities, fmt.Sprintf("%s (%s)", city, country))
	//             }
	//         }
	//     }
	// }

	// fmt.Println("city:", cities)
	// if list, ok := data["list"].([]interface{}); ok && len(list) > 0 {
	//     if item, ok := list[0].(map[string]interface{}); ok {
	//         if cityID, ok := item["id"].(float64); ok {
	//             fmt.Println("city_id=", int(cityID))
	//         }
	//     }
	// }

}

func main() {
	MakeRequestToOpWether("Petersburg")
}
