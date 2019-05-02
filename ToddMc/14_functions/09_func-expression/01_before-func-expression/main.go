package main

import "fmt"

func greet () {
	fmt.Println("Before Func Expression")
}

func main () {
greet()  // No 'func' expression in this program. Func main calls a func greet
}
