package models

import "fmt"

type Board struct {
	Size  int
	Board [][]*PlayingPiece
}

func NewBoard(size int) *Board {
	board := make([][]*PlayingPiece, size)
	for i := range board {
		board[i] = make([]*PlayingPiece, size)
	}
	return &Board{
		Size:  size,
		Board: board,
	}
}

func (b *Board) Print() {
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			if b.Board[i][j] != nil {
				fmt.Print(b.Board[i][j].PieceType)
			} else {
				fmt.Print(" ")
			}
			fmt.Print(" | ")
		}
		fmt.Println()
	}
}

func (b *Board) GetEmptyCells() [][2]int {
	emptyCells := [][2]int{}
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			if b.Board[i][j] == nil {
				emptyCells = append(emptyCells, [2]int{i, j})
			}
		}
	}
	return emptyCells
}

func (b *Board) AddPiece(row, col int, piece *PlayingPiece) bool {
	if b.Board[row][col] != nil {
		return false
	}
	b.Board[row][col] = piece
	return true
}
