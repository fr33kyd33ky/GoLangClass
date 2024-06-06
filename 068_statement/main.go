package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// fmt.Println("Running the first statement")
	// fmt.Println("Running the second statement")

	x := 40
	y := 5
	fmt.Printf(" x=%v \n y=%v\n", x, y)

	// statement-statement idiom because two statements z := 2 * rand.Intn(x) and z >= x
	// scope of z is limited to this statement
	// assigns a value to z.
	// this is a random int value 0 and less than x.
	// check to see if z is greater that or equal to x
	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and that is greater than or equal to x which is %v\n", z, x)
	} else {
		fmt.Printf("z is %v and that is less than x which is %v\n", z, x)
	}

}
