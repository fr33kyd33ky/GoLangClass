package main

import "fmt"

func main() {

	var x string = "one"
	var y = 1
	z := 1.11

	fmt.Printf("the value of x is %v and type is %T\n", x, x)
	fmt.Printf("the value of y is %v and type is %T\n", y, y)
	fmt.Printf("the value of z is %v and type is %T\n", z, z)

	a, b, c := "two", 2, 2.22
	fmt.Printf("the value of a is %v and type is %T\n", a, a)
	fmt.Printf("the value of b is %v and type is %T\n", b, b)
	fmt.Printf("the value of c is %v and type is %T\n", c, c)

}
