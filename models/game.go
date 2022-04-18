package models

// Game is the proccess where players solves quiz.
type Game interface {
	// Play the game.
	Play() (FinishedGame, error)
}

// Round is the single round of game.
type Round interface {
	// Play the round.
	Play() (Results, error)
}

// FinishedGame is the summary of ended game.
type FinishedGame interface {
	// Results of the game.
	Results() (Results, error)
}

// Results in one of game moment.
type Results map[PlayerName]Points
