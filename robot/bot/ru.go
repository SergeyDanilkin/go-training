package bot

type RuBot struct {
	Name string
}

func (rb RuBot) Command(command string) string {
	switch command {
	case "Привет":
		return rb.Phrase1()

	case "Время":
		return rb.Phrase2()

	case "Дата":
		return rb.Phrase3()

	case "День":
		return rb.Phrase4()

	case rb.GetExitCommand():
		return rb.Phrase5()

	default:
		return "Я не понимаю"
	}
}
func (rb RuBot) Phrase1() string {
	return "Привет, я " + rb.Name + "!"
}
func (rb RuBot) Phrase2() string {
	return "Сейчас " + GetTime()
}
func (rb RuBot) Phrase3() string {
	return "Сегодня " + GetDate()
}
func (rb RuBot) Phrase4() string {
	return "Сегодня " + GetWeekday()
}
func (rb RuBot) Phrase5() string {
	return "Пока"
}
func (rb RuBot) GetExitCommand() string {
	return "Пока"
}
