package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 5; i++ {

		fmt.Printf("outer loop = %v\n", i)
		for j := 0; j < 5; j++ {

			fmt.Printf("inner loop = %v\n", j)

		}
	}
}
