package main

import (
	"fmt"
	"math/rand"
)

func main() {

	//countit()

	for i := 0; i < 100; i++ {
		x := rand.Intn(10)
		y := rand.Intn(10)
		fmt.Printf("x = %v - y = %v\n", x, y)

		switch {
		case x < 4 && y < 4:
			fmt.Println("Both are less than 4", i)
		case x > 6 && y > 6:
			fmt.Println("Both are more than 6", i)
		case x >= 4 && x <= 6:
			fmt.Println("x between 4 and 6", i)
		case y != 5:
			fmt.Println("y is not 5", i)
		default:
			fmt.Println("Nothing is true", i)

		}

	}
}

func countit() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
}
