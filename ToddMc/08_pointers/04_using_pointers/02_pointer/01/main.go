package main

import "fmt"

func zero(z *int) {
	*z = 0
}

func main() {

	x := 5
	zero(&x)
	fmt.Println(x) // x is now 0
}

// x is 5
// pass over memory address to func zero
// In func  zero receive memory address
// It is dereferenced at that memory location
// Assign the value 0 at the memory address
// Print 0
