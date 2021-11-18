package tenniskata

type tennisGame2 struct {
	P1point int
	P2point int

	P1res       string
	P2res       string
	player1Name string
	player2Name string
}

func TennisGame2(player1Name string, player2Name string) TennisGame {
	game := &tennisGame2{
		player1Name: player1Name,
		player2Name: player2Name,
	}

	return game
}

func (game *tennisGame2) GetScore() string {
	score := ""
	if game.P1point < 4 {
		scoreResults := []string{"Love-All", "Fifteen-All", "Thirty-All", "-All"}
		score = scoreResults[game.P1point]
	}
	if game.P1point >= 3 {
		score = "Deuce"
	}

	if game.P1point != game.P2point {
		if game.P1point < 4 || game.P2point < 4 {
			scoreResults := []string{"Love", "Fifteen", "Thirty", "Forty", ""}
			game.P1res = scoreResults[game.P1point]
			game.P2res = scoreResults[game.P2point]
			score = game.P1res + "-" + game.P2res
		}
	}

	if game.P1point > game.P2point && game.P2point >= 3 {
		score = "Advantage player1"
	}

	if game.P2point > game.P1point && game.P1point >= 3 {
		score = "Advantage player2"
	}

	if game.P1point >= 4 && (game.P1point-game.P2point) >= 2 {
		score = "Win for player1"
	}
	if game.P2point >= 4 && (game.P2point-game.P1point) >= 2 {
		score = "Win for player2"
	}
	return score
}

func (game *tennisGame2) SetP1Score(number int) {
	for i := 0; i < number; i++ {
		game.P1Score()
	}
}

func (game *tennisGame2) SetP2Score(number int) {
	for i := 0; i < number; i++ {
		game.P2Score()
	}
}

func (game *tennisGame2) P1Score() {
	game.P1point++
}

func (game *tennisGame2) P2Score() {
	game.P2point++
}

func (game *tennisGame2) WonPoint(player string) {
	if player == "player1" {
		game.P1Score()
	} else {
		game.P2Score()
	}
}
