package main

import "fmt"

func main() {

	if !true {
		fmt.Println("This did not run")
	}

	if !false { // Double Negative
		fmt.Println("Now this ran!")
	}
}
