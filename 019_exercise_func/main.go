package main

import "fmt"

func main() {
	fmt.Println(add(42, 13))
	sayHello()
	fmt.Println("-------------")

	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println("-------------")

	fmt.Println(split(17))
	fmt.Println("-------------")

}

//both parameters are int
func add(x int, y int) int {
	return x + y
}

//both parameters are int
func add2(x, y int) int {
	return x + y
}

func sayHello() {
	fmt.Println("Hello")
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 9 / 4
	y = sum - x
	//assign x and y and return both
	return
}
