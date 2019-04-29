package main

import "fmt"

func main () {
<<<<<<< HEAD

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
=======
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
>>>>>>> 39b0053f0c6a797308f16b2e14f2b07602734a84
		}
	}
}
