package main

import "fmt"

func celebrate () func () string {
	return func () string {
		return "Another way to represent a func expression"
	}
}
func main () {
	greet := celebrate()  // func expression
	fmt.Println(greet())
}
