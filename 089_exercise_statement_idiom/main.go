package main

import (
	"fmt"
	"math/rand"
)

func main() {

	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	age := m["James"]
	fmt.Println(age)

	age2 := m["Q"]
	fmt.Println(age2)

	if res, ok := m["Q"]; !ok {
		fmt.Printf("not ok result %v\n", res)
	}

	c := 1
	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("iteration %v \t count %v \t x is %v\n", i, c, x)
			c++
		}
	}
}
