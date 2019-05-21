package main

import (
	"weather/meteo"
)

func main() {
	var Weather meteo.Weather

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
		Weather := Weather.GetWeather(city)
		Weather.Print()
	}
}
