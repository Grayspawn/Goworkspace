package main

import "fmt"

func main () {

	a := 75
	fmt.Println(a)  // 75
	fmt.Println(&a) // 0xc000018060

	var b = &a      // referencing a memory address
	fmt.Println(b)  // 0xc000018060
	fmt.Println(*b) // 0xc000018060 *b is de-referencing a memory address


	// b is an int pointer
	// b points to the memory address that an int (a) is stored
	// to see the value "derefence" in that memory address add a * in front of b
	// This is know na s dereferencing
	// the * is an operator in this case

}