package main

import "fmt"

func main () {
	greet("Jane")    		// greet is declared here with a parameter
	greet("John")
}

func greet (name string){ 		// when calling greet pass in an argument
	fmt.Println(name)
}
