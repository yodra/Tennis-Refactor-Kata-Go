package tenniskata

import (
	"strings"
)

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
	if playerName == "player1" {
		game.m_score1 += 1
	} else {
		game.m_score2 += 1
	}
}

func (game *tennisGame1) GetScore() (score string) {

	if game.m_score1 == game.m_score2 {
		return resultEqual(game.m_score1)

	} else if game.m_score1 >= 4 || game.m_score2 >= 4 {
		return checkIfWins(game.m_score1, game.m_score2)

	} else {
		return scoreDuringMatch(game.m_score1, game.m_score2)
	}
}

func resultEqual(input1 int) (score string) {

	switch input1 {
	case 0:
		score = "Love-All"
	case 1:
		score = "Fifteen-All"
	case 2:
		score = "Thirty-All"
	default:
		score = "Deuce"
	}

	return
}

func checkIfWins(input1 int, input2 int) (score string) {

	difference := input1 - input2

	switch difference {
	case -2, -3, -4:
		score = "Win for player2"
	case -1:
		score = "Advantage player2"
	case 1:
		score = "Advantage player1"
	case 2, 3, 4:
		score = "Win for player1"
	}

	return
}

func scoreDuringMatch(input1 int, input2 int) (score string) {

	slicedScore := []string{"", "-", ""}

	switch input1 {
	case 0:
		slicedScore[0] = "Love"
	case 1:
		slicedScore[0] = "Fifteen"
	case 2:
		slicedScore[0] = "Thirty"
	case 3:
		slicedScore[0] = "Forty"
	}
	switch input2 {
	case 0:
		slicedScore[2] = "Love"
	case 1:
		slicedScore[2] = "Fifteen"
	case 2:
		slicedScore[2] = "Thirty"
	case 3:
		slicedScore[2] = "Forty"
	}

	score = strings.Join(slicedScore, "")

	return
}
