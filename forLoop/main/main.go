package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		// fmt.Println(i)
		for j := 0; j < 10; j++ {
			fmt.Println("i: ", i, "j: ", j, "i*j: ", i*j)
		}
	}

	k := 0
	for {
		k++
		if k%2 == 1 {
			continue
		}
		fmt.Println(k)
		if k > 10 {
			break
		}
	}
}
