package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type apiConfigData struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

func loadApiConfig(filename string) (apiConfigData, error) {
	//read file - apiConfigFile
	bytes, err := os.ReadFile(filename)

	//handle error
	if err != nil {
		return apiConfigData{}, err
	}

	//define var c as type struct
	var c apiConfigData

	//decode json to struct
	err = json.Unmarshal(bytes, &c)

	if err != nil {
		return apiConfigData{}, err
	}

	return c, nil

}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func getCityWeather(city string) (weatherData, error) {
	apiConfig, err := loadApiConfig(".apiConfig")

	if err != nil {
		return weatherData{}, err
	}

	res, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + apiConfig.OpenWeatherMapApiKey + "&q=" + city)

	if err != nil {
		return weatherData{}, err
	}

	var d weatherData

	//read data and store value in d
	if err := json.NewDecoder(res.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}

	return d, nil
}

func main() {
	http.HandleFunc("/hello", hello)

	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		//split the URL using "/" into 3 parts
		//get the 3 element (city) using index 2 from array created using SplitN
		//assign value to var city
		city := strings.SplitN(r.URL.Path, "/", 3)[2]

		//pass city value to getCityWeather func
		data, err := getCityWeather(city)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf=8")

		//create data and write out to user
		json.NewEncoder(w).Encode(data)

	})

	//start server
	http.ListenAndServe(":8080", nil)
}
