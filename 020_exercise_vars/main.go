package main

import "fmt"

//assigned as False
var c, python, java bool

const d int = 42

var i, j = 1, 2.0

func main() {
	//var i int //assigned as 0
	fmt.Println(i, c, python, java, d)
	fmt.Println("--------------")

	var c, python, java = true, false, "No"
	fmt.Println(i, j, c, python, java)
	fmt.Printf("%T %T %T %T %T", i, j, c, python, java)
	fmt.Println("--------------")

	k := 3
	fmt.Println(k, i, j, c, python, java)
}
