package main

import "fmt"

func main () {
	greeting :=func () {
		// Above: when you assign a func to a variable it becomes a func expression
		fmt.Println("Hellow World")
	}
	greeting()
	fmt.Printf("%T \n", greeting)
}
