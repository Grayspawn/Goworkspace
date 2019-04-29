package main

import "fmt"

func main () {
	i :=0
	for {  // with no condition the program runs infinitely
		fmt.Println(i)
		i++
	}
}
