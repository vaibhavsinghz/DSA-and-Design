package main

import "fmt"

func main() {
	TTTGame := NewTicTacToeGame()
	winner := TTTGame.Start()
	fmt.Println("Gamer winner is", winner)
}
