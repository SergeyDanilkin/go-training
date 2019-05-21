package meteo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Weather struct {
	City string `json:"name"`
	//General WeatherMain `json:"main"`
	Wind     Wind    `json:"wind"`
	Time     Sun     `json:"sys"`
	Temp     float64 `json:"temp"`
	Pressure float64 `json:"pressure"`
	Humidity int     `json:"humidity"`
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
	Clouds   []struct {
		Descript string `json:"description"`
	} `json:"weather"`
}

func (w *Weather) GetWeather(city string) Weather {
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
	var weath Weather
	err1 := json.Unmarshal(data, &weath)
	if err1 != nil {
		fmt.Println("error:", err)
	}
	return weath
}

func (w *Weather) Print() {

	t, _, _ := w.GetTemperature()
	wSpeed, wGust, dir := w.GetWind()
	temp := strconv.FormatFloat(t, 'f', -1, 32)
	gust := ""
	h := strconv.Itoa(w.GetHumidity())
	sunrise, sunset := w.GetTime()

	if t > 0 {
		temp = "+" + strconv.FormatFloat(t, 'f', -1, 32)
	}

	if wGust != 0 {
		gust = " с порывами до " + strconv.Itoa(wGust) + "м/с"
	}
	fmt.Println("\n" + "Сегодня в городе " + w.GetCity() + " " + w.GetCloudiness() +
		", температура воздуха " + temp + "°С, ветер " + dir + " " + strconv.Itoa(wSpeed) + "м/с" + gust + ".\n" +
		"Влажность воздуха " + h + "%. Восход солнца " + w.getTime(sunrise) +
		" ,заход солнца " + w.getTime(sunset) + ".")
}
func (w *Weather) getTime(t int64) string {
	return time.Unix(t, 0).Format("15:04")
}
func (w *Weather) GetTemperature() (temp float64, tempMin float64, tempMax float64) {
	return w.Temp, w.TempMin, w.TempMax
}

func (w *Weather) GetCloudiness() (description string) {
	return w.Clouds[0].Descript
}
func (w *Weather) GetHumidity() (humidity int) {
	return w.Humidity
}
func (w *Weather) GetWind() (speed int, gust int, direction string) {
	return int(w.Wind.Speed), int(w.Wind.Gust), w.Wind.GetWindDirection()
}
func (w *Weather) GetTime() (sunrise int64, sunset int64) {
	return w.Time.Sunrise, w.Time.Sunset
}
func (w *Weather) GetCity() string {
	return w.City
}
