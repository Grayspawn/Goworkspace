//Package level scope for the variable x
// The scope of x is the entire package as it is NOT within the curly braces of func main

package main

import "fmt"

var x = 42 //Declaration of variable x is type int & Variable x is is assigned value 42 which initialises the variable x 	fmt.Println(x)


func main() {
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println(x)
}
