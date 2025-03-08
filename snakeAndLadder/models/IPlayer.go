package models

type IPlayer interface {
	MoveTo(newPosition int)
	GetPosition() int
	GetID() int
}
