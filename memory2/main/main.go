package main

import "fmt"

// func parameters are passed by value
// so don't expect to manipulate the original variable
// unless you pass the reference and de-reference inside the func
func zero(x int) {
	fmt.Printf("in zero, &x: %p\n", &x)
	x = 0
}

func zeroPointer(x *int) {
	fmt.Printf("in zeroPointer, x: %p\n", x)
	*x = 0
}

func main() {
	x := 5
	fmt.Printf("in main, &x: %p\n", &x)
	zero(x)
	fmt.Printf("after zero(x), x: %d\n", x)

	zeroPointer(&x)
	fmt.Printf("after zeroPointer(&x), x: %d\n", x)

}
