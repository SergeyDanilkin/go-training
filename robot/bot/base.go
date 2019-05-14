package bot

import (
	"time"
)

const TimeZone = "Europe/Minsk"

func GetTime() string {
	t := time.Now()
	loc, _ := time.LoadLocation(TimeZone)
	t = t.In(loc)
	return t.Format("15:04")
}

func GetDate() string {
	t := time.Now()
	return t.Format("January 2, 2006")
}
func GetWeekday() string {
	t := time.Now()
	return t.Format("Monday")
}
