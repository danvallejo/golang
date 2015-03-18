package main

import (
	"fmt"
)

func modify(x *int) {
	*x = 0
}

func main() {
	score := 22
	modify(&score)
	fmt.Println(score)

	highscore := new(int)
	*highscore = 22
	fmt.Println(highscore, *highscore)
}
