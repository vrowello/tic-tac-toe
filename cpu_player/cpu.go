package cpu_player

import (
	"math/rand"
)

func Comp_choice_col(comp_guess chan int) {
	comp_guess <-rand.Intn(3)
}

func Comp_choice_row(comp_guess chan int) {
	comp_guess <-rand.Intn(3)
}
