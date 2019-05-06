package main

import "fmt"

func main () {
	name := "Leslie"
	fmt.Println(name)  // Leslie

	changeMe(name)

	fmt.Println(name) // Leslie
}

func changeMe (z string) {
	fmt.Println(z)   // Leslie
	z = "- Rocky -"
	fmt.Println(z)   // Rocky
}
