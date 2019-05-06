package main

import "fmt"

func main () {

	half := func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}

	i, even := half(8489)
	fmt.Println(i, even)
}