package main

import (
	"weather/meteo"
)

func main() {
	var Meteorologist meteo.Meteorologist
	var WeatherForecast meteo.WeatherForecast

	var cities = []string{
		"Mahilyow",
		"Rome",
		"Madrid",
		"Paris",
		"London",
		"Brasilia",
		"Calgary",
		"Sydney",
		"Tokio",
		"Moscow",
		"Oslo",
		"Murmansk",
	}

	for _, city := range cities {
		Weather := Meteorologist.GetWeather(city)
		WeatherForecast.Print(Weather)
	}
}
