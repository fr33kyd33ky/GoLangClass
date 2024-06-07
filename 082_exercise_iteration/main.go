package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i <= 42; i++ {
		x := rand.Intn(5)

		switch x {
		case 0:
			fmt.Printf("x is 0 - num %v\n", i)
		case 1:
			fmt.Printf("x is 1 - num %v\n", i)
		case 2:
			fmt.Printf("x is 2 - num %v\n", i)
		case 3:
			fmt.Printf("x is 3 - num %v\n", i)
		case 4:
			fmt.Printf("x is 4 - num %v\n", i)
		default:
			fmt.Printf("Nothing - num %v\n", i)
		}
	}

}
