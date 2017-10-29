package cpu_player

import (
	"math/rand"
)

var (
	comp_row int
	comp_col int
)

func Comp_choice_col() int {
	comp_col = rand.Intn(3)
	return (comp_col)
}

func Comp_choice_row() int {
	comp_row = rand.Intn(3)
	return (comp_row)
}
