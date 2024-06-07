package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is init")
}

func main() {
	var x int = rand.Intn(250)
	fmt.Printf("x = %v\n", x)

	switch {
	case x <= 100:
		fmt.Println("between 0 and 100")
	case x >= 101 && x <= 200:
		fmt.Println("between 101 and 200")
	case x >= 201 && x <= 250:
		fmt.Println("between 201 and 250")
	default:
		fmt.Println("this was more than 250")
	}

}
