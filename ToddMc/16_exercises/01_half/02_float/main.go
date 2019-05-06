package main

import (
	"fmt"
)

func halve (n int) (float64, bool) {
	return float64(n) , n%2 == 0
}

func main () {
	i, even := halve(8)
	fmt.Println(i, even)
}
