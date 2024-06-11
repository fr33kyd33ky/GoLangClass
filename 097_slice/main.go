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

func main() {
	/*
		SLICE
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
		fmt.Println("-----------------")

		for i := range xs {
			fmt.Printf("i %v \t val %v\n", i, xs[i])
		}
		fmt.Println("-----------------")

		for i, v := range xs {
			fmt.Printf("i %v \t val %v\n", i, v)
		}
		fmt.Println("-----------------")

		xs = []string{
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
		fmt.Println("-----------------------------")
	*/

	/*
		APPEND
		xi := []int{42, 43, 44}
		fmt.Println(xi)
		fmt.Println("-----------------------------")

		xi = append(xi, 45)
		fmt.Println(xi)
		fmt.Println("-----------------------------")

		xi = append(xi, 46, 47, 48)
		fmt.Println(xi)
		fmt.Println("-----------------------------")
	*/

	/*
		//SLICE A SLICE
		xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		fmt.Printf("xi - %v\n", xi)
		fmt.Println("-----------------------------")

		fmt.Printf("xi - %v\n", xi[0:4])
		fmt.Println("-----------------------------")

		fmt.Printf("xi - %v\n", xi[:7])
		fmt.Println("-----------------------------")

		fmt.Printf("xi - %v\n", xi[4:])
		fmt.Println("-----------------------------")
	*/
	/*
		//DELETE FROM SLICE

		xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		fmt.Printf("xi - %v\n", xi)
		fmt.Println("-----------------------------")

		xi = append(xi[:4], xi[5:]...)
		fmt.Printf("xi - %v\n", xi)
		fmt.Println("-----------------------------")
	*/

	/*
		//MAKE
		si := []string{"a", "b", "c"}
		fmt.Println(si)
		fmt.Println("-----------------------------")

		xi := make([]int, 0, 10)
		fmt.Println(xi)
		fmt.Println(len(xi))
		fmt.Println(cap(xi))
		fmt.Println("-----------------------------")

		xi = append(xi, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
		fmt.Println(xi)
		fmt.Println(len(xi))
		fmt.Println(cap(xi))
		fmt.Println("-----------------------------")

		xi = append(xi, 10, 11, 12, 13)
		fmt.Println(xi)
		fmt.Println(len(xi))
		fmt.Println(cap(xi))
		fmt.Println("-----------------------------")
	*/

	/*
		//MULTIDIMENSION SLICE
		//like a spreadsheet
		//stores slice of strings
		jb := []string{"James", "Bond", "Martini", "Chocolate"}
		jm := []string{"Jenny", "Moneypenny", "Guiness", "Rocky Road"}
		fmt.Println(jb)
		fmt.Println(jm)
		fmt.Println("-----------------------------")

		//stores slice of a slice of strings
		xxs := [][]string{jb, jm}
		fmt.Println(xxs)
		fmt.Println("-----------------------------")
	*/
	/*
		a := []int{0, 1, 2, 3, 4, 5}
		b := a //important
		fmt.Println("a ", a)
		fmt.Println("b", b)
		fmt.Println("-----------------------------")

		//this changes a AND b. both are pointing to the same array because b:= a
		a[0] = 7
		fmt.Println("a ", a)
		fmt.Println("b", b)
		fmt.Println("-----------------------------")
	*/

	/*
		a := []int{0, 1, 2, 3, 4, 5}
		b := make([]int, len(a)) //important, create new object to store the values
		copy(b, a)               //important, make a copy of b from a
		fmt.Println("a ", a)
		fmt.Println("b ", b)
		fmt.Println("-----------------------------")

		a[0] = 7
		fmt.Println("a ", a)
		fmt.Println("b ", b)
		fmt.Println("-----------------------------")
	*/

}
