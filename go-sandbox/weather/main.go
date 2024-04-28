package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type WeatherResponse struct {
	Name string `json:"name"`
	Main Main   `json:"main"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"fells_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
	SeaLevel  int     `json:"sea_level"`
	GrndLevel int     `json:"grnd_level"`
}

func main() {

	response, err := http.Get(`https://api.openweathermap.org/data/2.5/weather?lat=44.34&lon=10.99&appid=699b6bdedd2baf679a53adece3b0ab1d&units=metric`)
	if err != nil {
		log.Fatal(err)
	}

	// log.Print(response.StatusCode)
	// log.Print(response.Body)
	bytes, errRead := ioutil.ReadAll(response.Body)
	defer func() {
		e := response.Body.Close()
		if e != nil {
			log.Fatal(e)
		}
	}()

	if errRead != nil {
		log.Fatal(errRead)
	}
	// log.Print(string(bytes))

	var weatherResponse WeatherResponse
	errUnmarshal := json.Unmarshal(bytes, &weatherResponse)
	if errUnmarshal != nil {
		log.Fatal(errUnmarshal)
	}

	log.Printf("%+v", weatherResponse)
}
