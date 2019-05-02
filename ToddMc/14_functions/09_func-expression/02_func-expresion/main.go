package main

import "fmt"

func main () {

	greeting := func () {     // "func ()" is an Anonymous function lacking a titular identifier
		fmt.Println("This is an 'Anonymous function'")
	}
	greeting()
}
