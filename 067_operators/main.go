package main

import "fmt"

func main() {
	fmt.Println("Running the first statement")
	fmt.Println("Running the second statement")

	x := 40
	y := 5
	fmt.Printf(" x=%v \n y=%v\n", x, y)

	if x < 42 && y < 42 {
		fmt.Println("Both are less than 42")
	}

	if x > 40 || x < 42 {
		fmt.Println("Getting close x")
	}

	if x != 42 && y != 10 {
		fmt.Println("Both not true")
	}
}
