package main

import "fmt"

func main() {
	fmt.Println("Running the first statement")
	fmt.Println("Running the second statement")

	x := 45
	y := 5
	fmt.Printf(" x=%v \n y=%v\n", x, y)

	if x < 42 {
		fmt.Println("Less than x")
	}

	if x < 42 {
		fmt.Println("Less than x")
	} else {
		fmt.Println("Equal to or greater x")
	}

	if x < 42 {
		fmt.Println("Less than x")
	} else if x == 42 {
		fmt.Println("Equal to x")
	} else {
		fmt.Println("Greater than x")
	}
}
