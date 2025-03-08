package models

type Player struct {
	ID              int
	CurrentPosition int
}

func NewPlayer(id int) *Player {
	return &Player{
		ID: id,
	}
}

func (p *Player) MoveTo(newPosition int) {
	p.CurrentPosition = newPosition
}

func (p *Player) GetPosition() int {
	return p.CurrentPosition
}

func (p *Player) GetID() int {
	return p.ID
}
