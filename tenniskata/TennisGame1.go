package tenniskata

type tennisGame1 struct {
	pointPlayer1 int
	pointPlayer2 int
	player1Name  string
	player2Name  string
}

func TennisGame1(player1Name string, player2Name string) TennisGame {
	game := &tennisGame1{
		player1Name: player1Name,
		player2Name: player2Name}
	return game
}

func (game *tennisGame1) WonPoint(playerName string) {
	if playerName == "player1" {
		game.pointPlayer1 += 1
	} else {
		game.pointPlayer2 += 1
	}
}

func (game *tennisGame1) GetScore() string {
	score := ""
	if game.pointPlayer1 == game.pointPlayer2 {
		switch game.pointPlayer1 {
		case 0:
			score = "Love-All"
		case 1:
			score = "Fifteen-All"
		case 2:
			score = "Thirty-All"
		default:
			score = "Deuce"
		}
	} else if game.pointPlayer1 >= 4 || game.pointPlayer2 >= 4 {

		minusResult := game.pointPlayer1 - game.pointPlayer2

		if minusResult == 1 {
			score = "Advantage player1"
		} else if minusResult == -1 {
			score = "Advantage player2"
		} else if minusResult >= 2 {
			score = "Win for player1"
		} else {
			score = "Win for player2"
		}
	} else {
		_ = getPoint(game.pointPlayer1, score)
		// setMatch = game.pointPlayer1
		score += "-"
		// setMatch = game.pointPlayer2
		_ = getPoint(game.pointPlayer2, score)
	}
	return score
}

func getPoint(matchPoint int, score string) string {

	switch matchPoint {

	case '0':

		score += "Love"

	case 1:

		score += "Fifteen"

	case 2:

		score += "Thirty"

	case 3:

		score += "Forty"

	}

	return score

}