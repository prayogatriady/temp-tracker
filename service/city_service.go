package service

import (
	"encoding/json"
	"net/http"

	"github.com/prayogatriady/temp-tracker/domain"
)

var OpenWeatherMapApiKey string = "97d50198dfe66bd023e9b1bc9d90e61c"

func FindTemp(city string) (domain.WeatherData, error) {

	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + OpenWeatherMapApiKey + "&q=" + city)
	if err != nil {
		return domain.WeatherData{}, err
	}

	defer resp.Body.Close()

	var data domain.WeatherData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return domain.WeatherData{}, err
	}
	return data, nil
}
