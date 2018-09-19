package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c) // when a channel is closed, no value can be put into it any more
	}()

	// Below is waiting to receive from the c until c is closed
	for n := range c {
		fmt.Println(n)
	}
}
