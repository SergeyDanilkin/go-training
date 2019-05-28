package meteo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Meteorologist struct{}

func (m *Meteorologist) GetWeather(city string) Weather {
	url := "http://api.openweathermap.org/data/2.5/weather?q=" + city + "&lang=ru&units=metric&appid=2c19a8c670afc70f2ae7a81f229fce3d"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("error:", err)
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error:", err)
	}
	res.Body.Close()
	var weather Weather
	print(res.Body)
	err1 := json.Unmarshal(data, &weather)
	if err1 != nil {
		fmt.Println("error:", err)
	}
	return weather
}
