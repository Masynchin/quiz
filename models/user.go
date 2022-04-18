package models

// User of our app.
type User interface {
	// Id of the user.
	Id() (UserID, error)
	// Name of the user.
	Name() (string, error)
}

// ID of User.
type UserID int
