package bot

type EnBot struct {
	Name string
}

func (eb EnBot) Command(command string) string {
	switch command {
	case "Hello":
		return eb.Phrase1()

	case "Time":
		return eb.Phrase2()

	case "Date":
		return eb.Phrase3()

	case "Day":
		return eb.Phrase4()

	case eb.GetExitCommand():
		return eb.Phrase5()

	default:
		return "I don`t understand"
	}
}
func (eb EnBot) Phrase1() string {
	return "Hello, I'm " + eb.Name + "!"
}
func (eb EnBot) Phrase2() string {
	return "Now is " + GetTime()
}
func (eb EnBot) Phrase3() string {
	return "Today is " + GetDate()
}
func (eb EnBot) Phrase4() string {
	return "Today is " + GetWeekday()
}
func (e EnBot) Phrase5() string {
	return "Bye"
}
func (e EnBot) GetExitCommand() string {
	return "Bye"
}
