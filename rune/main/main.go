package main

import "fmt"

func main() {
	for i := 50; i <= 140; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
	fmt.Println("hello 世界", []byte("hello 世界"))
	fmt.Println(`I have
many lines
yeah`)
	fmt.Sprint()
}
