package match

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Match struct {
	Player1, Player2 Player
	MatchScore       []Score
}

var Ball *rand.Rand

func init() {
	r := rand.NewSource(time.Now().UnixNano())
	Ball = rand.New(r)
}

func (tm *Match) start() {
	var wg sync.WaitGroup
	wg.Add(2)

	c := make(chan int)
	winner := make(chan *Player)

	go func() {
		tm.Player1.Play(c, winner)
		wg.Done()
	}()
	go func() {
		tm.Player2.Play(c, winner)
		wg.Done()
	}()

	c <- Ball.Intn(100)

	for {
		w, ok := <-winner
		if w != nil {
			w.PlayerScore += 15
			close(winner)
		}
		if !ok {
			break
		}
	}
	wg.Wait()
}

func (tm *Match) playGame() {

	for (tm.Player1.PlayerScore < 40) && (tm.Player2.PlayerScore < 40) {
		tm.start()
	}
	winner := tm.getWinner(tm.Player1.PlayerScore, tm.Player2.PlayerScore)
	winner.PlayerGames++
}

func (tm *Match) playSet() {
	for (tm.Player1.PlayerGames < 6) && (tm.Player2.PlayerGames < 6) {
		tm.playGame()
		tm.clearPlayerScore()
	}
	winner := tm.getWinner(tm.Player1.PlayerGames, tm.Player2.PlayerGames)
	winner.PlayerSets++
	winner.PlayerWins++
	setScore := Score{tm.Player1.PlayerGames, tm.Player2.PlayerGames}
	tm.MatchScore = append(tm.MatchScore, setScore)
}

func (tm *Match) PlayMatch() {
	for (tm.Player1.PlayerSets < 2) && (tm.Player2.PlayerSets < 2) {
		tm.playSet()
		tm.clearPlayerScore()
		tm.Player1.PlayerGames = 0
		tm.Player2.PlayerGames = 0
	}
	tm.printWinner()
}

func (tm *Match) clearPlayerScore() {
	tm.Player1.PlayerScore = 0
	tm.Player2.PlayerScore = 0
}

func (tm *Match) getWinner(score1, score2 int) *Player {
	if (score1) > (score2) {
		return &tm.Player1
	}
	return &tm.Player2
}

func (tm *Match) printWinner() {
	fmt.Printf("\n%v: ", tm.Player1.Name)
	for _, set := range tm.MatchScore {
		fmt.Printf("%v ", set.PlayerScore1)
	}

	fmt.Printf("\n%v: ", tm.Player2.Name)
	for _, set := range tm.MatchScore {
		fmt.Printf("%v ", set.PlayerScore2)
	}
	fmt.Println("")
	if tm.Player1.PlayerWins > tm.Player2.PlayerWins {
		fmt.Printf("Win %v", tm.Player1.Name)
	} else {
		fmt.Printf("Win %v", tm.Player2.Name)
	}
	fmt.Println("")
}
