package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello")
	fmt.Println("The time is ", time.Now())
	fmt.Println("--------------")

	fmt.Println("My favorite number is", rand.Intn(1000))
	fmt.Println("--------------")

	fmt.Printf("Now you have %g problems\n", math.Sqrt(7))
	fmt.Println("--------------")

}
