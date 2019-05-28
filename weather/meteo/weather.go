package meteo

type WeatherMain struct {
	Temp     float64 `json:"temp"`
	Pressure float64 `json:"pressure"`
	Humidity int     `json:"humidity"`
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
}

type Weather struct {
	City    string      `json:"name"`
	Wind    Wind        `json:"wind"`
	General WeatherMain `json:"main"`
	Time    Sun         `json:"sys"`
	Clouds  []struct {
		Descript string `json:"description"`
	} `json:"weather"`
}

func (w *Weather) GetTemperature() (temp float64, tempMin float64, tempMax float64) {
	return w.General.Temp, w.General.TempMin, w.General.TempMax
}

func (w *Weather) GetCloudiness() (description string) {
	return w.Clouds[0].Descript
}
func (w *Weather) GetHumidity() (humidity int) {
	return w.General.Humidity
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
