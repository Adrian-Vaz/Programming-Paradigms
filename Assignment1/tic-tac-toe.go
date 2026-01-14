package main

import "fmt"

var board [9]rune = [9]rune{'-', '-', '-', '-', '-', '-', '-', '-', '-'}

func display() {
	fmt.Println(string(board[0:3]))
	fmt.Println(string(board[3:6]))
	fmt.Println(string(board[6:9]))
}

var winningCombos [8][3]int = [8][3]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {0, 3, 6}, {1, 4, 7}, {2, 5, 8}, {0, 4, 8}, {2, 4, 6}}

func checkCombo(combo [3]int, player rune) bool {
	return board[combo[0]] == player && board[combo[1]] == player && board[combo[2]] == player
}

func winner(player rune) bool {
	for _, combo := range winningCombos {
		if checkCombo(combo, player) {
			return true
		}
	}
	return false
}

func takeTurn(player rune) {
	fmt.Println("it's player ", string(player), "'s turn")
	fmt.Println("Enter move index (1-9)")
	var move int
	fmt.Scan(&move)
	board[move-1] = player
}

func next() func() rune {
	player := 'x'
	return func() rune {
		if player == 'o' {
			player = 'x'
		} else {
			player = 'o'
		}
		return player
	}
}

func main() {
	nextPlayer := next()
	for i := 0; i < 9; i++ {
		player := nextPlayer()
		display()
		takeTurn(player)
		if winner(player) {
			display()
			fmt.Println("Player ", string(player), "wins!")
			break
		}
	}
}
