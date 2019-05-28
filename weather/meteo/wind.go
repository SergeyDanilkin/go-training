package meteo

type Wind struct {
	Speed float64 `json:"speed"`
	Gust  float64 `json:"gust"`
	Deg   float64 `json:"deg"`
}

func (w *Wind) GetWindDirection() string {
	switch {
	case ((w.Deg < 30) && (w.Deg > 0)) || ((w.Deg <= 360) && (w.Deg > 330)):
		{
			return "северный"
		}
	case (w.Deg < 60) && (w.Deg > 30):
		{
			return "северо-восточный"
		}
	case (w.Deg < 120) && (w.Deg > 60):
		{
			return "восточный"
		}
	case (w.Deg < 150) && (w.Deg > 120):
		{
			return "юго-восточный"
		}
	case (w.Deg < 210) && (w.Deg > 150):
		{
			return "южный"
		}
	case (w.Deg < 240) && (w.Deg > 210):
		{
			return "юго-западный"
		}
	case (w.Deg < 300) && (w.Deg > 240):
		{
			return "западный"
		}
	case (w.Deg < 330) && (w.Deg > 300):
		{
			return "северо-западный"
		}
	default:
		{
			return "изменчивый"
		}
	}
}
