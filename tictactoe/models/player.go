package models

type Player struct {
	Name        string
	PlayerPiece *PlayingPiece
}

func NewPlayer(name string, playerPiece *PlayingPiece) *Player {
	return &Player{
		Name:        name,
		PlayerPiece: playerPiece,
	}
}
