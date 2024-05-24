package main

import (
	"fmt"

	golangpuppy "github.com/fr33kyd33ky/GoLangPuppy"
)

func main() {
	s1 := golangpuppy.Bark()
	s2 := golangpuppy.Barks()

	fmt.Println(s1)
	fmt.Println(s2)
}
