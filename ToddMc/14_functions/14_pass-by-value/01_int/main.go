package main

import "fmt"

func main () {
	age := 54
	changeMe(age)
	fmt.Println(age)
}

func changeMe (z int) {
	fmt.Println(z)
	z = 22
}

// when changeMe is called on line 8
// the value 44 is being passed as an argument
