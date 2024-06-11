package main

import "fmt"

func main() {

	a := [5]int{42, 43, 44}
	fmt.Println(a)

	a[3] = 45
	a[4] = 46

	fmt.Println(a)

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	fmt.Println("----------------------------")

	b := [5]int{}
	for i := 0; i < 5; i++ {
		b[i] = i
	}
	fmt.Println(b)

}
