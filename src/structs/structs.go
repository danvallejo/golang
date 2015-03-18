package main

import (
	"fmt"
)

type Television struct {
	channel int
}

func (tv *Television) changeChannel(newChannel int) {
	if newChannel > 1 && newChannel < 50 {
		tv.channel = newChannel
	}
}

func main() {
	var tv Television

	tv.channel = 2

	tv.changeChannel(13)

	fmt.Println(tv)

	bedroomTv := new(Television)
	fmt.Println(bedroomTv)

	spareTv := Television{channel: 3}
	fmt.Println(spareTv)
}
