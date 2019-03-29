package main

import "fmt"

func zero (x *int) {
	*x = 0
}

func main () {
	x := 5
	zero(&x)
	fmt.Println(x)
}

//pointers(*int) reference a memory location where a value is stored versus the value itself.By using a pointer (*int) the zero function is able to modify the variable x

// A pointer is represented by the asterisk '*' characted followed by the type of the value to be stored. In the zero function xPtr is pointer to an integer

// the asterisk '*' is also used to "dereference" pointer variables. Dereferencing a pointer gives us access to the value of the pointer it points to

// *x=0 means "store the into 0 in the memory location assigned to *x

// the amperand operator '&'is used to find the memory address of a variable
