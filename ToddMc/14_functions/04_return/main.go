package main

import "fmt"

func main () {
 	greet("Mick", "Steve")
}

func greet (fname, lname string) {
	fmt.Println(fname, lname)
}
