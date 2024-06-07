/*
● array
○ a numbered sequence of VALUES of the same TYPE
○ does not change in size
○ used for Go internals; generally not used in your code
○ https://golang.org/ref/spec#Array_types

● slice
○ built on top of an array
○ holds VALUES of the same TYPE
○ changes in size
○ has a length and a capacity
○ https://golang.org/ref/spec#Slice_types

● map
○ key / value storage
○ an unordered group of VALUES of one TYPE, called the element type, indexed
by a set of unique keys of another type, called the key type.
○ https://golang.org/ref/spec#Map_types

● struct
○ a data structure
○ a composite / aggregate
*/

/*
	slice is a data structure with 3 elements
	len, cap, pointer to underlying array

	type slice struct{
		array	unsafe.Pointer
		len		int
		cap		int
	}

*/

package main

import "fmt"

func main() {
	// xs := []string{"hello", "world"}
	// fmt.Println(xs)
	// fmt.Println(len(xs))
	// fmt.Printf("%T", xs)

	xs := []string{
		"Almond Biscotti Café",
		"Banana Pudding ",
		"Balsamic Strawberry (GF)",
		"Bittersweethocolate (GF)",
		"Blueberry Pancake (GF)",
		"Bourbon Turtle (GF)",
		"Browned Butter	Cookie Dough",
		"Chocolate Covered Black Cherry (GF)",
		"Chocolate Fudge Brownie",
		"Chocolate Peanut Butter (GF)",
		"Coffee (GF)",
		"Cookies & Cream",
		"Fresh Basil (GF)",
		"Garden Mint Chip (GF)",
		"Lavender Lemon Honey (GF)",
		"Lemon Bar",
		"Madagascar	Vanilla (GF)",
		"Matcha (GF)",
		"Midnight Chocolate Sorbet (GF, V)",
		"Non-Dairy Chocolate Peanut Butter (GF, V)",
		"Non-Dairy Coconut Matcha (GF, V)",
		"Non-Dairy SunbutterCinnamon (GF, V)",
		"Orange Cream (GF) ",
		"Peanut Butter Cookie Dough",
		"RaspberrySorbet (GF, V)",
		"Salty Caramel (GF)",
		"Slate Slate Different",
		"Strawberry LemonadeSorbet (GF, V)",
		"Vanilla Caramel Blondie",
		"Vietnamese Cinnamon (GF)",
		"WolverineTracks (GF)",
	}

	fmt.Println(xs)
	// len is not part of the data type
	fmt.Println(len(xs))
	fmt.Printf("%T", xs)

	// for i := range xs {
	// 	fmt.Printf("i %v \t val %v\n", i, xs[i])
	// }

	for i, v := range xs {
		fmt.Printf("i %v \t val %v\n", i, v)
	}

}
