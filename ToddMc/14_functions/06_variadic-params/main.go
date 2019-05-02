package main

import "fmt"

func main () {
	n := average(45, 48, 455.789, 599, 2233, 0.99941)
	fmt.Println(n)
}

func average (sf ...float64) float64 {
	fmt.Println(len(sf))
	fmt.Printf("%T \n", sf)
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

// variadic is defined by a prefix of 3 dots '...' meaning can pass an infinite number of parameters. In this case of type float64