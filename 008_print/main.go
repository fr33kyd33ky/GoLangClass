package main

import "fmt"

func main() {
	const name, age = "Kim", 22
	//%s = string
	//%d = decimal
	//%T = type
	fmt.Printf("%s is %d years old.\n", name, age)
	fmt.Printf("%s is %d years old.\t and type %T and %T", name, age, name, age)
}
