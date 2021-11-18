package tenniskata

type tennisGame1 struct {
	m_score1    int
	m_score2    int
	player1Name string
	player2Name string
}

func TennisGame1(player1Name string, player2Name string) TennisGame {
	game := &tennisGame1{
		player1Name: player1Name,
		player2Name: player2Name}

	return game
}

func (game *tennisGame1) WonPoint(playerName string) {
	if playerName == game.player1Name {
		game.m_score1 += 1
	} else {
		game.m_score2 += 1
	}
}

func (game *tennisGame1) GetScore() string {
	if game.m_score1 == game.m_score2 {
		return getDrawScore(game.m_score1)
	} else if game.m_score1 >= 4 || game.m_score2 >= 4 {
		return getWinScore(game.m_score1, game.m_score2)
	} else {
		return getNormalScore(game.m_score1, game.m_score2)
	}
}

func getDrawScore(score int) string {
	switch score {
	case 0:
		return "Love-All"
	case 1:
		return "Fifteen-All"
	case 2:
		return "Thirty-All"
	default:
		return "Deuce"
	}
}

func getWinScore(player1Score int, player2Score int) string {
	minusResult := player1Score - player2Score

	switch {
	case minusResult == -1:
		return "Advantage player2"
	case minusResult == 1:
		return "Advantage player1"
	case minusResult >= 2:
		return "Win for player1"
	default:
		return "Win for player2"
	}
}

func getNormalScore(player1Score int, player2Score int) string {
	pointsValue := map[int]string{
		0: "Love",
		1: "Fifteen",
		2: "Thirty",
		3: "Forty",
	}

	score := pointsValue[player1Score]
	score += "-"
	score += pointsValue[player2Score]

	return score
}
