package models

type IBoard interface {
	GetCell(position int) ICell
	GetBoardSize() int
}
