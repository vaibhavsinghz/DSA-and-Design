package models

import "Self/snakeAndLadder/utils"

type Board struct {
	Size  int
	Cells [][]ICell
}

func NewBoard(size, numOfSnakes, numOfLadders int) *Board {
	cells := make([][]ICell, size)
	for i := range cells {
		cells[i] = make([]ICell, size)
		for j := range cells[i] {
			cells[i][j] = NewCell()
		}
	}
	board := &Board{
		Size:  size,
		Cells: cells,
	}
	board.addSnakesLadders(numOfSnakes, numOfLadders)
	return board
}

func (board *Board) addSnakesLadders(numOfSnakes, numOfLadders int) {
	for numOfSnakes > 0 {
		snakeHead, snakeTail := utils.RandInt(1, board.Size*board.Size-1), utils.RandInt(1, board.Size*board.Size-1)
		if snakeTail >= snakeHead {
			continue
		}

		board.GetCell(snakeHead).SetJump(NewJump(snakeHead, snakeTail))
		numOfSnakes--
	}

	for numOfLadders > 0 {
		ladderStart, ladderEnd := utils.RandInt(1, board.Size*board.Size-1), utils.RandInt(1, board.Size*board.Size-1)
		if ladderStart >= ladderEnd {
			continue
		}

		board.GetCell(ladderStart).SetJump(NewJump(ladderStart, ladderEnd))
		numOfLadders--
	}
}

func (board *Board) GetCell(position int) ICell {
	x := position / board.Size
	y := position % board.Size
	return board.Cells[x][y]
}

func (board *Board) GetBoardSize() int {
	return board.Size
}
