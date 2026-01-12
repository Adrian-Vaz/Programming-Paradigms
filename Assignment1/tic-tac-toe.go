package main

import (
	"fmt"
)

type SquareState int
type Player int

const (
	none   = iota
	cross  = iota
	circle = iota
)

func (e Player) String() string {
	switch e {
	case none:
		return "none"
	case cross:
		return "cross"
	case circle:
		return "circle"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

type GameState struct {
	board      [3][3]SquareState
	turnPlayer Player
}

func (state *GameState) drawBoard() {
	// (# challenge: use Stringer to simplify this function)
	for i, row := range state.board {
		for j, square := range row {
			fmt.Print(" ")
			switch square {
			case none:
				fmt.Print(" ")
			case cross:
				fmt.Print("X")
			case circle:
				fmt.Print("O")
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

type AlreadyExistError struct {
	row    int
	column int
}

type positionOutOfBoundError struct {
	row    int
	column int
}

func (e *AlreadyExistError) Error() string {
	return fmt.Sprintf("position (%d,%d) already has a mark on it.", e.row, e.column)
}

func (e *positionOutOfBoundError) Error() string {
	return fmt.Sprintf("position (%d,%d) is out of bound.", e.row, e.column)
}

func (state *GameState) placeMark(row int, column int) error {
	if row < 0 || column < 0 || row >= len(state.board) || column >= len(state.board[row]) {
		return &positionOutOfBoundError{row, column}
	}
	if state.board[row][column] != none {
		return &AlreadyExistError{row, column}
	}

	state.board[row][column] = SquareState(state.turnPlayer)
	return nil // no error
}

type gameResult int

const (
	noWinnerYet = iota
	crossWon
	circleWon
	draw
)

func (state *GameState) whoIsNext() Player {
	return state.turnPlayer
}

func (state *GameState) nextTurn() {
	if state.turnPlayer == cross {
		state.turnPlayer = circle
	} else {
		state.turnPlayer = cross
	}
}

func (state *GameState) checkWinner() gameResult {
	boardSize := len(state.board)

	checkLine := func(startRow int, startColumn int, deltaRow int, deltaColumn int) gameResult {
		var lastSquare SquareState = state.board[startRow][startColumn]
		row, column := startRow+deltaRow, startColumn+deltaColumn

		for row >= 0 && column >= 0 && row < boardSize && column < boardSize {

			if state.board[row][column] == none {
				return noWinnerYet
			}

			if lastSquare != state.board[row][column] {
				return noWinnerYet
			}

			lastSquare = state.board[row][column]
			row, column = row+deltaRow, column+deltaColumn
		}

		if lastSquare == cross {
			return crossWon
		} else if lastSquare == circle {
			return circleWon
		}

		return noWinnerYet
	}

	for row := 0; row < boardSize; row++ {
		if result := checkLine(row, 0, 0, 1); result != noWinnerYet {
			return result
		}
	}
	for column := 0; column < boardSize; column++ {
		if result := checkLine(column, 0, 0, 1); result != noWinnerYet {
			return result
		}
	}
	if result := checkLine(0, 0, 1, 1); result != noWinnerYet {
		return result
	}
	if result := checkLine(0, boardSize-1, 1, -1); result != noWinnerYet {
		return result
	}
	for _, row := range state.board {
		for _, square := range row {
			if square == none {
				return noWinnerYet
			}
		}
	}
	return draw
}

func main() {
	state := GameState{}
	state.turnPlayer = cross // cross goes first

	var result gameResult = noWinnerYet

	for {
		fmt.Printf("next player to place a mark is: %v\n", state.whoIsNext())

		state.drawBoard()

		fmt.Printf("where to place a %v? (input row then column, separated by space)\n> ", state.whoIsNext())

		for {
			var row, column int
			fmt.Scan(&row, &column)

			e := state.placeMark(row-1, column-1)

			if e == nil {
				break
			}

			fmt.Println(e)
			fmt.Printf("please re-enter a position:\n> ")
		}

		result = state.checkWinner()
		if result != noWinnerYet {
			break
		}

		state.nextTurn()

		fmt.Println()
	}

	state.drawBoard()

	switch result {
	case crossWon:
		fmt.Printf("cross won the game!\n")
	case circleWon:
		fmt.Printf("circle won the game!\n")
	case draw:
		fmt.Printf("the game has ended with a draw...\n")
	}
}
