package main

import (
	"sync"
	"fmt"
)

func main() {
	in := gen()

	// FAN OUT
	// Multiple functions reading from the same channel until that channel
	// is closed
	// Distribute work across multiple functions (goroutines) that all
	// read from it
	c0 := factorial(in)
	c1 := factorial(in)
	c2 := factorial(in)
	c3 := factorial(in)
	c4 := factorial(in)
	c5 := factorial(in)
	c6 := factorial(in)
	c7 := factorial(in)
	c8 := factorial(in)
	c9 := factorial(in)

	// FAN IN
	// Multiplex multiple channels onto a single channel
	// Merge the cahnnels from c0 through c9 onto a single channel
	for n := range merge(c0, c1, c2, c3, c4, c5, c6, c7, c8, c9) {
		fmt.Println(n)
	}
}

// This function makes a channel and puts 1000 values into it
// Then closes the channel and returns it
func gen() <-chan int {
	out := make(chan int)
	go func() {
		// make 10 to 100 so we're doing 1000 factorials
		for i := 0; i < 100; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

// This function 
func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
	}
}
