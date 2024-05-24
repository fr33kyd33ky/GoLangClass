package main

import "fmt"

func main() {

	//declare h as int with value 44
	var h int = 44
	fmt.Println(h)
	fmt.Println("------------")

	//declare g hold value of 0 and used when you want to use to check if it's 0.
	var g int
	fmt.Println(g)
	g = 43
	fmt.Println(g)
	fmt.Println("------------")

	//better way to declare vars. short declaration operator := determines type
	a := 42
	fmt.Print(a)
	fmt.Println("------------")

	//blank identiier _ is a place holder, prolly for a function retun.
	b, c, d, _, f := 1, 2, 3, 4, "happy"
	fmt.Println(b, c, d, f)
	fmt.Println("------------")
}
