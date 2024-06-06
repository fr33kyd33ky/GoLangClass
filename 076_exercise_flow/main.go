package main

import (
	"fmt"
	"math/rand"
)

func main() {

	x := rand.Intn(250)
	fmt.Printf("x = %v\n", x)

	// if x > 0 && x <= 100 {
	// 	fmt.Println("between 0 and 100")
	// } else if x > 100 && x <= 200 {
	// 	fmt.Println("between 101 and 200")
	// } else if x > 200 && x <= 250 {
	// 	fmt.Println("between 201 and 250")
	// }

	//0 is included
	if x <= 100 {
		fmt.Println("between 0 and 100")
	} else if x >= 101 && x <= 200 {
		fmt.Println("between 101 and 200")
	} else if x >= 201 && x <= 250 {
		fmt.Println("between 201 and 250")
	} else {
		fmt.Println("this was more than 250")
	}

}
