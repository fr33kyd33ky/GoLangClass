package main

import "fmt"

func main() {

	x := 40
	y := 5
	fmt.Printf(" x=%v \n y=%v\n", x, y)

	// c# style
	for i := 0; i < 5; i++ {
		fmt.Printf("counting from 1 to 5: %v \t first loop\n", i)
	}

	// like c# while statement
	for y < 10 {
		fmt.Printf("y is: %v \t second loop\n", y)
		y++
	}

	// break leaves the loop
	for {
		fmt.Printf("y is: %v \t third loop\n", y)
		if y > 20 {
			break
		}
		y++
	}

	// if statement is true then continue to
	for i := 0; i < 20; i++ {
		// if value is odd then continue to next iteration of loop
		// if value is even then it moves to the next step in the iteration
		if i%2 != 0 {
			continue
		}
		fmt.Printf("counting even numbers: %v \t fourth loop\n", i)
	}

}
