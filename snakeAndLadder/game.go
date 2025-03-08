package main

import (
	"Self/snakeAndLadder/models"
	"fmt"
)

type Game struct {
	Board  models.IBoard
	Dice   models.IDice
	Player []models.IPlayer
}

func NewGame() IGame {
	game := &Game{
		Board: models.NewBoard(10, 5, 4),
		Dice:  models.NewDice1(1, 6),
	}
	game.addPlayer(2)
	return game
}

func (game *Game) addPlayer(count int) {
	for i := 1; i <= count; i++ {
		game.Player = append(game.Player, models.NewPlayer(i))
	}
}

func (game *Game) Play() {
	for true {
		currentPlayer := game.Player[0]
		fmt.Printf("Current Player is: %d, current position is %d\n", currentPlayer.GetID(), currentPlayer.GetPosition())

		diceNumber := game.Dice.Roll()
		newPosition := currentPlayer.GetPosition() + diceNumber
		newPosition = game.checkJump(newPosition)
		currentPlayer.MoveTo(newPosition)

		fmt.Printf("Current Player is: %d, new position is %d\n", currentPlayer.GetID(), currentPlayer.GetPosition())

		if currentPlayer.GetPosition() >= game.Board.GetBoardSize()*game.Board.GetBoardSize()-1 {
			fmt.Printf("Winner is %d\n", currentPlayer.GetID())
			break
		}

		game.Player = game.Player[1:]
		game.Player = append(game.Player, currentPlayer)
	}
}

func (game *Game) checkJump(playerPosition int) int {
	if playerPosition >= game.Board.GetBoardSize()*game.Board.GetBoardSize()-1 {
		return playerPosition
	}

	cell := game.Board.GetCell(playerPosition)
	jump := cell.GetJump()
	if jump != nil && jump.Start == playerPosition {
		if jump.End > jump.Start {
			fmt.Println("cell jump by ladder")
		} else {
			fmt.Println("cell jump by snake")
		}
		return jump.End
	}
	return playerPosition
}
