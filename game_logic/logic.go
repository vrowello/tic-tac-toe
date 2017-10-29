package game_logic

import (
	"fmt"
)

var (
	board     [3][3]string
	game_over bool
	turns     int
	guess_row int
	guess_col int
	comp_row  int
	comp_col  int
)

func Print_board(board [3][3]string) {
	fmt.Println(board[0])
	fmt.Println(board[1])
	fmt.Println(board[2])
}

func Game_check(board [3][3]string) bool {
	if board[0][0] == "x" && board[0][1] == "x" && board[0][2] == "x" {
		fmt.Println("You Win!")
		game_over = true
	} else if board[1][0] == "x" && board[1][1] == "x" && board[1][2] == "x" {
		fmt.Println("You Win!")
		game_over = true
	} else if board[2][0] == "x" && board[2][1] == "x" && board[2][2] == "x" {
		fmt.Println("You Win!")
		game_over = true
	} else if board[0][0] == "x" && board[1][1] == "x" && board[2][2] == "x" {
		fmt.Println("You Win!")
		game_over = true
	} else if board[2][0] == "x" && board[1][1] == "x" && board[0][2] == "x" {
		fmt.Println("You Win!")
		game_over = true
	} else if board[0][0] == "x" && board[1][0] == "x" && board[2][0] == "x" {
		fmt.Println("You Win!")
		game_over = true
	} else if board[0][1] == "x" && board[1][1] == "x" && board[2][1] == "x" {
		fmt.Println("You Win!")
		game_over = true
	} else if board[0][2] == "x" && board[1][2] == "x" && board[2][2] == "x" {
		fmt.Println("You Win!")
		game_over = true
	} else if board[0][0] == "o" && board[0][1] == "o" && board[0][2] == "o" {
		fmt.Println("You lose!")
		game_over = true
	} else if board[1][0] == "o" && board[1][1] == "o" && board[1][2] == "o" {
		fmt.Println("You lose!")
		game_over = true
	} else if board[2][0] == "o" && board[2][1] == "o" && board[2][2] == "o" {
		fmt.Println("You lose!")
		game_over = true
	} else if board[0][0] == "o" && board[1][1] == "o" && board[2][2] == "o" {
		fmt.Println("You lose!")
		game_over = true
	} else if board[2][0] == "o" && board[1][1] == "o" && board[0][2] == "o" {
		fmt.Println("You lose!")
		game_over = true
	} else if board[0][0] == "o" && board[1][0] == "o" && board[2][0] == "o" {
		fmt.Println("You lose!")
		game_over = true
	} else if board[0][1] == "o" && board[1][1] == "o" && board[2][1] == "o" {
		fmt.Println("You lose!")
		game_over = true
	} else if board[0][2] == "o" && board[1][2] == "o" && board[2][2] == "o" {
		fmt.Println("You lose!")
		game_over = true
	} else if turns == 10 {
		fmt.Println("Draw!")
		game_over = true
	}
	return game_over
}
