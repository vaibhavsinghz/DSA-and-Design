package main

import (
	"Self/tictactoe/models"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TicTacToeGame struct {
	Players   []*models.Player
	GameBoard *models.Board
}

func NewTicTacToeGame() *TicTacToeGame {
	ticTacToeGame := &TicTacToeGame{
		Players:   []*models.Player{},
		GameBoard: models.NewBoard(3),
	}

	player1 := models.NewPlayer("Player1", models.NewPlayingPieceX())
	player2 := models.NewPlayer("Player2", models.NewPlayingPieceO())

	ticTacToeGame.Players = append(ticTacToeGame.Players, player1)
	ticTacToeGame.Players = append(ticTacToeGame.Players, player2)

	return ticTacToeGame
}

func (t *TicTacToeGame) Start() string {
	for true {
		t.GameBoard.Print()

		emptyCells := t.GameBoard.GetEmptyCells()
		if len(emptyCells) == 0 { //this means that all cells are filled but there is no winner
			break
		}

		playerTurn := t.Players[0]

		fmt.Printf("Player: %s Enter row,column: ", playerTurn.Name)
		reader := bufio.NewReader(os.Stdin)
		s, _ := reader.ReadString('\n')
		values := strings.Split(strings.TrimSpace(s), ",")
		inputRow, _ := strconv.Atoi(values[0])
		inputColumn, _ := strconv.Atoi(values[1])

		if !t.GameBoard.AddPiece(inputRow, inputColumn, playerTurn.PlayerPiece) {
			fmt.Println("incorrect position chosen, try again")
			continue
		}

		t.Players = t.Players[1:]
		t.Players = append(t.Players, playerTurn)

		if t.IsCurrentPlayerWinner(inputRow, inputColumn, playerTurn.PlayerPiece.PieceType) {
			return playerTurn.Name
		}

	}
	return "Match Draw"
}

func (t *TicTacToeGame) IsCurrentPlayerWinner(row, col int, pieceType models.PieceType) bool {
	rowMatch, colMatch, diagonalMatch, antiDiagonalMatch := true, true, true, true
	gameBoard := t.GameBoard

	for i := 0; i < gameBoard.Size; i++ {
		if gameBoard.Board[row][i] == nil || gameBoard.Board[row][i].PieceType != pieceType {
			rowMatch = false
			break
		}
	}

	for i := 0; i < gameBoard.Size; i++ {
		if gameBoard.Board[i][col] == nil || gameBoard.Board[i][col].PieceType != pieceType {
			colMatch = false
			break
		}
	}

	for i, j := 0, 0; i < gameBoard.Size && j < gameBoard.Size; i, j = i+1, j+1 {
		if gameBoard.Board[i][j] == nil || gameBoard.Board[i][j].PieceType != pieceType {
			diagonalMatch = false
			break
		}
	}

	for i, j := 0, gameBoard.Size-1; i < gameBoard.Size && j >= 0; i, j = i+1, j-1 {
		if gameBoard.Board[i][j] == nil || gameBoard.Board[i][j].PieceType != pieceType {
			antiDiagonalMatch = false
			break
		}
	}

	return rowMatch || colMatch || diagonalMatch || antiDiagonalMatch

}
