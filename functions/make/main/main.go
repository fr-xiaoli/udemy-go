package main

import (
	"fmt"
)

func main() {
	// slice is a reference type which points to the underlying array
	// it should be created with "make"
	s := make([]string, 1, 25)
	fmt.Println("main -> before calling -> s:", s)
	changeMySlice(s)
	fmt.Println("main -> after calling -> s:", s)

	// map is the same
	m := make(map[string]int)
	fmt.Println("main -> m", m)
	changeMyMap(m)
	fmt.Println("main after calling -> m", m)
}

func changeMySlice(z []string) {
	z[0] = "Me"
	fmt.Println("changeMe -> z: ", z)
}

func changeMyMap(z map[string]int) {
	z["Me"] = 42
}
