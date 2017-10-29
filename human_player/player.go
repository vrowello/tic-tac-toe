package human_player

import (
	"fmt"
)

var (
	guess_row int
	guess_col int
)

func Player_choice_col() int {
	fmt.Printf("Column (0. 1, or 2): ")
	fmt.Scanln(&guess_col)
	return (guess_col)
}

func Player_choice_row() int {
	fmt.Printf("Row (0. 1, or 2): ")
	fmt.Scanln(&guess_row)
	return (guess_row)
}
