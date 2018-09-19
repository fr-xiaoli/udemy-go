package main

import (
	"fmt"
)

func filter(numbers []int, callback func(int) bool) []int {
	var result []int
	for _, v := range numbers {
		if callback(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	numbers := filter([]int{1, 2, 3, 4, 5, 6}, func(n int) bool {
		return n%2 == 1
	})
	fmt.Println(numbers)
}
