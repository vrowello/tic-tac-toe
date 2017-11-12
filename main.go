package main

import (
	"github.com/vrowello/tic-tac-toe/cpu_player"
	"github.com/vrowello/tic-tac-toe/game_logic"
	"github.com/vrowello/tic-tac-toe/human_player"
)

var (
	board     [3][3]string
	game_over bool
	turns     int
	guess_row int
	guess_col int
	comp_row  int
	comp_col  int
	attempts  int = 1
)

func main() {
	board[0] = [3]string{"_", "_", "_"}
	board[1] = [3]string{"_", "_", "_"}
	board[2] = [3]string{"_", "_", "_"}

	comp_guess := make(chan int, 2)

	game_over = false
	turns = 1

	for game_over == false {

		go cpu_player.Comp_choice_col(comp_guess)
		go cpu_player.Comp_choice_row(comp_guess)

		guess_col = human_player.Player_choice_col()
		guess_row = human_player.Player_choice_row()

		if board[guess_col][guess_row] == "_" {
			board[guess_col][guess_row] = "x"
		} else {
			guess_col = human_player.Player_choice_col()
			guess_row = human_player.Player_choice_row()
		}

		turns += 1

		game_over = game_logic.Game_check(board)
		if game_over == true {
			break
		}

		comp_col = <-comp_guess
		comp_row = <-comp_guess

		if board[comp_col][comp_row] == "_" {
			board[comp_col][comp_row] = "o"
		} else {
			for board[comp_col][comp_row] != "_" {
				go cpu_player.Comp_choice_col(comp_guess)
				go cpu_player.Comp_choice_row(comp_guess)
				comp_col = <-comp_guess
				comp_row = <-comp_guess
			}
			board[comp_col][comp_row] = "o"
		}

		turns += 1

		game_logic.Print_board(board)

		game_over = game_logic.Game_check(board)
	}

	if game_over == true {
		game_logic.Print_board(board)
	}
}
