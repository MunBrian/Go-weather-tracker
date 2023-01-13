package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type apiConfigData struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp`
	} `json:"main"`
}

func loadApiConfig(filename string) (apiConfigData, error) {
	//read file - apiConfigFile
	bytes, err := ioutil.ReadFile(filename)

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

func main() {
	fmt.Println("The app is up and running")
}
