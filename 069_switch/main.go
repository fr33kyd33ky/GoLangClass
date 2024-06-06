package main

import (
	"fmt"
)

func main() {
	x := 40
	y := 5
	fmt.Printf(" x=%v \n y=%v\n", x, y)

	// switch

	//with default
	switch {
	case x < 42:
		fmt.Println("x is less than 42")
	case x == 42:
		fmt.Println("x equal to 42")
	case x > 42:
		fmt.Println("x is greater than 42")
	default:
		fmt.Println("default case for x")
	}

	//no fallthrough so exits when condition is met
	switch x {
	case 40:
		fmt.Println("x is 40")
	case 41:
		fmt.Println("x is 41")
	case 42:
		fmt.Println("x is 42")
	case 43:
		fmt.Println("x is 43")
	default:
		fmt.Println("default case")
	}

	//fallthrough means continue to the next case
	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 41:
		fmt.Println("printing because all fallthrough. x is 41")
		fallthrough
	case 42:
		fmt.Println("printing because all fallthrough. x is 42")
		fallthrough
	case 43:
		fmt.Println("printing because all fallthrough. x is 43")
		fallthrough
	default:
		fmt.Println("printing because all fallthrough. default for x")
	}

}
