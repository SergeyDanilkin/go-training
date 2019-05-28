package main

import (
	"tennis/match"
)

func main() {
	pl1 := match.Player{Name: "P1", Skill: 90}
	pl2 := match.Player{Name: "P2", Skill: 90}

	m := match.Match{Player1: pl1, Player2: pl2}
	m.PlayMatch()
}
