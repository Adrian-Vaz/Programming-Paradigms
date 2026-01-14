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
	for {
		fmt.Scan(&move)
		if move < 1 || move > 9 {
			fmt.Println("Invalid index, choose 1-9")
			continue
		}
		if board[move-1] != '-' {
			fmt.Println("Cell already occupied, choose another")
			continue
		}
		board[move-1] = player
		break
	}
}

func removeOpponentMove(player rune) {
	opponent := 'x'
	if player == 'x' {
		opponent = 'o'
	}
	// check if opponent has any move
	has := false
	for i := 0; i < 9; i++ {
		if board[i] == opponent {
			has = true
			break
		}
	}
	if !has {
		fmt.Println("No opponent moves to remove.")
		return
	}
	fmt.Println("You may remove one of your opponent's moves. Enter index (1-9):")
	var idx int
	for {
		fmt.Scan(&idx)
		if idx < 1 || idx > 9 {
			fmt.Println("Invalid index, choose 1-9")
			continue
		}
		if board[idx-1] != opponent {
			fmt.Println("That cell is not occupied by your opponent. Pick again.")
			continue
		}
		board[idx-1] = '-'
		fmt.Println("Removed opponent's move at", idx)
		break
	}
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
	for turn := 1; ; turn++ {
		player := nextPlayer()
		display()
		takeTurn(player)
		if winner(player) {
			display()
			fmt.Println("Player ", string(player), "wins!")
			break
		}
		// Choose whether to remove opponent's move every 3 turns.
		if turn%3 == 0 {
			fmt.Println("Removal opportunity: do you want to remove an opponent's move? (y/n)")
			var resp string
			for {
				fmt.Scan(&resp)
				if resp == "y" || resp == "Y" {
					removeOpponentMove(player)
					break
				} else if resp == "n" || resp == "N" {
					fmt.Println("No removal this turn.")
					break
				} else {
					fmt.Println("Please enter 'y' or 'n'.")
				}
			}
			display()
			// re-check winner in case game state changed (unlikely, but safe)
			if winner(player) {
				display()
				fmt.Println("Player ", string(player), "wins!")
				break
			}
		}
		// Safety cap to avoid infinite games when moves are continually removed
		if turn > 100 {
			fmt.Println("Game ended in a draw (turn limit reached).")
			break
		}
	}
}
