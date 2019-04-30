package main

import "fmt"

func main() {

	a := 75

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a // code mankes variable b a pointer to the memory address where an int is stored (a).
	// (b) is of type int pointer
	// *int -- the * is part of the type -- (b) is of type *int
	fmt.Println(b)
}
