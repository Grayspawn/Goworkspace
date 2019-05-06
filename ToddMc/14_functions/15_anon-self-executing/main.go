package main

import "fmt"

func main () {
	func () {
		fmt.Println("Anonymous Self executing Function")
	} ()
}
