package bot

type Bot interface {
	Phrase1() string
	Phrase2() string
	Phrase3() string
	Phrase4() string
	Phrase5() string
	Command(command string) string
	GetExitCommand() string
}
