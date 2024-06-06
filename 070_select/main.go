package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	d1 := time.Duration(rand.Int63n(250))
	d2 := time.Duration(rand.Int63n(250))

	//go routines run concurrent
	//defined by channel
	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- 41
	}()

	go func() {
		time.Sleep(d2 * time.Millisecond)
		ch2 <- 41
	}()

	//select
	// will run whichever finishes sleep faster
	select {
	case v1 := <-ch1:
		fmt.Printf("value from channel 1: %v\n", v1)
	case v2 := <-ch2:
		fmt.Printf("value from channel 2: %v\n", v2)
	}

}
