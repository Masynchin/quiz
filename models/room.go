package models

// Room is the lobby for game.
type Room interface {
	// Id of room.
	Id() (RoomID, error)
	// Register player to this room.
	Join(Player) error
	// Remove player from this room.
	Quit(Player) error
	// Game for this room.
	Game() (Game, error)
	// Stop room for accepting players.
	Stop() error
}

// ID of Room.
type RoomID int
