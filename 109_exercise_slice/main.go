package main

import "fmt"

func main() {

	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(a)

	for i, v := range a {
		fmt.Printf("index %v\t%v\t%T\n", i, v, v)
	}

}
