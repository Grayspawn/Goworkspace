package main

import "fmt"

func main (){
	var x [58] int // Array
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[34])
	x[34] = 666

	fmt.Println(x[34])

}


// With numbers in the braces its an array (line 4) without [] then its a slice
// the numbers in the braces identify the number of values within the array