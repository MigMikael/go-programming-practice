package main

import "fmt"

func printBoard(board [3][3]int) {
	for _, array := range board {
		for _, value := range array {
			fmt.Print(value, " ")
		}
		fmt.Println()
	}
}

func checkWinner(board [3][3]int, turn int) bool {
	isWinner := false

	// check earh row
	for _, array := range board {
		markCount := 0
		for _, value := range array {
			if value == turn {
				markCount++
			}
		}
		if markCount == 3 {
			return true
		}
	}

	// check each col
	for i := 0; i < 3; i++ {
		markCount := 0
		for j := 0; j < 3; j++ {
			if board[j][i] == turn {
				markCount++
			}
		}
		if markCount == 3 {
			return true
		}
	}

	// check diagonal
	markCount := 0
	for i := 0; i < 3; i++ {
		if board[i][i] == turn {
			markCount++
		}
	}
	if markCount == 3 {
		return true
	}

	markCount = 0
	for i := 0; i < 3; i++ {
		if board[2-i][i] == turn {
			markCount++
		}
	}
	if markCount == 3 {
		return true
	}

	return isWinner
}

func gameXO() {
	var board [3][3]int
	printBoard(board)

	turnCounter := 0
	for ; ; turnCounter++ {
		var input int
		var winner bool
		var winMsg string
		fmt.Scan(&input)
		input--
		if turnCounter%2 == 0 { // first player turn
			board[input/3][(input % 3)] = 1
			winner = checkWinner(board, 1)
			winMsg = "Player 1 Win"
		} else { // second player turn
			board[input/3][(input % 3)] = 2
			winner = checkWinner(board, 2)
			winMsg = "Player 2 Win"
		}
		printBoard(board)

		if winner {
			fmt.Println(winMsg)
			break
		}
	}
}
