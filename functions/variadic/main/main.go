package main

import (
	"fmt"
)

func main() {
	data := []float64{45, 32, 45, 63, 24} // a slice
	fmt.Println(average(data...))

	func() {
		fmt.Println("I'm anonymous")
	}()
}

func average(sf ...float64) float64 {
	var total float64 // varaibles declared with var keyword will be initialized
	// with zero value of that type
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
