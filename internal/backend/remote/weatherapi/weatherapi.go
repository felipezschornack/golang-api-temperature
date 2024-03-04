package weatherapi

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/felipezschornack/golang-api-temperature/internal/erro"
)

type WeatherAPI struct {
	Location struct {
		Name           string  `json:"name,omitempty"`
		Region         string  `json:"region,omitempty"`
		Country        string  `json:"country,omitempty"`
		Lat            float64 `json:"lat,omitempty"`
		Lon            float64 `json:"lon,omitempty"`
		TzID           string  `json:"tz_id,omitempty"`
		LocaltimeEpoch int     `json:"localtime_epoch,omitempty"`
		Localtime      string  `json:"localtime,omitempty"`
	} `json:"location,omitempty"`
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch,omitempty"`
		LastUpdated      string  `json:"last_updated,omitempty"`
		TempC            float64 `json:"temp_c,omitempty"`
		TempF            float64 `json:"temp_f,omitempty"`
		IsDay            int     `json:"is_day,omitempty"`
		Condition        struct {
			Text string `json:"text,omitempty"`
			Icon string `json:"icon,omitempty"`
			Code int    `json:"code,omitempty"`
		} `json:"condition,omitempty"`
		WindMph    float64 `json:"wind_mph,omitempty"`
		WindKph    float64 `json:"wind_kph,omitempty"`
		WindDegree int     `json:"wind_degree,omitempty"`
		WindDir    string  `json:"wind_dir,omitempty"`
		PressureMb float64 `json:"pressure_mb,omitempty"`
		PressureIn float64 `json:"pressure_in,omitempty"`
		PrecipMm   float64 `json:"precip_mm,omitempty"`
		PrecipIn   float64 `json:"precip_in,omitempty"`
		Humidity   int     `json:"humidity,omitempty"`
		Cloud      int     `json:"cloud,omitempty"`
		FeelslikeC float64 `json:"feelslike_c,omitempty"`
		FeelslikeF float64 `json:"feelslike_f,omitempty"`
		VisKm      float64 `json:"vis_km,omitempty"`
		VisMiles   float64 `json:"vis_miles,omitempty"`
		Uv         float64 `json:"uv,omitempty"`
		GustMph    float64 `json:"gust_mph,omitempty"`
		GustKph    float64 `json:"gust_kph,omitempty"`
	} `json:"current,omitempty"`
}

type WeatherAPIResult struct {
	Celsius    float32 `json:"temp_C"`
	Fahrenheit float32 `json:"temp_F"`
	Kelvin     float32 `json:"temp_K"`
}

func GetWeather(name string, apiKey string) (*WeatherAPIResult, *erro.Erro) {
	url := "https://api.weatherapi.com/v1/current.json?key=" + apiKey + "&q=" + url.QueryEscape(name) + "&aqi=no"

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, erro.New(http.StatusBadRequest, err.Error())
	}

	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, erro.New(resp.StatusCode, err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, erro.New(http.StatusInternalServerError, err.Error())
	}

	var data WeatherAPI
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, erro.New(http.StatusInternalServerError, err.Error())
	}

	//util.PrintDataAsJson("Weather API", data)
	return data.convert(), nil
}

func (a *WeatherAPI) convert() *WeatherAPIResult {
	return &WeatherAPIResult{Celsius: float32(a.Current.TempC), Fahrenheit: float32(a.Current.TempF), Kelvin: float32(a.Current.TempC) + 273}
}
