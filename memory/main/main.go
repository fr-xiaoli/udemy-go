package main

import (
	"fmt"
)

func main() {
	a := 43
	fmt.Printf("a:\t%d\n", a)
	fmt.Printf("&a:\t%p\n", &a)

	var b *int = &a // & is the "reference operator" which gets the address of the memory (the reference)
	fmt.Printf("b:\t%v\n", b)
	fmt.Printf("*b:\t%v\n", *b) // * is the "dereference operator" which gets the value saved in the address of the memory

	fmt.Println("~~~~~~~ *b = 42 ~~~~~~~")
	*b = 42 // change the value saved in address b to 42

	fmt.Printf("a:\t%d\n", a)
	fmt.Printf("&a:\t%p\n", &a)
}
