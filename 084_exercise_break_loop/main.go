package main

import (
	"fmt"
)

func main() {

	y := 0
	for {
		fmt.Printf("y is: %v\n", y)
		if y > 20 {
			break
		}
		y++
	}

	x := 20
	for {
		if x < 0 {
			break
		}
		fmt.Printf("x is: %v\n", x)
		x--
	}

}
