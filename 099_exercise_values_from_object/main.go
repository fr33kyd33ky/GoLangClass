package main

import "fmt"

func main() {

	xs := []string{
		"Almond Biscotti Café",
		"Banana Pudding ",
		"Balsamic Strawberry (GF)",
	}

	fmt.Println(xs)
	// len is not part of the data type
	fmt.Println(len(xs))
	fmt.Printf("%T\n", xs)
	fmt.Println("-----------------------------")

	//blank will not return a value. only used to iterate the slice for the values
	//we don't want the index
	for _, v := range xs {
		fmt.Printf("val %v\n", v)
	}
	fmt.Println("-----------------------------")

	fmt.Println(xs[0])
	fmt.Println(xs[1])
	fmt.Println(xs[2])
	//this doesn't work
	//fmt.Println(xs[3])
	fmt.Println("-----------------------------")

	for i := 0; i < len(xs); i++ {
		fmt.Println(xs[i])

	}

}
