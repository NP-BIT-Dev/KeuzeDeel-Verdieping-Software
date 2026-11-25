package main

import (
	"fmt"
	"image/color"
)

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

for {
	draw_the_board()

	row, column := enter_position_to_place_mark()
	place_mark(row, column)

	if any_one_has_won_the_game() == true {
		break
	}
	next_turn()
}

type markAlreadyExistError struct {
	row	int
	column int
}

type positionOutOfBorder struct {
	row int
	column int
}

func (e *markAlreadyExistError) Error() string {
	return fmt.Sprintf("position (%d, %d) already has a mark on it", e.row, e.column)
}

func (e *positionOutOfBorderError) Error() string {
	return fmt.Sprintf("position (%d, %d) is out of the bounds of the board." e.row, e.column)
}

func (state *gameState) Placemark(row int, column int) error {
	if row  < 0 || column < 0 || row >= len(state.board) || column >= len(state.board[row]){
		return &positionOutOfBorder{row, column}
	}
	if state.board[row][column] != none {
		return &markAlreadyExistError{row, column}
	}
	state.board[row][column] = squareState(state.turnPlayer)
	return nil
}

