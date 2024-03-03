package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/felipezschornack/golang-api-temperature/configs"
	"github.com/felipezschornack/golang-api-temperature/internal/backend/remote/viacep"
	"github.com/felipezschornack/golang-api-temperature/internal/backend/remote/weatherapi"
)

var apiKey string

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		log.Panicln(err)
	}
	apiKey = configs.WeatherAPIKey

	http.HandleFunc("/weather", getWeatherHandler)
	http.ListenAndServe(":8080", nil)
}

func getWeatherHandler(w http.ResponseWriter, r *http.Request) {
	zipcode := r.URL.Query().Get("zipcode") // queryparam

	data, err := viacep.BuscaCep(zipcode)
	if err != nil {
		w.WriteHeader(err.Status)
		json.NewEncoder(w).Encode(err.Message)
		return
	}

	// buscar dados do weatherAPI
	weatherData, err := weatherapi.GetWeather(data.Localidade, apiKey)
	if err != nil {
		w.WriteHeader(err.Status)
		json.NewEncoder(w).Encode(err.Message)
		return
	} else {
		w.Header().Set(http.CanonicalHeaderKey("content-type"), "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(weatherData)
	}

}
