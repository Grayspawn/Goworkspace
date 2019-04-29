package main

import "fmt"

func main () {

	x := 13 % 3 // % = remainder command
	fmt.Println(x)

	if x == 1{
		fmt.Println("Odd")
	} else {
		fmt.Print("Even")
	}

	for i := 0; i < 100; i++ {
		if i % 2 == 1 {
			fmt.Println("Odd")
		} else {
			fmt.Println("Even")
		}
	}
}
