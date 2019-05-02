package main

import "fmt"

func celebrate() func() string {
	return func () string {
		return "Printout of func expression"
	}
}

func main () {
	greet := celebrate()
	fmt.Println(greet())
	fmt.Printf("%T \n", greet())
}
