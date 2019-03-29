package main

import "fmt"

func main () {
	x := 13 / 3 // result is 4
	fmt.Println(x)

	y := 13 % 3
	fmt.Println(y) // result is 1 using the remainder operator '&'

	if y == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}


	for i :=1; i <23; i++ {
		if i%2 == 1 {
			fmt.Println("Still Odd")
		} else {
			fmt.Println("Now very even")
		}
	}
}
