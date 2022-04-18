package models

// Player is the user who participates in the room.
type Player interface {
	// Id of player for exact room.
	Id() (PlayerID, error)
	// Player nickname for exact room.
	Name() (PlayerName, error)
	// ReceiveQuestion is way to communicate between player and game.
	ReceiveQuestion(Question) (<-chan PlayerAnswer, error)
	// ReceiveResults is used to send player latest results.
	ReceiveResults(Results) error
	// ReceiveFinishResults is used to send player a finished game.
	ReceiveFinishResults(FinishedGame) error
}

// Id of player.
type PlayerID int

// Name of player.
type PlayerName string
