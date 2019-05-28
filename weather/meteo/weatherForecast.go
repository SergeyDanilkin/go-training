package meteo

import (
	"fmt"
	"strconv"
	"time"
)

type WeatherForecast struct {
}

func (wf *WeatherForecast) Print(w Weather) {
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
		"Влажность воздуха " + h + "%. Восход солнца " + wf.getTime(sunrise) +
		" ,заход солнца " + wf.getTime(sunset) + ".")

}
func (wf *WeatherForecast) getTime(t int64) string {
	return time.Unix(t, 0).Format("15:04")
}
