package main

import "fmt"

func hello () {
	fmt.Print("hello ")
}

func world () {
	fmt.Println("world")
}


func main () {
	defer world()  // defer delays or defers whatever it is going to run to just before function exits
	hello()
}