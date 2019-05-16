package main

import "fmt"

func main () {
	var y [256] byte

	fmt.Println(len(y))
	fmt.Println(y[42])

	for g :=0; g < 256; g++{
		y[g] = byte(g)
	}

	for g, v := range y {
	fmt.Printf("%v - %T - %b\n", v, v, v)

	if g > 100 {
		break
	}
	}
}
