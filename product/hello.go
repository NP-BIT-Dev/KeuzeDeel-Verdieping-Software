package main

import "fmt"

type squareState int 

const (
	none = iota
	cross = iota
	circle = iota
)

var board [3][3] squareState

type player int

var turnPlayer player

type gameState struct {
	board[3][3]squareState
	turnPlayer player
}

func (state *gameState) drawBoard() {
	for i, row := range state.board {
		for j, square := range row{
			fmt.Print(" ")
			switch square{
		case none:
			fmt.Print(" ")
		case cross: 
			fmt.Print("X")
		case circle:
			fmt.Print("0")
		}
		if j != len(row)-1 {
			fmt.Print(" |")
		}
	}
	if i != len(state.board)-1 {
		fmt.Print("\n------------")
	}
			fmt.Print("\n")
	}
}

func main() {
	state := gameState{}
	
	state.board[0][1] = cross
	state.board[1][1] = circle

	state.drawBoard()
}
