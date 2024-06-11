package main

import "fmt"

func main() {

	jb := []string{"James", "Bond", "Shaken, not stirred"}
	jm := []string{"Miss", "Moneypenny", "I'm 008"}

	xxs := [][]string{jb, jm}

	// for x := range xxs {
	// 	for y := range jb {
	// 		fmt.Println(xxs[x][y])
	// 	}
	// }

	for i, v := range xxs {
		fmt.Println(i, v)
		for a, b := range v {
			fmt.Println(a, b)
		}
	}
}
