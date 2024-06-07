package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Printf("x = %v - y = %v\n", x, y)

	if x < 4 && y < 4 {
		fmt.Println("Both are less than 4")
	} else if x > 6 && y > 6 {
		fmt.Println("Both are more than 6")
	} else if x >= 4 && x <= 6 {
		fmt.Println("x between 4 and 6")
	} else if y != 5 {
		fmt.Println("y is not 5")
	} else {
		fmt.Println("Nothing is true")
	}
}
