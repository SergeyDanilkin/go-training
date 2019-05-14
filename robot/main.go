package main

import (
	"errors"
	"fmt"
	"os"
	"robot/bot"
)

func main() {
	var language string
	fmt.Scanln(&language)
	bot, err := createBot(language)
	if err != nil {
		fmt.Println("Error has happend: ", err)
		os.Exit(1)
	}
	listenUser(bot)
}

func createBot(language string) (bot.Bot, error) {
	switch language {
	case "English":
		return bot.EnBot{Name: "T-800"}, nil
	case "Russian":
		return bot.RuBot{Name: "Громозека"}, nil
	default:
		return nil, errors.New("Unknown laguage")
	}
}
func listenUser(b bot.Bot) {
	var mess string
	for {
		fmt.Scanln(&mess)
		fmt.Println(b.Command(mess))

		if mess == b.GetExitCommand() {
			break
		}
	}
}
