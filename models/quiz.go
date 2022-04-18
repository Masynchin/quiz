package models

import "time"

// Quiz is the combination of questions related to same theme.
type Quiz interface {
	// Description of the quiz.
	Description() (string, error)
	// Question of the quiz.
	Questions() ([]QuizQuestion, error)
}

// QuizQuestion is the question with options to answer it.
type QuizQuestion interface {
	// Question for viewing.
	Question() (Question, error)
	// Check for correctness of the player answer.
	Check(PlayerAnswer) (Points, error)
}

// QuestionAnswer is the answer to the question.
type QuestionAnswer interface {
	// Check if player answer is correct.
	Check(PlayerAnswer) (bool, error)
}

// Question for viewing by player.
type Question interface {
	// Description of the question.
	Description() (string, error)
	// Options for the answer.
	Options() ([]string, error)
	// Duration of the question.
	Duration() (time.Duration, error)
}

// PlayerAnswer is the answer provided by player.
type PlayerAnswer interface {
	// Options of the player answer.
	Options() ([]string, error)
}

// Points of the game.
type Points int
