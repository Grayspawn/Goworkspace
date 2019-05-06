package main

import "fmt"

func max (numbers ...int) int {
	fmt.Printf("%T\n", numbers)
	var largest int
	for _, r := range numbers {
		if r > largest {
			largest = r
		}
	}
	return largest
}
func main () {
	greatest := max(24, 52, 14, 141, 511, 22, 45, 78999)
	fmt.Println(greatest)
}
