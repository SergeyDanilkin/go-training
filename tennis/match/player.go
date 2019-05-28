package match

type Player struct {
	Name        string
	Skill       int
	PlayerScore int
	PlayerGames int
	PlayerSets  int
	PlayerWins  int
}

func (p *Player) Play(c chan int, w chan *Player) {
	for {
		ball, ok := <-c
		if !ok {
			w <- p
			break
		}
		if ball <= p.Skill {
			nb := Ball.Intn(100)
			c <- nb
		}
		if ball > p.Skill {
			close(c)
			break
		}
	}
}
